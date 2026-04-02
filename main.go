package main

import (
	"blog_backend_go/core"
	"blog_backend_go/global"
)

func main() {
	global.VP = core.InitViper()
	global.DB = core.InitDataBase()
	core.AutoMigrate()

	core.RunServer()
}
