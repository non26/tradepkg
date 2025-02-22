package dynamodbconfig

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type newDynamodb struct {
	endPoint   dynamodb.EndpointResolverV2
	credential aws.CredentialsProvider
	awsconfig  aws.Config
}

func (c newDynamodb) NewLocal() *dynamodb.Client {
	svc := dynamodb.NewFromConfig(c.awsconfig, func(o *dynamodb.Options) {
		o.Credentials = c.credential
		o.EndpointResolverV2 = c.endPoint
	})
	return svc
}

func (c newDynamodb) NewPrd() *dynamodb.Client {
	svc := dynamodb.NewFromConfig(c.awsconfig, func(o *dynamodb.Options) {
		o.Credentials = c.credential
	})
	return svc
}

func DynamoDB(
	endPoint dynamodb.EndpointResolverV2,
	credential aws.CredentialsProvider,
	awsconfig aws.Config,
) *newDynamodb {
	return &newDynamodb{
		endPoint,
		credential,
		awsconfig,
	}
}
