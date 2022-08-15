package http

import uAuth "github.com/fajarabdillahfn/merchants/internal/usecase/auth"

type Delivery struct {
	authUC uAuth.UseCase
}

func NewTestApplicationHTTP(authUC uAuth.UseCase) *Delivery {
	return &Delivery{
		authUC: authUC,
	}
}
