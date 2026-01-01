package simplequeue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type ISimpleQueue interface {
	Send(ctx context.Context, item string) error
}

type SimpleQueue struct {
	url string
	sqs *sqs.Client
}

func NewSimpleQueue(
	url string,
	sqs *sqs.Client,
) ISimpleQueue {
	return &SimpleQueue{
		url: url,
		sqs: sqs,
	}
}

func (q *SimpleQueue) Send(ctx context.Context, item string) error {
	_, err := q.sqs.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    &q.url,
		MessageBody: &item,
	})
	if err != nil {
		return err
	}

	return nil
}
