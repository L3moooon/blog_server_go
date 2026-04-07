package initialize

import (
	"blog_backend_go/global"
	"log"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dypnsapi "github.com/alibabacloud-go/dypnsapi-20170525/v3/client"
	"github.com/alibabacloud-go/tea/tea"
)

func InitSMS() *dypnsapi.Client {
	// 1. 创建配置
	config := &openapi.Config{
		AccessKeyId:     tea.String(global.CONFIG.SMS.AccessKeyId),
		AccessKeySecret: tea.String(global.CONFIG.SMS.AccessKeySecret),
	}
	// 号码认证服务固定域名
	config.Endpoint = tea.String("dypnsapi.aliyuncs.com")

	// 2. 创建客户端
	sms, err := dypnsapi.NewClient(config)
	if err != nil {
		panic("初始化阿里云号码认证客户端失败: " + err.Error())
	}
	log.Println("短信服务（号码认证）初始化成功")
	return sms
}
