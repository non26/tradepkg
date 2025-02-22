package dynamodbspot

import (
	"context"

	models "github.com/non26/tradepkg/pkg/bn/dynamodb_spot/models"
)

type IBnSpotOpeningPositionRepository interface {
	GetAll(ctx context.Context) (map[string]*models.BnSpotOpeningPosition, error)
	Get(ctx context.Context, data *models.BnSpotOpeningPosition) (*models.BnSpotOpeningPosition, error)
	ScanWith(ctx context.Context, clientId string) (*models.BnSpotOpeningPosition, error)
	Insert(ctx context.Context, openOrder *models.BnSpotOpeningPosition) error
	UpdateAmountB(ctx context.Context, openOrder *models.BnSpotOpeningPosition) error
	Delete(ctx context.Context, openOrder *models.BnSpotOpeningPosition) error
	Upsert(ctx context.Context, openOrder *models.BnSpotOpeningPosition) error
}

type IBnSpotCryptoRepository interface {
	Get(ctx context.Context, symbol string) (*models.BnSpotCrypto, error)
	Update(ctx context.Context, qouteUSDT *models.BnSpotCrypto) error
	Insert(ctx context.Context, data *models.BnSpotCrypto) error
}

type IBnSpotHistoryRepository interface {
	Get(ctx context.Context, clientId string) (*models.BnSpotHistory, error)
	Insert(ctx context.Context, history *models.BnSpotHistory) error
}
