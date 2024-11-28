package dynamodbrepository

import (
	"context"

	models "github.com/non26/tradepkg/pkg/bn/dynamodb_repository/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (d *dynamoDBRepository) GetAllOpenOrders(ctx context.Context) ([]models.BinanceFutureOpeningPosition, error) {
	var err error
	var response *dynamodb.QueryOutput
	table := models.NewBinanceFutureOpeningPositionTable()
	result := make([]models.BinanceFutureOpeningPosition, 0)
	queryPaginator := dynamodb.NewQueryPaginator(d.dynamodb, &dynamodb.QueryInput{
		TableName: aws.String(table.GetTableName()),
	})
	for queryPaginator.HasMorePages() {
		response, err = queryPaginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		err = attributevalue.UnmarshalListOfMaps(response.Items, &result)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (d *dynamoDBRepository) GetOpenOrderBySymbol(ctx context.Context, symbol string) (*models.BinanceFutureOpeningPosition, error) {
	var err error
	var response *dynamodb.GetItemOutput
	result := &models.BinanceFutureOpeningPosition{}
	table := models.NewBinanceFutureOpeningPositionTable()
	table.Symbol = symbol
	response, err = d.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
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

func (d *dynamoDBRepository) GetOpenOrderByClientID(ctx context.Context, client_id string) (*models.BinanceFutureOpeningPosition, error) {
	var err error
	var response *dynamodb.GetItemOutput
	result := &models.BinanceFutureOpeningPosition{}
	table := models.NewBinanceFutureOpeningPositionTable()
	table.ClientId = client_id
	response, err = d.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyByClientID(),
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

func (d *dynamoDBRepository) DeleteOpenOrderBySymbol(ctx context.Context, symbol string) error {
	table := models.NewBinanceFutureOpeningPositionTable()
	table.Symbol = symbol
	_, err := d.dynamodb.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyBySymbol(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (d *dynamoDBRepository) NewOpenOrder(ctx context.Context, openOrder *models.BinanceFutureOpeningPosition) error {
	table := models.NewBinanceFutureOpeningPositionTable()
	table.BinanceFutureOpeningPosition = openOrder
	item, err := attributevalue.MarshalMap(table.BinanceFutureOpeningPosition)
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
