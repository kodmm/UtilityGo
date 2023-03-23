package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Item struct {
	ID            string `dynamodbav:"id"`
	ProcessDate   string `dynamodbav:"process_date"`
	Text          string `dynamodbav:"text"`
	TextOmitEmpty string `dynamodbav:"text_omit_empty,omitempty"`
}

func main() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("failed to load SDK configuration, %v", err)
	}
	db := dynamodb.NewFromConfig(cfg)

	item := Item{
		ID:            "0001",
		ProcessDate:   time.Now().Format("2006-01-02"),
		Text:          "example text",
		TextOmitEmpty: "",
	}
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		log.Fatalf("failed to marshal, item = %v, %v", item, err)
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String("example"),
		Item:      av,
	}

	_, err = db.PutItem(ctx, input)
	if err != nil {
		log.Fatalf("failed to put item, %v", err)
	}

	inputGet := &dynamodb.GetItemInput{
		TableName: aws.String("example"),
		Key: map[string]types.AttributeValue{
			"id":           &types.AttributeValueMemberS{Value: "00001"},
			"process_date": &types.AttributeValueMemberS{Value: "2020-11-01"},
		},
	}
	out, err := db.GetItem(ctx, inputGet)
	if err != nil {
		// error handling
	}
	var itemGet Item
	err = attributevalue.UnmarshalMap(out.Item, &itemGet)
	if err != nil {
		// error handling
	}
	fmt.Printf("get item: %+v", item)

	inputGet2 := &dynamodb.QueryInput{
		TableName: aws.String("example"),
		ExpressionAttributeNames: map[string]string{
			"#ID": "id",
		},
		ExpressionAttributeValues: map[string]string{
			":id":    &types.AttributeValueMemberS{Value: "00001"},
			":month": &types.AttributeValueMemberS{Value: "2020-10"},
		},
		KeyConditionExpression: aws.String("#ID = :id and begins_with(process_date, :month)"),
	}
	out2, err := db.Query(ctx, inputGet2)
	if err != nil {
		// error handling
	}
	var items2 []Item
	err = attributevalue.UnmarshalListOfMaps(out.Item, &items2)
	if err != nil {
		// error handling
	}
	for _, v := range items2 {
		fmt.Printf("query: %+v\n", v)
	}
}
