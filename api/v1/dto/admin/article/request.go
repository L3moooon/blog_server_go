package auth


// 登录请求
type ListRequest struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Title    string `json:"title"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Code      string    `json:"code"`
}

// 获取邮箱验证码请求
type EmailCaptchaRequest struct {
	Email string `json:"email"`
}

// 获取手机验证码请求
type SmsCaptchaRequest struct {
	Phone string `json:"phone"`
}

// 注册请求
type RegisterRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Code     string `json:"code"`
}
