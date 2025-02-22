package dynamodbconfig

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type dynamodbConfig struct{}

func (d *dynamodbConfig) LoadConfig() aws.Config {
	cfg, _ := config.LoadDefaultConfig(context.TODO())
	return cfg
}
func NewDynamodbConfig() *dynamodbConfig {
	return &dynamodbConfig{}
}
