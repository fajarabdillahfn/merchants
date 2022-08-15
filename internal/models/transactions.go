package models

import "time"

type Transaction struct {
	Id         int       `gorm:"column:id;primaryKey;"`
	MerchantId int64     `gorm:"column:merchant_id"`
	OutletId   int64     `gorm:"column:outlet_id"`
	BillTotal  float64   `gorm:"column:bill_total"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	CreatedBy  int64     `gorm:"column:created_by"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
	UpdatedBy  int64     `gorm:"column:updated_by"`
}

func (Transaction) TableName() string {
	return "transactions"
}