package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"

	"blog_backend_go/config"
	"blog_backend_go/router"
	"blog_backend_go/utils"
)

func main() {
	//读取命令行环境参数
	var env string
	flag.StringVar(&env, "env", "dev", "运行环境: dev/test/prod")
	flag.Parse() // 必须调用 Parse() 才能拿到值
	log.Printf("当前运行环境: %s", env)

	//加载环境变量文件
	err := godotenv.Load(".env." + env)
	if err != nil {
		log.Println("No .env file found")
	}

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080" // 默认端口
	}

	//初始化数据库
	config.InitDB()

	//执行GORM模型自动迁移
	utils.AutoMigrate(config.DB)

	//初始化路由
	r := router.InitRouter()

	//启动服务
	log.Printf("🚀 服务已启动：http://localhost:%s", port)
	r.Run(":" + port)
}
