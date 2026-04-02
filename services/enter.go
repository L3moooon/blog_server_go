package service

import (
	"blog_backend_go/services/admin"
	"blog_backend_go/services/web"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	AdminServiceGroup admin.ServiceGroup
	WebServiceGroup   web.ServiceGroup
}
