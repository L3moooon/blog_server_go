package utils

import (
	"blog_backend_go/global"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "时雨博客"+"<"+global.CONFIG.Email.AuthUser+">") // 设置发件人
	m.SetHeader("To", to)                                            // 设置收件人
	m.SetHeader("Subject", subject)                                  // 标题
	m.SetBody("text/html; charset=utf-8", body)                      // 内容（HTML格式）
	return global.Mailer.DialAndSend(m)
}
