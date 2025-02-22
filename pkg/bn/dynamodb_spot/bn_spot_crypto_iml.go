package dynamodbspot

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	models "github.com/non26/tradepkg/pkg/bn/dynamodb_spot/models"
)

type bnSpotCryptoRepository struct {
	client *dynamodb.Client
}

func NewConnectionBnSpotCryptoRepository(client *dynamodb.Client) IBnSpotCryptoRepository {
	return &bnSpotCryptoRepository{
		client: client,
	}
}

func (d *bnSpotCryptoRepository) Get(ctx context.Context, symbol string) (*models.BnSpotCrypto, error) {
	var err error
	var response *dynamodb.GetItemOutput
	result := &models.BnSpotCrypto{}
	table := models.NewBinanceSpotCryptoTable()
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

func (d *bnSpotCryptoRepository) Update(ctx context.Context, qouteUSDT *models.BnSpotCrypto) error {
	table := models.NewBinanceSpotCryptoTableWith(qouteUSDT)
	table.Transform()
	countingLong, _ := table.GetCountingTableField()
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyBySymbol(),
		UpdateExpression: aws.String(fmt.Sprintf(
			"set %v = :counting",
			countingLong,
		)),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":counting": &types.AttributeValueMemberN{Value: strconv.FormatInt(table.GetCounting(), 10)},
		},
	}
	_, err := d.client.UpdateItem(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (d *bnSpotCryptoRepository) Insert(ctx context.Context, data *models.BnSpotCrypto) error {
	table := models.NewBinanceSpotCryptoTableWith(data)
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
