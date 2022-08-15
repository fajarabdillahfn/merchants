package models

type Revenue struct {
	Date    string            `json:"date"`
	Revenue []MerchantDetails `json:"revenue,omitempty"`
}

type MerchantDetails struct {
	MerchantName string  `json:"merchant_name,omitempty"`
	OutletName   string  `json:"outlet_name,omitempty"`
	DailyRevenue float64 `json:"daily_revenue"`
}
