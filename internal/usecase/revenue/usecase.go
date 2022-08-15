package revenue

import (
	"context"

	"github.com/fajarabdillahfn/merchants/internal/models"
)

type UseCase interface {
	TotalMerchantRevenue(ctx context.Context, userId int, dataType string) (revenues []models.Revenue, err error)
}
