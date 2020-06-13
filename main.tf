terraform {
  backend "s3" {
    bucket = "atlas-terraform-state-dev"
    key    = "covid-tf.state"
    region = "us-east-2"
  }
}
provider "aws" {
  region = "us-east-2"
}
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