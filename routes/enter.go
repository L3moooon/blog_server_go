package router

import (
	"blog_backend_go/routes/admin"
	"blog_backend_go/routes/web"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	Admin admin.RouterGroup
	Web   web.RouterGroup
}
