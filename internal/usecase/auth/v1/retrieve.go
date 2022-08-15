package v1

import (
	"context"
	"time"

	"github.com/fajarabdillahfn/merchants/internal/models"
)

func (u *useCase) UserGetByUsername(ctx context.Context, username string) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*5000)
	defer cancel()

	return u.AuthRepo.GetUserByUserName(ctx, username)
}
