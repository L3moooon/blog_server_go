package utils

import (
	"blog_backend_go/global"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", global.CONFIG.Email.AuthUser, "时雨博客") // 设置发件人
	m.SetHeader("To", to)                                            // 设置收件人
	m.SetHeader("Subject", subject)                                  // 标题
	m.SetBody("text/html; charset=utf-8", body)                      // 内容（HTML格式）
	return global.Mailer.DialAndSend(m)
}

