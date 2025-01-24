package dynamodbfuture

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	models "github.com/non26/tradepkg/pkg/bn/dynamodb_future/models"
)

type bnFtBotOnRunRepository struct {
	client *dynamodb.Client
}

func NewConnectionBnFtBotOnRunRepository(client *dynamodb.Client) IBnFtBotOnRunRepository {
	return &bnFtBotOnRunRepository{
		client: client,
	}
}

func (d *bnFtBotOnRunRepository) Get(ctx context.Context, botOnRun *models.BnFtBotOnRun) (*models.BnFtBotOnRun, error) {
	table := models.NewBinanceFutureBotOnRunTable(botOnRun)
	result := models.BnFtBotOnRun{}
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

func (d *bnFtBotOnRunRepository) Insert(ctx context.Context, botOnRun *models.BnFtBotOnRun) error {
	table := models.NewBinanceFutureBotOnRunTable(botOnRun)
	table.Transform()
	item, err := attributevalue.MarshalMap(table.GetData())
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

func (d *bnFtBotOnRunRepository) Update(ctx context.Context, botOnRun *models.BnFtBotOnRun) error {
	table := models.NewBinanceFutureBotOnRunTable(botOnRun)
	table.Transform()
	updateExpression := fmt.Sprintf(
		"set %v = :symbol, %v = :position_side, %v = :amount_q, %v = :is_active",
		table.GetSymbolTableField(),
		table.GetPositionSideTableField(),
		table.GetAmountQtyTableField(),
		table.GetIsActiveTableField(),
	)
	_, err := d.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:        aws.String(table.GetTableName()),
		Key:              table.GetKeyBotIDAndOrderID(),
		UpdateExpression: aws.String(updateExpression),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":symbol":        &types.AttributeValueMemberS{Value: table.Symbol},
			":position_side": &types.AttributeValueMemberS{Value: table.PositionSide},
			":amount_b":      &types.AttributeValueMemberS{Value: table.AmountB},
			":is_active":     &types.AttributeValueMemberBOOL{Value: table.IsActive},
		},
	})
	return err
}

func (d *bnFtBotOnRunRepository) Delete(ctx context.Context, botOnRun *models.BnFtBotOnRun) error {
	table := models.NewBinanceFutureBotOnRunTable(botOnRun)
	_, err := d.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyBotIDAndOrderID(),
	})
	return err
}
