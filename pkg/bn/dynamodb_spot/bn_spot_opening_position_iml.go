package dynamodbspot

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	dynamodbconfig "github.com/non26/tradepkg/pkg/bn/dynamodb_config"
	models "github.com/non26/tradepkg/pkg/bn/dynamodb_spot/models"
)

type bnSpotOpeningPositionRepository struct {
	client *dynamodb.Client
}

func NewConnectionBnSpotOpeningPositionRepository(client *dynamodb.Client) IBnSpotOpeningPositionRepository {
	return &bnSpotOpeningPositionRepository{
		client: client,
	}
}

// return map[string]*models.BnFtOpeningPosition all position in db
// key is symbol + position_side
func (d *bnSpotOpeningPositionRepository) GetAll(ctx context.Context) (map[string]*models.BnSpotOpeningPosition, error) {
	var err error
	var response *dynamodb.ScanOutput
	table := models.NewBinanceSpotOpeningPositionTable()
	result := make([]models.BnSpotOpeningPosition, 0)
	response, err = d.client.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(table.GetTableName()),
	})
	if err != nil {
		return nil, err
	}
	openOrders := make(map[string]*models.BnSpotOpeningPosition)
	if len(response.Items) == 0 {
		return openOrders, nil
	}

	err = attributevalue.UnmarshalListOfMaps(response.Items, &result)
	if err != nil {
		return nil, err
	}

	for _, openOrder := range result {
		openOrders[openOrder.Symbol] = &openOrder
	}
	return openOrders, nil
}

func (d *bnSpotOpeningPositionRepository) Get(ctx context.Context, data *models.BnSpotOpeningPosition) (*models.BnSpotOpeningPosition, error) {
	var err error
	var response *dynamodb.GetItemOutput
	result := &models.BnSpotOpeningPosition{}
	table := models.NewBinanceSpotOpeningPositionTableWith(data)
	response, err = d.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyBySymbol(),
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

func (d *bnSpotOpeningPositionRepository) ScanWith(ctx context.Context, client_id string) (*models.BnSpotOpeningPosition, error) {
	var err error
	var response *dynamodb.ScanOutput
	result := []models.BnSpotOpeningPosition{}
	table := models.NewBinanceSpotOpeningPositionTable()
	table.ClientId = client_id
	response, err = d.client.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(table.GetTableName()),
		// Optional: Add a filter expression
		FilterExpression: aws.String("contains(#client_id, :value)"),
		ExpressionAttributeNames: map[string]string{
			"#client_id": "client_id", // Field name in DynamoDB table
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":value": &types.AttributeValueMemberS{Value: client_id}, // Filter condition
		},
	})
	if err != nil {
		log.Fatalf("Failed to perform scan: %v", err)
		return nil, err
	}

	if len(response.Items) == 0 {
		return &models.BnSpotOpeningPosition{}, nil
	}

	err = attributevalue.UnmarshalListOfMaps(response.Items, &result)
	if err != nil {
		log.Fatalf("Failed to unmarshal items: %v", err)
		return nil, err
	}
	if len(result) == 0 {
		return &models.BnSpotOpeningPosition{}, nil
	}
	return &result[0], nil
}

func (d *bnSpotOpeningPositionRepository) Delete(ctx context.Context, openOrder *models.BnSpotOpeningPosition) error {
	table := models.NewBinanceSpotOpeningPositionTable()
	table.BnSpotOpeningPosition = openOrder
	_, err := d.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyBySymbol(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (d *bnSpotOpeningPositionRepository) Insert(ctx context.Context, openOrder *models.BnSpotOpeningPosition) error {
	table := models.NewBinanceSpotOpeningPositionTableWith(openOrder)
	table.Transform()
	table.SetCreatedAt()
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

func (d *bnSpotOpeningPositionRepository) UpdateAmountB(ctx context.Context, openOrder *models.BnSpotOpeningPosition) error {
	table := models.NewBinanceSpotOpeningPositionTableWith(openOrder)
	table.Transform()
	amountB, _ := table.GetAmountBTableField()
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyBySymbol(),
		UpdateExpression: aws.String(fmt.Sprintf(
			"set %v = :amount_b",
			amountB,
		)),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":amount_b": &types.AttributeValueMemberS{Value: openOrder.AmountB},
		},
	}
	_, err := d.client.UpdateItem(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (d *bnSpotOpeningPositionRepository) Upsert(ctx context.Context, openOrder *models.BnSpotOpeningPosition) error {
	table := models.NewBinanceSpotOpeningPositionTableWith(openOrder)
	table.SetCreatedAt()
	update_config := dynamodbconfig.NewUpdateTable(openOrder)
	update_config.Set(table.GetAmountBTableField, openOrder.AmountB)
	update_config.Set(table.GetCreatedAtTableField, openOrder.CreatedAt)
	update_config.Set(table.GetClientIdTableField, openOrder.ClientId)
	expression := update_config.BuildExpression()
	_, err := d.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 aws.String(table.GetTableName()),
		Key:                       table.GetKeyBySymbol(),
		UpdateExpression:          expression,
		ExpressionAttributeValues: update_config.GetExpressionAttributeValues(),
	})
	if err != nil {
		return err
	}
	return nil
}
