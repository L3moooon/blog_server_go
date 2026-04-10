package utils

import (
	"blog_backend_go/global"
	"fmt"
	"strings"

	dypnsapi "github.com/alibabacloud-go/dypnsapi-20170525/v3/client"
	"github.com/alibabacloud-go/tea/tea"
)

func SendSms(phone string, code string) error {
	// log.Println("发送短信", phone, code)
	param := global.CONFIG.SMS.TemplateParam
	param = strings.ReplaceAll(param, "##code##", code)
	sendSmsRequest := &dypnsapi.SendSmsVerifyCodeRequest{
		PhoneNumber:   tea.String(phone),
		SignName:      tea.String(global.CONFIG.SMS.SignName),
		TemplateCode:  tea.String(global.CONFIG.SMS.TemplateCode),
		TemplateParam: tea.String(param),
	}

	resp, err := global.SMS.SendSmsVerifyCode(sendSmsRequest)
	if err != nil {
		return err
	}

	// 业务代码不是 OK 则视为失败
	if *resp.Body.Code != "OK" {
		return fmt.Errorf("号码认证发送验证码失败: %s, %s", *resp.Body.Code, *resp.Body.Message)
	}

	return nil
}
