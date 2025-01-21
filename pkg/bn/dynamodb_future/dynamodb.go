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
	Update(ctx context.Context, openOrder *models.BnFtOpeningPosition) error
	Delete(ctx context.Context, openOrder *models.BnFtOpeningPosition) error
}

type IBnFtQouteUSDTRepository interface {
	Get(ctx context.Context, symbol string) (*models.BnFtQouteUSDT, error)
	Update(ctx context.Context, qouteUSDT *models.BnFtQouteUSDT) error
	Insert(ctx context.Context, data *models.BnFtQouteUSDT) error
}

type IBnFtHistoryRepository interface {
	Get(ctx context.Context, clientId string) (*models.BnFtHistory, error)
	Insert(ctx context.Context, history *models.BnFtHistory) error
}

type IBnFtBotOnRunRepository interface {
	Get(ctx context.Context, botOnRun *models.BnFtBotOnRun) (*models.BnFtBotOnRun, error)
	Insert(ctx context.Context, botOnRun *models.BnFtBotOnRun) error
	Update(ctx context.Context, botOnRun *models.BnFtBotOnRun) error
	Delete(ctx context.Context, botOnRun *models.BnFtBotOnRun) error
}

type IBnFtBotRepository interface {
	Get(ctx context.Context, botID string) (*models.BnFtBot, error)
}

type IBnFtAdvancedPositionRepository interface {
	Get(ctx context.Context, advancedPosition *models.BnFtAdvancedPositionModel) (*models.BnFtAdvancedPositionModel, error)
	Insert(ctx context.Context, advancedPosition *models.BnFtAdvancedPositionModel) error
	Delete(ctx context.Context, advancedPosition *models.BnFtAdvancedPositionModel) error
}

type IBnFtAccumulationRepository interface {
	Get(ctx context.Context, accumulation *models.BnFtAccumulation) (*models.BnFtAccumulation, error)
	Update(ctx context.Context, accumulation *models.BnFtAccumulation) error
	Insert(ctx context.Context, accumulation *models.BnFtAccumulation) error
	Delete(ctx context.Context, accumulation *models.BnFtAccumulation) error
}
