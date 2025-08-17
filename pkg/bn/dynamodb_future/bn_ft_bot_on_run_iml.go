package dynamodbfuture

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	dynamodbconfig "github.com/non26/tradepkg/pkg/bn/dynamodb_config"
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
	// if !result.IsFound() {
	// 	return nil, nil
	// }
	return &result, nil
}

func (d *bnFtBotOnRunRepository) Insert(ctx context.Context, botOnRun *models.BnFtBotOnRun) error {
	table := models.NewBinanceFutureBotOnRunTable(botOnRun)
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
	// table := models.NewBinanceFutureBotOnRunTable(botOnRun)
	// table.Transform()
	// updateExpression := fmt.Sprintf(
	// 	"set %v = :symbol, %v = :position_side, %v = :amount_q, %v = :is_active",
	// 	table.GetSymbolTableField(),
	// 	table.GetPositionSideTableField(),
	// 	table.GetAmountQtyTableField(),
	// 	table.GetIsActiveTableField(),
	// )
	// _, err := d.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
	// 	TableName:        aws.String(table.GetTableName()),
	// 	Key:              table.GetKeyBotIDAndOrderID(),
	// 	UpdateExpression: aws.String(updateExpression),
	// 	ExpressionAttributeValues: map[string]types.AttributeValue{
	// 		":symbol":        &types.AttributeValueMemberS{Value: table.Symbol},
	// 		":position_side": &types.AttributeValueMemberS{Value: table.PositionSide},
	// 		":amount_b":      &types.AttributeValueMemberS{Value: table.AmountB},
	// 		":is_active":     &types.AttributeValueMemberBOOL{Value: table.IsActive},
	// 	},
	// })
	// return err
	return nil
}

func (d *bnFtBotOnRunRepository) Delete(ctx context.Context, botOnRun *models.BnFtBotOnRun) error {
	table := models.NewBinanceFutureBotOnRunTable(botOnRun)
	_, err := d.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyBotIDAndOrderID(),
	})
	return err
}

func (d *bnFtBotOnRunRepository) Upsert(ctx context.Context, botOnRun *models.BnFtBotOnRun) error {
	table := models.NewBinanceFutureBotOnRunTable(botOnRun)
	update_config := dynamodbconfig.NewUpdateTable(botOnRun)
	update_config.Set(table.GetAccountIdTableField, botOnRun.AccountId)
	expression := update_config.BuildExpression()
	_, err := d.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 aws.String(table.GetTableName()),
		Key:                       table.GetKeyBotIDAndOrderID(),
		UpdateExpression:          expression,
		ExpressionAttributeValues: update_config.GetExpressionAttributeValues(),
	})
	return err
}

func (d *bnFtBotOnRunRepository) GetAll(ctx context.Context) ([]models.BnFtBotOnRun, error) {
	table := models.NewBinanceFutureBotOnRunTable(nil)
	response, err := d.client.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(table.GetTableName()),
	})
	if err != nil {
		return nil, err
	}
	items := []models.BnFtBotOnRun{}
	err = attributevalue.UnmarshalListOfMaps(response.Items, &items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (d *bnFtBotOnRunRepository) ScanWith(ctx context.Context, clientId string) ([]models.BnFtBotOnRun, error) {
	table := models.NewBinanceFutureBotOnRunTable(nil)
	response, err := d.client.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(table.GetTableName()),
		// Optional: Add a filter expression
		FilterExpression: aws.String("contains(#bot_order_id, :value)"),
		ExpressionAttributeNames: map[string]string{
			"#bot_order_id": "bot_order_id", // Field name in DynamoDB table
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":value": &types.AttributeValueMemberS{Value: clientId}, // Filter condition
		},
	})
	if err != nil {
		return nil, err
	}
	if len(response.Items) == 0 {
		return nil, nil
	}
	items := []models.BnFtBotOnRun{}
	err = attributevalue.UnmarshalListOfMaps(response.Items, &items)
	if err != nil {
		return nil, err
	}
	return items, nil
}
