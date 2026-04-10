package auth

import "blog_backend_go/model"

type Permission struct {
	Routes     []string `json:"routes"`
	Btns       []string `json:"btns"`
	Components []string `json:"components"`
}

// 登录响应
type LoginResponse struct {
	User        model.SysUser `json:"user"`
	Token       string        `json:"token"`
	ExpiresTime int64         `json:"expires_time"`
	Permissions Permission    `json:"permissions"`
}

// 注册响应
type RegisterResponse struct {
}

// 获取邮箱验证码响应
type EmailCaptchaResponse struct {
	Code string `json:"code"`
}

// 获取手机验证码响应
type SmsCaptchaResponse struct {
	Code string `json:"code"`
}
