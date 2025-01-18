package dynamodbrepository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	models "github.com/non26/tradepkg/pkg/bn/dynamodb_repository/models"
)

type bnFtAdvancedPositionRepository struct {
	client *dynamodb.Client
}

func NewConnectionBnFtAdvancedPositionRepository(client *dynamodb.Client) IBnFtAdvancedPositionRepository {
	return &bnFtAdvancedPositionRepository{
		client: client,
	}
}

func (d *bnFtAdvancedPositionRepository) Get(ctx context.Context, advancedPosition *models.BnFtAdvancedPositionModel) (*models.BnFtAdvancedPositionModel, error) {
	table := models.NewBinanceFutureAdvancedPositionTableWith(advancedPosition)
	table.Transform()
	table.SetCreatedAt()
	result := models.BnFtAdvancedPositionModel{}
	response, err := d.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKey(),
	})
	if err != nil {
		return nil, err
	}
	err = attributevalue.UnmarshalMap(response.Item, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (d *bnFtAdvancedPositionRepository) Insert(ctx context.Context, advancedPosition *models.BnFtAdvancedPositionModel) error {
	table := models.NewBinanceFutureAdvancedPositionTableWith(advancedPosition)
	table.Transform()
	table.SetCreatedAt()
	item, err := attributevalue.MarshalMap(advancedPosition)
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

func (d *bnFtAdvancedPositionRepository) Delete(ctx context.Context, advancedPosition *models.BnFtAdvancedPositionModel) error {
	table := models.NewBinanceFutureAdvancedPositionTableWith(advancedPosition)
	table.Transform()
	_, err := d.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKey(),
	})
	return err
}
