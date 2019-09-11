package main

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/kelseyhightower/envconfig"
)

// DynamoConfig aws dynamoDB config
type DynamoConfig struct {
	Region string `default:"ap-northeast-1"` // AWS_DYNAMO_REGION
	Table  string `default:"sns_alert"`      // AWS_DYNAMO_TABLE
}

type alertTable struct {
	Timestamp time.Time `dynamo:"timestamp"`
	Event     string    `dynamo:"event"`
	Message   string    `dynamo:"message"`
}

func getItemCount() (int, error) {
	var dynamoConf DynamoConfig
	if err := envconfig.Process("aws_dynamo", &dynamoConf); err != nil {
		return 0, err
	}
	table := setupDB(dynamoConf.Region, dynamoConf.Table)

	var results []alertTable
	if err := table.Scan().All(&results); err != nil {
		return 0, err
	}
	return len(results), nil
}

func setupDB(region, tableNm string) dynamo.Table {
	db := dynamo.New(session.New(), &aws.Config{Region: aws.String(region)})
	table := db.Table(tableNm)
	return table
}
