package models

type AuthToken struct {
	Token string `json:"token"`
}

type User struct {
	UserName string `json:"username"`
	IsStaff  bool   `json:"is_staff"`
	IsActive bool   `json:"is_active"`
	Id       int    `json:"id"`
}
