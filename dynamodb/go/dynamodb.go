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

func main() {
	sess, err := createSession("http://localhost:4566", "us-east-1")
	if err != nil {
		fmt.Println("Error creating session,", err)
		return
	}

	svc := dynamodb.New(sess)

	tableName := "SampleTable"
	err = createTable(svc, tableName)
	if err != nil {
		fmt.Println("Error creating table,", err)
	}

	item := Item{
		ID:   "1",
		Data: "Sample Data",
	}
	err = putItem(svc, tableName, item)
	if err != nil {
		fmt.Println("Error putting item,", err)
		return
	}

	itemData, err := getItem(svc, tableName, "1")
	if err != nil {
		fmt.Println("Error getting item,", err)
		return
	}

	fmt.Println("Item data:", itemData)

	err = listTables(svc)
	if err != nil {
		fmt.Println("Error listing tables,", err)
		return
	}
}

func createSession(endpoint, region string) (*session.Session, error) {
	return session.NewSession(&aws.Config{
		Region:   aws.String(region),
		Endpoint: aws.String(endpoint),
	})
}

func createTable(svc *dynamodb.DynamoDB, tableName string) error {
	_, err := svc.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	})

	// If table does not exist, create it
	if err != nil {
		_, err = svc.CreateTable(&dynamodb.CreateTableInput{
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

func putItem(svc *dynamodb.DynamoDB, tableName string, item Item) error {
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}

	_, err = svc.PutItem(&dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	})

	return err
}

func getItem(svc *dynamodb.DynamoDB, tableName, id string) (Item, error) {
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		return Item{}, err
	}

	item := Item{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	return item, err
}

func listTables(svc *dynamodb.DynamoDB) error {
	result, err := svc.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		return err
	}

	fmt.Println("Tables:")
	for _, table := range result.TableNames {
		fmt.Println(*table)
	}

	return nil
}
