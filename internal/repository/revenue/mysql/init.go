package revenue

import (
	"log"

	"gorm.io/gorm"

	rRevenue "github.com/fajarabdillahfn/merchants/internal/repository/revenue"
)

type repository struct {
	conn *gorm.DB
}

func NewRevenueRepo(db *gorm.DB) rRevenue.Repository {
	if db == nil {
		log.Panic("missing database connection")
	}

	repo := &repository{
		conn: db,
	}

	return repo
}
