package simplequeue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type credential struct {
	AccessKeyID     string
	SecretAccessKey string
}

func (c *credential) Retrieve(ctx context.Context) (aws.Credentials, error) {
	return aws.Credentials{
		AccessKeyID:     c.AccessKeyID,
		SecretAccessKey: c.SecretAccessKey,
	}, nil
}

func NewCredential(accessKeyID, secretAccessKey string) *credential {
	return &credential{
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
	}
}
