package dynamodbfuture

import (
	"context"

	models "github.com/non26/tradepkg/pkg/bn/dynamodb_future/models"
)

type IBnFtOpeningPositionRepository interface {
	GetAll(ctx context.Context) (map[string]*models.BnFtOpeningPosition, error)
	Get(ctx context.Context, data *models.BnFtOpeningPosition) (*models.BnFtOpeningPosition, error)
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

type IBnFtBotOnRunRepository interface {
	Get(ctx context.Context, botOnRun *models.BnFtBotOnRun) (*models.BnFtBotOnRun, error)
	// Insert(ctx context.Context, botOnRun *models.BnFtBotOnRun) error
	// Update(ctx context.Context, botOnRun *models.BnFtBotOnRun) error
	Delete(ctx context.Context, botOnRun *models.BnFtBotOnRun) error
	Upsert(ctx context.Context, botOnRun *models.BnFtBotOnRun) error
}

type IBnFtBotRepository interface {
	Get(ctx context.Context, botID string) (*models.BnFtBot, error)
}
