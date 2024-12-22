package dynamodbrepository

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	models "github.com/non26/tradepkg/pkg/bn/dynamodb_repository/models"
)

func (d *dynamoDBRepository) GetBotOnRunByBotIDAndOrderID(ctx context.Context, botOnRun *models.BnFtBotOnRun) (*models.BnFtBotOnRun, error) {
	table := models.NewBinanceFutureBotOnRunTable(botOnRun)
	result := models.BnFtBotOnRun{}
	response, err := d.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyBotIDAndOrderID(),
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

func (d *dynamoDBRepository) InsertBotOnRun(ctx context.Context, botOnRun *models.BnFtBotOnRun) error {
	table := models.NewBinanceFutureBotOnRunTable(botOnRun)
	item, err := attributevalue.MarshalMap(botOnRun)
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

func (d *dynamoDBRepository) UpdateBotOnRun(ctx context.Context, botOnRun *models.BnFtBotOnRun) error {
	table := models.NewBinanceFutureBotOnRunTable(botOnRun)
	item, err := attributevalue.MarshalMap(botOnRun)
	if err != nil {
		return err
	}
	_, err = d.dynamodb.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyBotID(),
		UpdateExpression: aws.String(fmt.Sprintf(
			"set %v = :symbol, %v = :position_side, %v = :position_condition",
			table.GetSymbolTableField(),
			table.GetPositionSideTableField(),
			table.GetPositionConditionTableField(),
		)),
		ExpressionAttributeValues: item,
	})
	return err
}

func (d *dynamoDBRepository) DeleteBotOnRun(ctx context.Context, botOnRun *models.BnFtBotOnRun) error {
	table := models.NewBinanceFutureBotOnRunTable(botOnRun)
	_, err := d.dynamodb.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyBotIDAndOrderID(),
	})
	return err
}
