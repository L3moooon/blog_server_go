package initialize

import (
	"blog_backend_go/global"

	"gopkg.in/gomail.v2"
)

func InitEmail() *gomail.Dialer {
	cfg := global.CONFIG.Email
	instance := gomail.NewDialer(
		cfg.Host,
		cfg.Port,
		cfg.AuthUser,
		cfg.AuthPass,
	)
	return instance
}
