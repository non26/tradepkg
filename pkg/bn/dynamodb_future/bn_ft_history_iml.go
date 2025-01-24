package dynamodbfuture

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	models "github.com/non26/tradepkg/pkg/bn/dynamodb_future/models"
)

type bnFtHistoryRepository struct {
	client *dynamodb.Client
}

func NewConnectionBnFtHistoryRepository(client *dynamodb.Client) IBnFtHistoryRepository {
	return &bnFtHistoryRepository{
		client: client,
	}
}

func (d *bnFtHistoryRepository) Get(ctx context.Context, clientId string) (*models.BnFtHistory, error) {
	table := models.NewBinanceFutureHistoryTable()
	table.ClientId = clientId
	result := &models.BnFtHistory{}
	response, err := d.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyClientID(),
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

func (d *bnFtHistoryRepository) Insert(ctx context.Context, history *models.BnFtHistory) error {
	table := models.NewBinanceFutureHistoryTableWith(history)
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

// func (d *bnFtHistoryRepository) Update(ctx context.Context, history *models.BnFtHistory) error {
// 	table := models.NewBinanceFutureHistoryTableWith(history)
// 	table.Transform()
// 	item, err := attributevalue.MarshalMap(table.GetData())
// 	if err != nil {
// 		return err
// 	}
// 	_, err = d.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
// 		TableName: aws.String(table.GetTableName()),
// 		Key:       table.GetKeyClientID(),
// 		UpdateExpression: aws.String(fmt.Sprintf(
// 			"set %v = :history_id, %v = :symbol, %v = :position_side",
// 			table.GetHistoryIdTableField(),
// 			table.GetSymbolTableField(),
// 			table.GetPositionSideTableField(),
// 		)),
// 		ExpressionAttributeValues: item,
// 	})
// 	return err
// }
