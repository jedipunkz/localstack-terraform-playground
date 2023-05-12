package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct {
	ID   string `json:"id"`
	Data string `json:"data"`
}

type DynamoDBClient struct {
	svc *dynamodb.DynamoDB
}

func main() {
	sess, err := createSession("http://localhost:4566", "us-east-1")
	if err != nil {
		fmt.Println("Error creating session,", err)
		return
	}

	client := DynamoDBClient{svc: dynamodb.New(sess)}

	tableName := "SampleTable"
	err = client.createTable(tableName)
	if err != nil {
		fmt.Println("Error creating table,", err)
	}

	item := Item{
		ID:   "1",
		Data: "Sample Data",
	}
	err = client.putItem(tableName, item)
	if err != nil {
		fmt.Println("Error putting item,", err)
		return
	}

	items, err := client.getAllItems(tableName)
	if err != nil {
		fmt.Println("Error getting items,", err)
		return
	}

	fmt.Println("Items:")
	for _, item := range items {
		fmt.Println(item)
	}
}

func createSession(endpoint, region string) (*session.Session, error) {
	return session.NewSession(&aws.Config{
		Region:   aws.String(region),
		Endpoint: aws.String(endpoint),
	})
}

func (client *DynamoDBClient) createTable(tableName string) error {
	_, err := client.svc.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	})

	// If table does not exist, create it
	if err != nil {
		_, err = client.svc.CreateTable(&dynamodb.CreateTableInput{
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String("id"),
					AttributeType: aws.String("S"),
				},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("id"),
					KeyType:       aws.String("HASH"),
				},
			},
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(5),
				WriteCapacityUnits: aws.Int64(5),
			},
			TableName: aws.String(tableName),
		})
	}

	return err
}

func (client *DynamoDBClient) putItem(tableName string, item Item) error {
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}

	_, err = client.svc.PutItem(&dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	})

	return err
}

func (client *DynamoDBClient) getAllItems(tableName string) ([]Item, error) {
	result, err := client.svc.Scan(&dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})
	if err != nil {
		return nil, err
	}

	items := []Item{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &items)

	return items, err
}
