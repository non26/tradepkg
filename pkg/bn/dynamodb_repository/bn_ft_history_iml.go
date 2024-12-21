package dynamodbrepository

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	models "github.com/non26/tradepkg/pkg/bn/dynamodb_repository/models"
)

func (d *dynamoDBRepository) GetAllHistory(ctx context.Context) ([]models.BnFtHistory, error) {
	table := models.NewBinanceFutureHistoryTable()
	result := []models.BnFtHistory{}
	response, err := d.dynamodb.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(table.GetTableName()),
	})
	if err != nil {
		return nil, err
	}

	err = attributevalue.UnmarshalListOfMaps(response.Items, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *dynamoDBRepository) GetHistoryByClientID(ctx context.Context, clientId string) (*models.BnFtHistory, error) {
	table := models.NewBinanceFutureHistoryTable()
	table.ClientId = clientId
	result := &models.BnFtHistory{}
	response, err := d.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyClientID(),
	})
	if err != nil {
		return nil, err
	}

	err = attributevalue.UnmarshalMap(response.Item, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *dynamoDBRepository) InsertHistory(ctx context.Context, history *models.BnFtHistory) error {
	table := models.NewBinanceFutureHistoryTable()
	item, err := attributevalue.MarshalMap(history)
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

func (d *dynamoDBRepository) UpdateHistory(ctx context.Context, history *models.BnFtHistory) error {
	table := models.NewBinanceFutureHistoryTable()
	item, err := attributevalue.MarshalMap(history)
	if err != nil {
		return err
	}
	_, err = d.dynamodb.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyClientID(),
		UpdateExpression: aws.String(fmt.Sprintf(
			"set %v = :history_id, %v = :symbol, %v = :position_side",
			table.GetHistoryIdTableField(),
			table.GetSymbolTableField(),
			table.GetPositionSideTableField(),
		)),
		ExpressionAttributeValues: item,
	})
	return err
}
