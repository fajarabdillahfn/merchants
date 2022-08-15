package models

import "time"

type Outlet struct {
	Id         int       `gorm:"column:id;primaryKey;"`
	MerchantId int64     `gorm:"column:merchant_id"`
	OutletName string    `gorm:"column:outlet_name"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	CreatedBy  int64     `gorm:"column:created_by"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
	UpdatedBy  int64     `gorm:"column:updated_by"`
}
