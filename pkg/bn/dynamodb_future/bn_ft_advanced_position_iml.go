package dynamodbfuture

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamodbconfig "github.com/non26/tradepkg/pkg/bn/dynamodb_config"
	models "github.com/non26/tradepkg/pkg/bn/dynamodb_future/models"
)

type bnFtAdvancedPositionRepository struct {
	client *dynamodb.Client
}

func NewConnectionBnFtAdvancedPositionRepository(client *dynamodb.Client) IBnFtAdvancedPositionRepository {
	return &bnFtAdvancedPositionRepository{
		client: client,
	}
}

func (d *bnFtAdvancedPositionRepository) Get(ctx context.Context, clientId string) (*models.BnFtAdvancedPosition, error) {
	table := models.NewBnFtAdvancedPositionTable()
	table.ClientID = clientId
	result := models.BnFtAdvancedPosition{}
	response, err := d.client.GetItem(ctx, &dynamodb.GetItemInput{
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
	return &result, nil
}

func (d *bnFtAdvancedPositionRepository) Delete(ctx context.Context, clientId string) error {
	table := models.NewBnFtAdvancedPositionTable()
	table.ClientID = clientId
	_, err := d.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyClientID(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (d *bnFtAdvancedPositionRepository) Upsert(ctx context.Context, advancedPosition *models.BnFtAdvancedPosition) error {
	table := models.NewBnFtAdvancedPositionTableWith(advancedPosition)
	table.Transform()
	update_config := dynamodbconfig.NewUpdateTable(advancedPosition)
	update_config.Set(table.GetAmountBTableField, advancedPosition.AmountB)
	update_config.Set(table.GetClientIDTableField, advancedPosition.ClientID)
	update_config.Set(table.GetPositionSideTableField, advancedPosition.PositionSide)
	update_config.Set(table.GetSideTableField, advancedPosition.Side)
	update_config.Set(table.GetSymbolTableField, advancedPosition.Symbol)
	expression := update_config.BuildExpression()
	_, err := d.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 aws.String(table.GetTableName()),
		Key:                       table.GetKeyClientID(),
		UpdateExpression:          expression,
		ExpressionAttributeValues: update_config.GetExpressionAttributeValues(),
	})
	if err != nil {
		return err
	}
	return nil
}
