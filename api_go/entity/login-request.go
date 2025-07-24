package entity

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}