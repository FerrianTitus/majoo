package models

type Outlets struct {
	ID         int    `json:"id" form:"id" query:"id"`
	MerchantId int    `json:"merchant_id" form:"merchant_id" query:"merchant_id"`
	OutletName string `json:"outlet_name" form:"outlet_name" query:"outlet_name"`
	CreatedAt  string `json:"created_at" form:"created_at" query:"created_at"`
	CreatedBy  string `json:"created_by" form:"created_by" query:"created_by"`
	UpdatedAt  string `json:"updated_at" form:"updated_at" query:"updated_at"`
	UpdatedBy  string `json:"updated_by" form:"updated_by" query:"updated_by"`
}

type GetOutletByUserId struct {
	OutletName   string `json:"outlet_name" form:"outlet_name" query:"outlet_name"`
	MerchantName string `json:"merchant_name" form:"merchant_name" query:"merchant_name"`
	UserName     string `json:"user_name" form:"user_name" query:"user_name"`
}
