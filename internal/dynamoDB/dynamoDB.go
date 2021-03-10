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

func AddFavoriteBusStop(userId string, busStop string) error {
	fav := FavoriteBusStop{UserId: userId, BusStop: busStop}
	data, err := dynamodbattribute.MarshalMap(fav)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      data,
		TableName: aws.String(tableName),
	}

	_, err = dynamodbClient.PutItem(input)
	if err != nil {
		log.Fatalf("Something went wrong when adding item to table %s: %s", tableName, err.Error())
		return err
	}

	return nil
}
