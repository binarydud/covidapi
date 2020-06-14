package db

import (
	"fmt"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/binarydud/covidapi/types"
	"github.com/rs/zerolog/log"
)

// DB dynamo wrapper
type DB struct {
	svc dynamodbiface.DynamoDBAPI
}

const stateTableName = "CovidState"
const usTableName = "CovidUS"

// New reates new dynamo struct
func New() *DB {
	mySession := session.Must(session.NewSession())

	svc := dynamodb.New(mySession, aws.NewConfig().WithRegion("us-east-2"))
	db := &DB{
		svc: svc,
	}
	return db
}

// PutState creates a new state record
func (db *DB) PutState(state types.State) error {
	svc := db.svc

	av, err := dynamodbattribute.MarshalMap(state)
	if err != nil {
		return err
	}
	_, err = svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(stateTableName),
		Item:      av,
	})
	return err
}

// PutUS ...
func (db *DB) PutUS(us types.US) error {
	svc := db.svc
	av, err := dynamodbattribute.MarshalMap(us)
	if err != nil {
		return err
	}
	_, err = svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(usTableName),
		Item:      av,
	})
	return err
}

// GetUSHistorical ...
func (db *DB) GetUSHistorical() ([]types.US, error) {
	svc := db.svc
	var records []types.US

	// Use the ScanPages method to perform the scan with pagination. Use
	// just Scan method to make the API call without pagination.
	err := svc.ScanPages(&dynamodb.ScanInput{
		TableName: aws.String(usTableName),
	}, func(page *dynamodb.ScanOutput, last bool) bool {
		recs := []types.US{}

		err := dynamodbattribute.UnmarshalListOfMaps(page.Items, &recs)
		if err != nil {
			panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
		}

		records = append(records, recs...)

		return true // keep paging
	})

	if err != nil {
		return nil, err
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].Date < records[j].Date
	})
	return records, nil
}

// GetUSCurrent ...
func (db *DB) GetUSCurrent() (*types.US, error) {
	records, err := db.GetUSHistorical()
	if err != nil {
		return nil, err
	}
	if len(records) < 1 {
		return nil, nil
	}
	state := records[len(records)-1]
	return &state, nil
}

// GetStateCurrent ...
func (db *DB) GetStateCurrent(name string) (*types.State, error) {
	states, err := db.GetStateHistorical(name)
	if err != nil {
		return nil, err
	}
	if len(states) < 1 {
		return nil, nil
	}
	state := states[len(states)-1]
	return &state, nil
}

// GetStateHistorical ...
func (db *DB) GetStateHistorical(name string) ([]types.State, error) {
	svc := db.svc
	keyCond := expression.Key("state").Equal(expression.Value(name))
	builder := expression.NewBuilder().WithKeyCondition(keyCond)
	expr, err := builder.Build()
	params := &dynamodb.QueryInput{
		TableName:                 aws.String(stateTableName),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}
	records := make([]types.State, 0)
	err = svc.QueryPages(params, func(page *dynamodb.QueryOutput, lastPage bool) bool {
		log.Debug().Msg("getting page")
		recs := []types.State{}

		err := dynamodbattribute.UnmarshalListOfMaps(page.Items, &recs)
		if err != nil {
			panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
		}

		records = append(records, recs...)
		return true
	})
	if err != nil {
		return nil, err
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].Date < records[j].Date
	})
	return records, nil

}
