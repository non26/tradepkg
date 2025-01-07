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
	// table bn_ft_opening_position
	GetAllOpenOrders(ctx context.Context) (map[string]*models.BnFtOpeningPosition, error)
	GetOpenOrderBySymbolAndPositionSide(ctx context.Context, data *models.BnFtOpeningPosition) (*models.BnFtOpeningPosition, error)
	GetOpenOrderByClientID(ctx context.Context, clientId string) (*models.BnFtOpeningPosition, error)
	GetOpenOrderByKey(ctx context.Context, key map[string]types.AttributeValue) (*models.BnFtOpeningPosition, error)
	InsertNewOpenOrder(ctx context.Context, openOrder *models.BnFtOpeningPosition) error
	UpdateOpenOrder(ctx context.Context, openOrder *models.BnFtOpeningPosition) error
	DeleteOpenOrderBySymbolAndPositionSide(ctx context.Context, openOrder *models.BnFtOpeningPosition) error

	// table bn_future_qoute_usdt
	GetQouteUSDT(ctx context.Context, symbol string) (*models.BnFtQouteUSDT, error)
	UpdateQouteUSDT(ctx context.Context, qouteUSDT *models.BnFtQouteUSDT) error
	InsertNewSymbolQouteUSDT(ctx context.Context, data *models.BnFtQouteUSDT) error

	// table bn_future_history
	GetAllHistory(ctx context.Context) ([]models.BnFtHistory, error)
	GetHistoryByClientID(ctx context.Context, clientId string) (*models.BnFtHistory, error)
	InsertHistory(ctx context.Context, history *models.BnFtHistory) error
	UpdateHistory(ctx context.Context, history *models.BnFtHistory) error

	// table bot_on_run
	GetBotOnRunByBotIDAndOrderID(ctx context.Context, botOnRun *models.BnFtBotOnRun) (*models.BnFtBotOnRun, error)
	InsertBotOnRun(ctx context.Context, botOnRun *models.BnFtBotOnRun) error
	UpdateBotOnRun(ctx context.Context, botOnRun *models.BnFtBotOnRun) error
	DeleteBotOnRun(ctx context.Context, botOnRun *models.BnFtBotOnRun) error

	// table bot
	GetBotByBotID(ctx context.Context, botID string) (*models.BnFtBot, error)

	// table bn_future_advanced_position
	GetAdvancedPositionBySymbolAndPositionSide(ctx context.Context, advancedPosition *models.BnFtAdvancedPositionModel) (*models.BnFtAdvancedPositionModel, error)
	InsertAdvancedPosition(ctx context.Context, advancedPosition *models.BnFtAdvancedPositionModel) error
	// UpdateAdvancedPosition(ctx context.Context, advancedPosition *models.BnFtAdvancedPositionTable) error
	DeleteAdvancedPosition(ctx context.Context, advancedPosition *models.BnFtAdvancedPositionModel) error
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
