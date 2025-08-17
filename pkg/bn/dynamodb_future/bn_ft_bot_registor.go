package dynamodbfuture

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamodbconfig "github.com/non26/tradepkg/pkg/bn/dynamodb_config"
	models "github.com/non26/tradepkg/pkg/bn/dynamodb_future/models"
)

type bnFtBotRegistorRepository struct {
	client *dynamodb.Client
}

func NewConnectionBnFtBotRegistorRepository(client *dynamodb.Client) IBnFtBotRegistorRepository {
	return &bnFtBotRegistorRepository{
		client: client,
	}
}

func (d *bnFtBotRegistorRepository) Get(ctx context.Context, botID string, botOrderID string) (*models.BnFtBotRegistor, error) {
	registor := models.NewBnFtBotRegistor()
	registor.BotID = botID
	registor.BotOrderID = botOrderID
	table := models.NewBnFtBotRegistorTable(registor)
	result := models.BnFtBotRegistor{}
	response, err := d.client.GetItem(ctx, &dynamodb.GetItemInput{
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

func (d *bnFtBotRegistorRepository) Upsert(ctx context.Context, botRegistor *models.BnFtBotRegistor) error {
	table := models.NewBnFtBotRegistorTable(botRegistor)
	update_config := dynamodbconfig.NewUpdateTable(botRegistor)
	update_config.Set(table.GetIsActiveTableField, botRegistor.IsActive)
	update_config.Set(table.GetSymbolTableField, botRegistor.Symbol)
	update_config.Set(table.GetPositionSideTableField, botRegistor.PositionSide)
	update_config.Set(table.GetAmountQtyTableField, botRegistor.AmountQ)
	update_config.Set(table.GetAccountIdTableField, botRegistor.AccountId)
	update_config.Set(table.GetSettingTableField, botRegistor.Setting)
	expression := update_config.BuildExpression()
	_, err := d.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 aws.String(table.GetTableName()),
		Key:                       table.GetKeyBotIDAndOrderID(),
		UpdateExpression:          expression,
		ExpressionAttributeValues: update_config.GetExpressionAttributeValues(),
	})
	return err
}

func (d *bnFtBotRegistorRepository) GetAll(ctx context.Context) ([]models.BnFtBotRegistor, error) {
	table := models.NewBnFtBotRegistorTable(nil)
	response, err := d.client.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(table.GetTableName()),
	})
	if err != nil {
		return nil, err
	}
	items := []models.BnFtBotRegistor{}
	err = attributevalue.UnmarshalListOfMaps(response.Items, &items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (d *bnFtBotRegistorRepository) Delete(ctx context.Context, botID string, botOrderID string) error {
	table := models.NewBnFtBotRegistorTable(&models.BnFtBotRegistor{
		BotID:      botID,
		BotOrderID: botOrderID,
	})
	_, err := d.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyBotIDAndOrderID(),
	})
	if err != nil {
		return err
	}
	return nil
}
