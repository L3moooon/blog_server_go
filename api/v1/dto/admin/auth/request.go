package auth

type LoginType string

const (
	LoginTypeAccount LoginType = "account" // 账号登录
	LoginTypeEmail   LoginType = "email"   // 邮箱登录
	LoginTypePhone   LoginType = "phone"   // 手机号登录
)

// 注册请求
type RegisterRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Username string `json:"username"`
}

// 登录请求
type LoginRequest struct {
	LoginType LoginType `json:"login_type"`
	Account   string    `json:"account"`
	Password  string    `json:"password"`
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

// 重置密码请求
type ResetPasswordRequest struct {
	Email    string `json:"email"`
	Code     string `json:"code"`
	Password string `json:"password"`
}
