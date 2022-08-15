package http

import uRevenue "github.com/fajarabdillahfn/merchants/internal/usecase/revenue"

type Delivery struct {
	revenueUC uRevenue.UseCase
}

func NewTestApplicationHTTP(revenueUC uRevenue.UseCase) *Delivery {
	return &Delivery{
		revenueUC: revenueUC,
	}
}
