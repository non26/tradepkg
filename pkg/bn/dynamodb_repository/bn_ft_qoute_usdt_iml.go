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

func (d *dynamoDBRepository) GetQouteUSDT(ctx context.Context, symbol string) (*models.BnFtQouteUSDT, error) {
	var err error
	var response *dynamodb.GetItemOutput
	result := &models.BnFtQouteUSDT{}
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

func (d *dynamoDBRepository) UpdateQouteUSDT(ctx context.Context, qouteUSDT *models.BnFtQouteUSDT) error {
	table := models.NewBinanceFutureQouteUSTDTable()
	table.BnFtQouteUSDT = qouteUSDT
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyBySymbol(),
		UpdateExpression: aws.String(fmt.Sprintf(
			"set %v = :counting_long, %v = :counting_short",
			table.GetCountingLongTableField(),
			table.GetCountingShortTableField(),
		)),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":counting_long":  &types.AttributeValueMemberN{Value: strconv.Itoa(table.GetCountingLong())},
			":counting_short": &types.AttributeValueMemberN{Value: strconv.Itoa(table.GetCountingShort())},
			// ":current_leverage": &types.AttributeValueMemberN{Value: strconv.Itoa(table.GetCurrentLeverage())},
		},
	}
	_, err := d.dynamodb.UpdateItem(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (d *dynamoDBRepository) InsertNewSymbolQouteUSDT(ctx context.Context, data *models.BnFtQouteUSDT) error {
	table := models.NewBinanceFutureQouteUSTDTable()
	table.BnFtQouteUSDT = data
	item, err := attributevalue.MarshalMap(table.BnFtQouteUSDT)
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
