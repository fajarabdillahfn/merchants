package auth

import (
	"context"

	"github.com/fajarabdillahfn/merchants/internal/models"
)

type UseCase interface {
	UserGetByUsername(ctx context.Context, username string) (models.User, error)
}
