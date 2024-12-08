package dynamodbrepository

import (
	"context"

	models "github.com/non26/tradepkg/pkg/bn/dynamodb_repository/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type dynamodbConfig struct{}

func (d *dynamodbConfig) LoadConfig() aws.Config {
	cfg, _ := config.LoadDefaultConfig(context.TODO())
	return cfg
}
func NewDynamodbConfig() *dynamodbConfig {
	return &dynamodbConfig{}
}

type credential struct {
	AccessKeyID     string
	SecretAccessKey string
}

func (c *credential) Retrieve(ctx context.Context) (aws.Credentials, error) {
	return aws.Credentials{
		AccessKeyID:     c.AccessKeyID,
		SecretAccessKey: c.SecretAccessKey,
	}, nil
}

func NewCredential(accessKeyID, secretAccessKey string) *credential {
	return &credential{
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
	}
}

type IRepository interface {
	GetAllOpenOrders(ctx context.Context) ([]models.BinanceFutureOpeningPosition, error)
	GetOpenOrderBySymbol(ctx context.Context, symbol string) ([]models.BinanceFutureOpeningPosition, error)
	GetOpenOrderByClientID(ctx context.Context, clientId string) ([]models.BinanceFutureOpeningPosition, error)
	GetOpenOrderByKey(ctx context.Context, key map[string]types.AttributeValue) (*models.BinanceFutureOpeningPosition, error)
	NewOpenOrder(ctx context.Context, openOrder *models.BinanceFutureOpeningPosition) error
	UpdateOpenOrder(ctx context.Context, openOrder *models.BinanceFutureOpeningPosition) error
	DeleteOpenOrderBySymbol(ctx context.Context, symbol string) error
	DeleteOpenOrderByKey(ctx context.Context, key map[string]types.AttributeValue) error
	// table bn_future_qoute_usdt
	GetQouteUSDT(ctx context.Context, symbol string) (*models.BinanceFutureQouteUSDT, error)
	UpdateCountingSymbolQouteUSDT(ctx context.Context, qouteUSDT *models.BinanceFutureQouteUSDT) error
	InsertNewSymbolUSDT(ctx context.Context, symbol string) error
}

type dynamoDBRepository struct {
	dynamodb *dynamodb.Client
}

func NewDynamoDBRepository(
	dynamodb *dynamodb.Client,
) IRepository {
	return &dynamoDBRepository{
		dynamodb: dynamodb,
	}
}

type newDynamodb struct {
	endPoint   dynamodb.EndpointResolverV2
	credential aws.CredentialsProvider
	awsconfig  aws.Config
}

func (c newDynamodb) NewLocal() *dynamodb.Client {
	svc := dynamodb.NewFromConfig(c.awsconfig, func(o *dynamodb.Options) {
		o.Credentials = c.credential
		o.EndpointResolverV2 = c.endPoint
	})
	return svc
}

func (c newDynamodb) NewPrd() *dynamodb.Client {
	svc := dynamodb.NewFromConfig(c.awsconfig, func(o *dynamodb.Options) {
		o.Credentials = c.credential
	})
	return svc
}

func DynamoDB(
	endPoint dynamodb.EndpointResolverV2,
	credential aws.CredentialsProvider,
	awsconfig aws.Config,
) *newDynamodb {
	return &newDynamodb{
		endPoint,
		credential,
		awsconfig,
	}
}
