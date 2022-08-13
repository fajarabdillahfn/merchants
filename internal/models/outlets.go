package models

import "time"

type Outlet struct {
	MerchantId int64
	OutletName string
	CreatedAt  time.Time
	CreatedBy  int64
	UpdatedAt  time.Time
	UpdatedBy  int64
}
