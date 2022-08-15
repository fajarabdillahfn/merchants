package v1

import (
	rRevenue "github.com/fajarabdillahfn/merchants/internal/repository/revenue"
	uRevenue "github.com/fajarabdillahfn/merchants/internal/usecase/revenue"
)

type useCase struct {
	RevenueRepo rRevenue.Repository
}

func NewRevenueUseCase(revenueRepository rRevenue.Repository) uRevenue.UseCase {
	return &useCase{
		RevenueRepo: revenueRepository,
	}
}
