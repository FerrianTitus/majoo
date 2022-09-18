package models

type LoginRequest struct {
	Username string `json:"user_name" form:"user_name" query:"user_name"`
	Password string `json:"password" form:"password" query:"password"`
}

type LoginResponse struct {
	Token string `json:"token" form:"token" query:"token"`
}

type DataFromToken struct {
	ID       int    `json:"id" form:"id" query:"id"`
	Name     string `json:"name" form:"name" query:"name"`
	Username string `json:"user_name" form:"user_name" query:"user_name"`
}
