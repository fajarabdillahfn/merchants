package revenue

import (
	"gorm.io/gorm"

	httpRevenue "github.com/fajarabdillahfn/merchants/internal/delivery/revenue/http"
	rRevenue "github.com/fajarabdillahfn/merchants/internal/repository/revenue"
	msRevenueRepository "github.com/fajarabdillahfn/merchants/internal/repository/revenue/mysql"
	uRevenue "github.com/fajarabdillahfn/merchants/internal/usecase/revenue"
	uRevenueV1 "github.com/fajarabdillahfn/merchants/internal/usecase/revenue/v1"
)

var (
	msRevenueRepo  rRevenue.Repository
	revenueUseCase uRevenue.UseCase
	HTTPDelivery   *httpRevenue.Delivery
)

func Initialize(msMerchant *gorm.DB) {
	msRevenueRepo = msRevenueRepository.NewRevenueRepo(msMerchant)
	revenueUseCase = uRevenueV1.NewRevenueUseCase(msRevenueRepo)
	HTTPDelivery = httpRevenue.NewTestApplicationHTTP(revenueUseCase)
}
