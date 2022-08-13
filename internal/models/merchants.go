package models

import "time"

type Merchant struct {
	UserId       int
	MerchantName string
	CreatedAt    time.Time
	CreatedBy    int64
	UpdatedAt    time.	Time
	UpdatedBy    int64
}
