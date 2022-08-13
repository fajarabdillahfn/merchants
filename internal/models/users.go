package models

import (
	"time"
)

type User struct {
	Name      string
	UserName  string
	Password  string
	CreatedAt time.Time
	CreatedBy int64
	UpdatedAt time.Time
	UpdatedBy int64
}
