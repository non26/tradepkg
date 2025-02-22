package dynamodbfuture

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	dynamodbconfig "github.com/non26/tradepkg/pkg/bn/dynamodb_config"
	models "github.com/non26/tradepkg/pkg/bn/dynamodb_future/models"
)

type bnFtCryptoRepository struct {
	client *dynamodb.Client
}

func NewConnectionBnFtCryptoRepository(client *dynamodb.Client) IBnFtCryptoRepository {
	return &bnFtCryptoRepository{
		client: client,
	}
}

func (d *bnFtCryptoRepository) Get(ctx context.Context, symbol string) (*models.BnFtCrypto, error) {
	var err error
	var response *dynamodb.GetItemOutput
	result := &models.BnFtCrypto{}
	table := models.NewBinanceFutureCryptoTable()
	table.Symbol = symbol
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

func (d *bnFtCryptoRepository) Update(ctx context.Context, qouteUSDT *models.BnFtCrypto) error {
	table := models.NewBinanceFutureCryptoTableWith(qouteUSDT)
	table.Transform()
	countingLong, _ := table.GetCountingLongTableField()
	countingShort, _ := table.GetCountingShortTableField()
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyBySymbol(),
		UpdateExpression: aws.String(fmt.Sprintf(
			"set %v = :counting_long, %v = :counting_short",
			countingLong,
			countingShort,
		)),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":counting_long":  &types.AttributeValueMemberN{Value: strconv.Itoa(int(table.GetCountingLong()))},
			":counting_short": &types.AttributeValueMemberN{Value: strconv.Itoa(int(table.GetCountingShort()))},
		},
	}
	_, err := d.client.UpdateItem(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (d *bnFtCryptoRepository) Insert(ctx context.Context, data *models.BnFtCrypto) error {
	table := models.NewBinanceFutureCryptoTableWith(data)
	table.Transform()
	item, err := attributevalue.MarshalMap(table.GetData())
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String(table.GetTableName()),
		Item:      item,
	}
	_, err = d.client.PutItem(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (d *bnFtCryptoRepository) Upsert(ctx context.Context, data *models.BnFtCrypto) error {
	table := models.NewBinanceFutureCryptoTableWith(data)
	table.Transform()
	update_config := dynamodbconfig.NewUpdateTable(data)
	update_config.Set(table.GetCountingLongTableField, data.CountingLong)
	update_config.Set(table.GetCountingShortTableField, data.CountingShort)
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
