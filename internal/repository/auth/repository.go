package auth

import (
	"context"

	"github.com/fajarabdillahfn/merchants/internal/models"
)

type Repository interface {
	//user
	GetUserByUserName(ctx context.Context, username string) (user models.User, err error)
}
