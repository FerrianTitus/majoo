package models

type Users struct {
	ID        int    `json:"id" form:"id" query:"id"`
	Name      string `json:"name" form:"name" query:"name"`
	Username  string `json:"user_name" form:"user_name" query:"user_name"`
	Password  string `json:"password" form:"password" query:"password"`
	CreatedAt string `json:"created_at" form:"created_at" query:"created_at"`
	CreatedBy string `json:"created_by" form:"created_by" query:"created_by"`
	UpdatedAt string `json:"updated_at" form:"updated_at" query:"updated_at"`
	UpdatedBy string `json:"updated_by" form:"updated_by" query:"updated_by"`
}
