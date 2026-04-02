package core

import (
	"blog_backend_go/global"
	"blog_backend_go/initialize"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// InitServer 初始化总路由
func RunServer() {
	// 从配置文件读取 gin 模式
	mode := global.VP.GetString("system.mode")
	if mode == "" {
		mode = "debug" // 默认调试模式
	}
	gin.SetMode(mode)

	// 注册路由
	Router := initialize.InitRouter()

	// 启动服务器
	address := fmt.Sprintf(":%d", global.CONFIG.System.Port)
	initServer(address, Router, 10*time.Minute, 10*time.Minute)
}
