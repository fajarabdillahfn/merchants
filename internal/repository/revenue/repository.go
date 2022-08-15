package revenue

import (
	"context"

	"github.com/fajarabdillahfn/merchants/internal/models"
)

type Repository interface {
	GetMerchantsByUserId(ctx context.Context, userId int) (merchants []models.Merchant, err error)

	GetMerchantRevenuePerDay(ctx context.Context, startDate string, endDate string, merchantId int) (revenue float64, err error)
}
