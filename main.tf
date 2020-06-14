terraform {
  required_version = "0.12.26"
  backend "s3" {
    bucket = "atlas-terraform-state-dev"
    key    = "covid-tf.state"
    region = "us-east-2"
  }
}
provider "aws" {
  region = "us-east-2"
}

data "aws_caller_identity" "current" {}

resource "aws_dynamodb_table" "covid-state-table" {
  name         = "CovidState"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "state"
  range_key    = "date"
  attribute {
    name = "state"
    type = "S"
  }
  attribute {
    name = "date"
    type = "N"
  }
}
resource "aws_dynamodb_table" "covid-us-table" {
  name         = "CovidUS"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "date"

  attribute {
    name = "date"
    type = "N"
  }
}
resource "aws_iam_role" "cacheRole" {
  name               = "covidCacheRole"
  assume_role_policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": {
    "Action": "sts:AssumeRole",
    "Principal": {
      "Service": "lambda.amazonaws.com"
    },
    "Effect": "Allow"
  }
}
POLICY
}
resource "aws_iam_role" "apiRole" {
  name               = "covidApiRole"
  assume_role_policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": {
    "Action": "sts:AssumeRole",
    "Principal": {
      "Service": "lambda.amazonaws.com"
    },
    "Effect": "Allow"
  }
}
POLICY
}
resource "aws_iam_policy" "lambda_logging" {
  name        = "lambda_logging"
  path        = "/"
  description = "IAM policy for logging from a lambda"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "arn:aws:logs:*:*:*",
      "Effect": "Allow"
    }
  ]
}
EOF
}
data "aws_iam_policy_document" "dynamoPut" {
  statement {
    actions = [
      "dynamodb:PutItem",
      "dynamodb:Scan",
      "dynamodb:Query"
    ]

    resources = [
      aws_dynamodb_table.covid-state-table.arn,
      aws_dynamodb_table.covid-us-table.arn
    ]
  }
}
resource "aws_iam_policy" "tablePolicy" {
  name   = "covidTablePolicy"
  path   = "/"
  policy = data.aws_iam_policy_document.dynamoPut.json
}
resource "aws_iam_role_policy_attachment" "cache_dynamo" {
  role       = aws_iam_role.cacheRole.name
  policy_arn = aws_iam_policy.tablePolicy.arn
}
resource "aws_iam_role_policy_attachment" "cache_logs" {
  role       = aws_iam_role.cacheRole.name
  policy_arn = aws_iam_policy.lambda_logging.arn
}
resource "aws_iam_role_policy_attachment" "api_logs" {
  role       = aws_iam_role.apiRole.name
  policy_arn = aws_iam_policy.lambda_logging.arn
}
resource "aws_iam_role_policy_attachment" "api_dynamo" {
  role       = aws_iam_role.apiRole.name
  policy_arn = aws_iam_policy.tablePolicy.arn
}
resource "aws_s3_bucket" "deployment_bucket" {
  bucket = "covidapideployment"
  acl    = "private"
  versioning {
    enabled = true
  }
}
resource "aws_s3_bucket_object" "cache_deployment" {
  bucket = aws_s3_bucket.deployment_bucket.id
  key    = "cache.zip"
  source = "dist/cache.zip"
  etag   = filemd5("dist/cache.zip")
}
resource "aws_lambda_function" "covidCache" {
  function_name     = "covidCache"
  s3_bucket         = aws_s3_bucket.deployment_bucket.id
  s3_key            = "cache.zip"
  s3_object_version = aws_s3_bucket_object.cache_deployment.version_id
  handler           = "cache"
  # source_code_hash  = filebase64("dist/cache.zip")
  role        = aws_iam_role.cacheRole.arn
  runtime     = "go1.x"
  memory_size = 128
  timeout     = 30
}
resource "aws_s3_bucket_object" "api_deployment" {
  bucket = aws_s3_bucket.deployment_bucket.id
  key    = "api.zip"
  source = "dist/api.zip"
  etag   = filemd5("dist/api.zip")
}
resource "aws_lambda_function" "covidAPIv2" {
  function_name     = "covidAPIv2"
  s3_bucket         = aws_s3_bucket.deployment_bucket.id
  s3_key            = "api.zip"
  s3_object_version = aws_s3_bucket_object.api_deployment.version_id
  handler           = "api"
  # source_code_hash = filebase64("dist/api.zip")
  role        = aws_iam_role.apiRole.arn
  runtime     = "go1.x"
  memory_size = 256
  timeout     = 30
}
resource "aws_apigatewayv2_api" "covidAPI" {
  name          = "covid-api"
  protocol_type = "HTTP"
}
resource "aws_lambda_permission" "apigw_lambda" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.covidAPIv2.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "arn:aws:execute-api:us-east-2:${data.aws_caller_identity.current.account_id}:${aws_apigatewayv2_api.covidAPI.id}/*/$default"
}
resource "aws_apigatewayv2_integration" "base" {
  api_id                 = aws_apigatewayv2_api.covidAPI.id
  integration_type       = "AWS_PROXY"
  description            = "Lambda example"
  integration_method     = "POST"
  passthrough_behavior   = "WHEN_NO_MATCH"
  integration_uri        = aws_lambda_function.covidAPIv2.invoke_arn
  payload_format_version = "1.0"
}
resource "aws_apigatewayv2_route" "default_route" {
  api_id    = aws_apigatewayv2_api.covidAPI.id
  route_key = "$default"
  target    = "integrations/${aws_apigatewayv2_integration.base.id}"
}
resource "aws_apigatewayv2_stage" "default_stage" {
  api_id      = aws_apigatewayv2_api.covidAPI.id
  name        = "$default"
  auto_deploy = true
}
resource "aws_apigatewayv2_domain_name" "covidapi" {
  domain_name = "covidapi.dev.cloudadaptr.com"

  domain_name_configuration {
    certificate_arn = var.certificate
    endpoint_type   = "REGIONAL"
    security_policy = "TLS_1_2"
  }
}
resource "aws_apigatewayv2_api_mapping" "apimapping" {
  api_id      = aws_apigatewayv2_api.covidAPI.id
  domain_name = aws_apigatewayv2_domain_name.covidapi.id
  stage       = aws_apigatewayv2_stage.default_stage.id
}

resource "aws_route53_record" "api" {
  name    = "covidapi.dev.cloudadaptr.com"
  type    = "A"
  zone_id = var.hosted_zone

  alias {
    evaluate_target_health = true
    name                   = aws_apigatewayv2_domain_name.covidapi.domain_name_configuration[0].target_domain_name
    zone_id                = aws_apigatewayv2_domain_name.covidapi.domain_name_configuration[0].hosted_zone_id
  }
}
