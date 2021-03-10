package db

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var dynamodbClient *dynamodb.DynamoDB
var tableName = "FavoriteBusStop"

// struct representing the db table
type FavoriteBusStop struct {
	UserId  string
	BusStop string
}

// Permissions must be granted to IAM user for DynamoDB
// and database should be created manually first
func InitializeDynamodbClient() {
	log.Printf("InitializeDynamodbClient")
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	dynamodbClient = dynamodb.New(session)
	log.Printf("InitializeDynamodbClient done")
}

// Get favorite bus stop from database
func GetFavoriteBusStop(userId string) (*FavoriteBusStop, error) {
	result, err := dynamodbClient.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"UserId": {
				S: aws.String(userId),
			},
		},
		TableName: aws.String(tableName),
	})

	if err != nil {
		log.Printf("Something went wrong while getting item from table %s: %s", tableName, err.Error())
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	favoriteBusStop := FavoriteBusStop{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &favoriteBusStop)
	if err != nil {
		log.Printf("Failed to unmarshal record, %s", err.Error())
		return nil, err
	}

	return &favoriteBusStop, nil
}

// Add favorite bus stop to database
func AddFavoriteBusStop(userId string, busStop string) error {
	fav := FavoriteBusStop{UserId: userId, BusStop: busStop}
	data, err := dynamodbattribute.MarshalMap(fav)
	if err != nil {
		log.Printf("Failed to marshal record, %s", err.Error())
		return err
	}

	_, err = dynamodbClient.PutItem(&dynamodb.PutItemInput{
		Item:      data,
		TableName: aws.String(tableName),
	})

	if err != nil {
		log.Printf("Something went wrong while adding item to table %s: %s", tableName, err.Error())
		return err
	}

	return nil
}

// Delete favorite bus stop to database
// You must make sure item is stored in database before deleting it
func DeleteFavoriteBusStop(userId string) error {
	_, err := dynamodbClient.DeleteItem(&dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"UserId": {
				S: aws.String(userId),
			},
		},
		TableName: aws.String(tableName),
	})

	if err != nil {
		log.Printf("Something went wrong while deleting item to table %s: %s", tableName, err.Error())
		return err
	}

	return nil
}
