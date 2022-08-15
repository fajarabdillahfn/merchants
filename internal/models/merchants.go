package models

import "time"

type Merchant struct {
	Id           int       `gorm:"column:id;primaryKey;"`
	UserId       int       `gorm:"column:user_id"`
	MerchantName string    `gorm:"column:merchant_name"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	CreatedBy    int64     `gorm:"column:created_by"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
	UpdatedBy    int64     `gorm:"column:updated_by"`
}
