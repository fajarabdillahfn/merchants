package models

import "time"

type Transaction struct {
	MerchantId int64
	OutletId   int64
	BillTotal  float64
	CreatedAt  time.Time
	CreatedBy  int64
	UpdatedAt  time.Time
	UpdatedBy  int64
}
