package dynamodbfuture

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	models "github.com/non26/tradepkg/pkg/bn/dynamodb_future/models"
)

type bnFtBotRepository struct {
	client *dynamodb.Client
}

func NewConnectionBnFtBotRepository(client *dynamodb.Client) IBnFtBotRepository {
	return &bnFtBotRepository{
		client: client,
	}
}

func (d *bnFtBotRepository) Get(ctx context.Context, botID string) (*models.BnFtBot, error) {
	table := models.NewBnFtBotTable()
	table.BotID = botID
	result := &models.BnFtBot{}
	response, err := d.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyByBotID(),
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
