package models

import (
	"time"
)

type User struct {
	Id        int       `gorm:"column:id;primaryKey;"`
	Name      string    `gorm:"column:name"`
	UserName  string    `gorm:"column:user_name"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at"`
	CreatedBy int64     `gorm:"column:created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	UpdatedBy int64     `gorm:"column:updated_by"`
}

func (User) TableName() string {
	return "users"
}
