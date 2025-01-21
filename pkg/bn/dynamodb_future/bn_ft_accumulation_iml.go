package dynamodbfuture

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	models "github.com/non26/tradepkg/pkg/bn/dynamodb_future/models"
)

type bnFtAccumulationRepository struct {
	client *dynamodb.Client
}

func NewConnectionBnFtAccumulationRepository(client *dynamodb.Client) IBnFtAccumulationRepository {
	return &bnFtAccumulationRepository{
		client: client,
	}
}

func (d *bnFtAccumulationRepository) Get(ctx context.Context, accumulation *models.BnFtAccumulation) (*models.BnFtAccumulation, error) {
	table := models.NewBnFtAccumulationTableWith(accumulation)
	table.Transform()
	result := models.BnFtAccumulation{}
	response, err := d.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKey(),
	})
	if err != nil {
		return nil, err
	}
	if response.Item == nil {
		return models.NewBnFtAccumulationWith(&models.BnFtAccumulation{
			TotalAmountQ:   "0",
			CurrentAmountQ: "0",
		}), nil
	}

	err = attributevalue.UnmarshalMap(response.Item, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (d *bnFtAccumulationRepository) Delete(ctx context.Context, accumulation *models.BnFtAccumulation) error {
	table := models.NewBnFtAccumulationTableWith(accumulation)
	table.Transform()
	_, err := d.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKey(),
	})
	return err
}

func (d *bnFtAccumulationRepository) Insert(ctx context.Context, accumulation *models.BnFtAccumulation) error {
	table := models.NewBnFtAccumulationTableWith(accumulation)
	table.Transform()
	item, err := attributevalue.MarshalMap(table.GetItem())
	if err != nil {
		return err
	}
	_, err = d.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(table.GetTableName()),
		Item:      item,
	})
	if err != nil {
		return err
	}
	return nil
}

func (d *bnFtAccumulationRepository) Update(ctx context.Context, accumulation *models.BnFtAccumulation) error {
	table := models.NewBnFtAccumulationTableWith(accumulation)
	_, err := d.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:        aws.String(table.GetTableName()),
		Key:              table.GetKey(),
		UpdateExpression: aws.String("set TotalAmountQ = :totalAmountQ, CurrentAmountQ = :currentAmountQ"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":totalAmountQ":   &types.AttributeValueMemberN{Value: table.TotalAmountQ},
			":currentAmountQ": &types.AttributeValueMemberN{Value: table.CurrentAmountQ},
		},
	})
	return err
}
