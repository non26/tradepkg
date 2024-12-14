package dynamodbrepository

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	models "github.com/non26/tradepkg/pkg/bn/dynamodb_repository/models"
)

func (d *dynamoDBRepository) GetQouteUSDT(ctx context.Context, symbol string) (*models.BinanceFutureQouteUSDT, error) {
	var err error
	var response *dynamodb.GetItemOutput
	result := &models.BinanceFutureQouteUSDT{}
	table := models.NewBinanceFutureQouteUSTDTable()
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

func (d *dynamoDBRepository) UpdateQouteUSDT(ctx context.Context, qouteUSDT *models.BinanceFutureQouteUSDT) error {
	table := models.NewBinanceFutureQouteUSTDTable()
	table.BinanceFutureQouteUSDT = qouteUSDT
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyBySymbol(),
		UpdateExpression: aws.String(fmt.Sprintf(
			"set %v = :counting_long, %v = :counting_short, %v = :current_leverage",
			table.GetCountingLongTableField(),
			table.GetCountingShortTableField(),
			table.GetCurrentLeverageTableField(),
		)),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":counting_long":    &types.AttributeValueMemberS{Value: table.GetNextCountingLong().String()},
			":counting_short":   &types.AttributeValueMemberS{Value: table.GetNextCountingShort().String()},
			":current_leverage": &types.AttributeValueMemberN{Value: strconv.Itoa(table.GetCurrentLeverage())},
		},
	}
	_, err := d.dynamodb.UpdateItem(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (d *dynamoDBRepository) InsertNewSymbolUSDT(ctx context.Context, symbol string, leverage int) error {
	table := models.NewBinanceFutureQouteUSTDTable()
	table.SetCountingLong(1)
	table.SetCountingShort(1)
	table.SetCurrentLeverage(leverage)
	table.SetSymbol(symbol)
	item, err := attributevalue.MarshalMap(table.BinanceFutureQouteUSDT)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String(table.GetTableName()),
		Item:      item,
	}
	_, err = d.dynamodb.PutItem(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
