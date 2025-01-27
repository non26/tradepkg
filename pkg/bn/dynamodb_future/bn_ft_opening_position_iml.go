package dynamodbfuture

import (
	"context"
	"fmt"
	"log"

	models "github.com/non26/tradepkg/pkg/bn/dynamodb_future/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type bnFtOpeningPositionRepository struct {
	client *dynamodb.Client
}

func NewConnectionBnFtOpeningPositionRepository(client *dynamodb.Client) IBnFtOpeningPositionRepository {
	return &bnFtOpeningPositionRepository{
		client: client,
	}
}

// return map[string]*models.BnFtOpeningPosition all position in db
// key is symbol + position_side
func (d *bnFtOpeningPositionRepository) GetAll(ctx context.Context) (map[string]*models.BnFtOpeningPosition, error) {
	var err error
	var response *dynamodb.ScanOutput
	table := models.NewBinanceFutureOpeningPositionTable()
	result := make([]models.BnFtOpeningPosition, 0)
	response, err = d.client.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(table.GetTableName()),
	})
	if err != nil {
		return nil, err
	}
	openOrders := make(map[string]*models.BnFtOpeningPosition)
	if len(response.Items) == 0 {
		return openOrders, nil
	}

	err = attributevalue.UnmarshalListOfMaps(response.Items, &result)
	if err != nil {
		return nil, err
	}

	for _, openOrder := range result {
		openOrders[openOrder.Symbol+openOrder.PositionSide] = &openOrder
	}
	return openOrders, nil
}

func (d *bnFtOpeningPositionRepository) Get(ctx context.Context, data *models.BnFtOpeningPosition) (*models.BnFtOpeningPosition, error) {
	var err error
	var response *dynamodb.GetItemOutput
	result := &models.BnFtOpeningPosition{}
	table := models.NewBinanceFutureOpeningPositionTableWith(data)
	response, err = d.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyByPositionSideAndSymbol(),
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

func (d *bnFtOpeningPositionRepository) ScanWith(ctx context.Context, client_id string) (*models.BnFtOpeningPosition, error) {
	var err error
	var response *dynamodb.ScanOutput
	result := []models.BnFtOpeningPosition{}
	table := models.NewBinanceFutureOpeningPositionTable()
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
		return &models.BnFtOpeningPosition{}, nil
	}

	err = attributevalue.UnmarshalListOfMaps(response.Items, &result)
	if err != nil {
		log.Fatalf("Failed to unmarshal items: %v", err)
		return nil, err
	}
	if len(result) == 0 {
		return &models.BnFtOpeningPosition{}, nil
	}
	return &result[0], nil
}

func (d *bnFtOpeningPositionRepository) Delete(ctx context.Context, openOrder *models.BnFtOpeningPosition) error {
	table := models.NewBinanceFutureOpeningPositionTable()
	table.BnFtOpeningPosition = openOrder
	_, err := d.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyByPositionSideAndSymbol(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (d *bnFtOpeningPositionRepository) Insert(ctx context.Context, openOrder *models.BnFtOpeningPosition) error {
	table := models.NewBinanceFutureOpeningPositionTableWith(openOrder)
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

func (d *bnFtOpeningPositionRepository) UpdateAmountB(ctx context.Context, openOrder *models.BnFtOpeningPosition) error {
	table := models.NewBinanceFutureOpeningPositionTableWith(openOrder)
	table.Transform()
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyByPositionSideAndSymbol(),
		UpdateExpression: aws.String(fmt.Sprintf(
			"set %v = :amount_b",
			table.GetAmountBTableField(),
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
