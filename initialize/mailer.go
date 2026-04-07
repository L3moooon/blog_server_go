package initialize

import (
	"blog_backend_go/global"
	"log"

	"gopkg.in/gomail.v2"
)

func InitEmail() *gomail.Dialer {
	config := global.CONFIG.Email
	instance := gomail.NewDialer(
		config.Host,
		config.Port,
		config.AuthUser,
		config.AuthPass,
	)
	log.Println("邮箱服务初始化成功")
	return instance
}
