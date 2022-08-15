package auth

import (
	"log"

	"gorm.io/gorm"

	rAuth "github.com/fajarabdillahfn/merchants/internal/repository/auth"
)

type repository struct {
	conn *gorm.DB
}

func NewAuthRepo(db *gorm.DB) rAuth.Repository {
	if db == nil {
		log.Panic("missing database connection")
	}

	repo := &repository{
		conn: db,
	}

	return repo
}
