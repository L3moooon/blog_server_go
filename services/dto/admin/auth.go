package admin

import "blog_backend_go/api/v1/dto/admin/auth"

type LoginParam struct {
	auth.LoginRequest
	Ip        string `json:"ip"`
	UserAgent string `json:"user_agent"`
}
