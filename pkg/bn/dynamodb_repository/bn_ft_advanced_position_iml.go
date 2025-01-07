package dynamodbrepository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	models "github.com/non26/tradepkg/pkg/bn/dynamodb_repository/models"
)

func (d *dynamoDBRepository) GetAdvancedPositionBySymbolAndPositionSide(ctx context.Context, advancedPosition *models.BnFtAdvancedPositionModel) (*models.BnFtAdvancedPositionModel, error) {
	table := models.NewBinanceFutureAdvancedPositionTableWithData(advancedPosition)
	result := models.BnFtAdvancedPositionModel{}
	response, err := d.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
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

func (d *dynamoDBRepository) InsertAdvancedPosition(ctx context.Context, advancedPosition *models.BnFtAdvancedPositionModel) error {
	table := models.NewBinanceFutureAdvancedPositionTableWithData(advancedPosition)
	item, err := attributevalue.MarshalMap(advancedPosition)
	if err != nil {
		return err
	}
	_, err = d.dynamodb.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(table.GetTableName()),
		Item:      item,
	})
	if err != nil {
		return err
	}
	return nil
}

func (d *dynamoDBRepository) DeleteAdvancedPosition(ctx context.Context, advancedPosition *models.BnFtAdvancedPositionModel) error {
	table := models.NewBinanceFutureAdvancedPositionTableWithData(advancedPosition)
	_, err := d.dynamodb.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKey(),
	})
	return err
}
