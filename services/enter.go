package service

import (
	"blog_backend_go/services/handler/admin"
	"blog_backend_go/services/handler/web"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	AdminServiceGroup admin.ServiceGroup
	WebServiceGroup   web.ServiceGroup
}
