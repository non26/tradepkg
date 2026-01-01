package simplequeue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type ISimpleQueueConfig interface {
	NewLocal() *sqs.Client
	NewPrd() *sqs.Client
}

type SimpleQueueConfig struct {
	cfg        aws.Config
	credential aws.CredentialsProvider
}

func NewSimpleQueueConfig(credential aws.CredentialsProvider) ISimpleQueueConfig {
	cfg, _ := config.LoadDefaultConfig(context.TODO())
	return &SimpleQueueConfig{cfg: cfg, credential: credential}
}

func (c *SimpleQueueConfig) NewLocal() *sqs.Client {
	return sqs.NewFromConfig(c.cfg, func(o *sqs.Options) {
		o.Credentials = c.credential
	})
}

func (c *SimpleQueueConfig) NewPrd() *sqs.Client {
	return sqs.NewFromConfig(c.cfg, func(o *sqs.Options) {
		o.Credentials = c.credential
	})
}
