package auth

import (
	"context"
	"log"

	models "github.com/fajarabdillahfn/merchants/internal/models"
)

func (r *repository) GetUserByUserName(ctx context.Context, username string) (user models.User, err error) {
	tableName := models.User.TableName(models.User{})

	dataPrep := r.conn.WithContext(ctx).Table(tableName)

	res := dataPrep.First(&user, "user_name = ?", username)
	if res.Error != nil {
		log.Println("ERROR get user: " + res.Error.Error())
		return models.User{}, res.Error
	}

	return
}
