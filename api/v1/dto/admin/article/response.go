package auth

// 登录响应
type LoginResponse struct {
	Token string `json:"token"`
}

// 注册响应
type RegisterResponse struct {
	Token string `json:"token"`
}

// 获取邮箱验证码响应
type EmailCaptchaResponse struct {
	Code string `json:"code"`
}

// 获取手机验证码响应
type SmsCaptchaResponse struct {
	Code string `json:"code"`
}
