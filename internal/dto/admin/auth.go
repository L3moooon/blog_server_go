package dto

type LoginRequest struct {
	LoginType int    `json:"login_type"`
	Account   string `json:"account"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Code      string `json:"code"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
