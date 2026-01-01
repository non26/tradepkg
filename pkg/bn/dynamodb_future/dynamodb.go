package dynamodbfuture

import (
	"context"

	models "github.com/non26/tradepkg/pkg/bn/dynamodb_future/models"
)

type IBnFtOpeningPositionRepository interface {
	GetAll(ctx context.Context) (map[string]*models.BnFtOpeningPosition, error)
	Get(ctx context.Context, symbol string, positionSide string) (*models.BnFtOpeningPosition, error)
	ScanWith(ctx context.Context, clientId string) (*models.BnFtOpeningPosition, error)
	Insert(ctx context.Context, openOrder *models.BnFtOpeningPosition) error
	UpdateAmountB(ctx context.Context, openOrder *models.BnFtOpeningPosition) error
	Delete(ctx context.Context, openOrder *models.BnFtOpeningPosition) error
	Upsert(ctx context.Context, openOrder *models.BnFtOpeningPosition) error
}

type IBnFtCryptoRepository interface {
	Get(ctx context.Context, symbol string) (*models.BnFtCrypto, error)
	Update(ctx context.Context, qouteUSDT *models.BnFtCrypto) error
	Insert(ctx context.Context, data *models.BnFtCrypto) error
	Upsert(ctx context.Context, data *models.BnFtCrypto) error
}

type IBnFtHistoryRepository interface {
	Get(ctx context.Context, clientId string) (*models.BnFtHistory, error)
	Insert(ctx context.Context, history *models.BnFtHistory) error
}

type IBnFtAdvancedPositionRepository interface {
	Get(ctx context.Context, clientId string) (*models.BnFtAdvancedPosition, error)
	Upsert(ctx context.Context, advancedPosition *models.BnFtAdvancedPosition) error
	Delete(ctx context.Context, clientId string) error
}

type IBnFtBotOnRunRepository interface {
	Get(ctx context.Context, botOnRun *models.BnFtBotOnRun) (*models.BnFtBotOnRun, error)
	Delete(ctx context.Context, botOnRun *models.BnFtBotOnRun) error
	Upsert(ctx context.Context, botOnRun *models.BnFtBotOnRun) error
	GetAll(ctx context.Context) ([]models.BnFtBotOnRun, error)
	ScanWith(ctx context.Context, clientId string) ([]models.BnFtBotOnRun, error)
}

type IBnFtBotRepository interface {
	Get(ctx context.Context, botID string) (*models.BnFtBot, error)
}

type IBnFtBotRegistorRepository interface {
	Get(ctx context.Context, botID string, botOrderID string) (*models.BnFtBotRegistor, error)
	Upsert(ctx context.Context, botRegistor *models.BnFtBotRegistor) error
	GetAll(ctx context.Context) ([]models.BnFtBotRegistor, error)
	Delete(ctx context.Context, botID string, botOrderID string) error
}
