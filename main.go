package main

import (
	"blog_backend_go/core"
	"blog_backend_go/global"
	"blog_backend_go/initialize"
)

func main() {
	global.VP = core.InitViper()    // 初始化viper
	global.DB = core.InitDataBase() // 初始化数据库
	global.LOG = core.Zap()         // 初始化日志
	global.REDIS = core.InitRedis() // 初始化redis

	global.IP2REGION = initialize.InitIP2Region() // 初始化ip2region
	global.Mailer = initialize.InitEmail()        // 初始化邮箱
	global.SMS = initialize.InitSMS()             // 初始化短信

	core.AutoMigrate() // 数据库自动迁移
	core.RunServer()   // 启动服务器

}
