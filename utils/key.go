package utils

// 定义常量前缀，统一管理
const (
	PrefixLoginEmailCaptcha = "login:code:email:" // 登录邮箱验证码
	PrefixLoginSmsCaptcha   = "login:code:sms:"   // 登录短信验证码
	PrefixUserJwt           = "login:jwt:"        // 用户JWT
	PrefixBlackJwt          = "login:black:jwt:"  // 黑名单JWT

	PrefixResetEmailCaptcha = "reset:code:email:" // 重置邮箱验证码
	PrefixResetSmsCaptcha   = "reset:code:sms:"   // 重置短信验证码
)

// EmailCaptchaKey 邮箱验证码 key
func LoginEmailCaptchaKey(email string) string {
	return PrefixLoginEmailCaptcha + email
}

// SmsCaptchaKey 短信验证码 key
func LoginSmsCaptchaKey(phone string) string {
	return PrefixLoginSmsCaptcha + phone
}

// UserJwtKey 用户JWT key
func UserJwtKey(username string) string {
	return PrefixUserJwt + username
}

// BlackJwtKey 黑名单JWT key
func BlackJwtKey(token string) string {
	return PrefixBlackJwt + token
}

// ResetEmailCaptchaKey 重置邮箱验证码 key
func ResetEmailCaptchaKey(email string) string {
	return PrefixResetEmailCaptcha + email
}

// ResetSmsCaptchaKey 重置短信验证码 key
func ResetSmsCaptchaKey(phone string) string {
	return PrefixResetSmsCaptcha + phone
}
