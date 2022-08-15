package auth

import (
	"gorm.io/gorm"

	
	httpAuth "github.com/fajarabdillahfn/merchants/internal/delivery/auth/http"
	rAuth "github.com/fajarabdillahfn/merchants/internal/repository/auth"
	msAuthRepository "github.com/fajarabdillahfn/merchants/internal/repository/auth/mysql"
	uAuth "github.com/fajarabdillahfn/merchants/internal/usecase/auth"
	uAuthV1 "github.com/fajarabdillahfn/merchants/internal/usecase/auth/v1"
)

var (
	msAuthRepo   rAuth.Repository
	authUseCase  uAuth.UseCase
	HTTPDelivery *httpAuth.Delivery
)

func Initialize(msMerchant *gorm.DB) {
	msAuthRepo = msAuthRepository.NewAuthRepo(msMerchant)
	authUseCase = uAuthV1.NewAuthUseCase(msAuthRepo)
	HTTPDelivery = httpAuth.NewTestApplicationHTTP(authUseCase)
}
