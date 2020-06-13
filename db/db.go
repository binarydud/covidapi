package db

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/binarydud/covidapi/types"
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

	svc := dynamodb.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
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

// GetStates ...
func (db *DB) GetStates() ([]types.State, error) {
	svc := db.svc
	var records []types.State

	// Use the ScanPages method to perform the scan with pagination. Use
	// just Scan method to make the API call without pagination.
	err := svc.ScanPages(&dynamodb.ScanInput{
		TableName: aws.String(stateTableName),
	}, func(page *dynamodb.ScanOutput, last bool) bool {
		recs := []types.State{}

		err := dynamodbattribute.UnmarshalListOfMaps(page.Items, &recs)
		if err != nil {
			panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
		}

		records = append(records, recs...)

		return true // keep paging
	})
	if err != nil {
		return records, err
	}
	return records, nil
}
