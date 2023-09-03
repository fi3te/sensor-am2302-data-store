package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func createDynamoDbClient(ctx context.Context) dynamodb.Client {
	sdkConfig, err := awsConfig.LoadDefaultConfig(ctx)
	if err != nil {
		panic(err)
	}
	return *dynamodb.NewFromConfig(sdkConfig)
}

func putItem(ctx context.Context, db *dynamodb.Client, tableName string, dataPoint DataPoint) error {
	item, err := attributevalue.MarshalMap(dataPoint)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	}
	_, err = db.PutItem(ctx, input)
	return err
}
