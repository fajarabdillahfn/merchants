package v1

import (
	rAuth "github.com/fajarabdillahfn/merchants/internal/repository/auth"
	uAuth "github.com/fajarabdillahfn/merchants/internal/usecase/auth"
)

type useCase struct {
	AuthRepo rAuth.Repository
}

func NewAuthUseCase(authRepository rAuth.Repository) uAuth.UseCase {
	return &useCase{
		AuthRepo: authRepository,
	}
}
