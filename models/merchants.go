package models

type Merchants struct {
	ID           int    `json:"id" form:"id" query:"id"`
	UserId       int    `json:"user_id" form:"user_id" query:"user_id"`
	MerchantName string `json:"merchant_name" form:"merchant_name" query:"merchant_name"`
	CreatedAt    string `json:"created_at" form:"created_at" query:"created_at"`
	CreatedBy    string `json:"created_by" form:"created_by" query:"created_by"`
	UpdatedAt    string `json:"updated_at" form:"updated_at" query:"updated_at"`
	UpdatedBy    string `json:"updated_by" form:"updated_by" query:"updated_by"`
}
