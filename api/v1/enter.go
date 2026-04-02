package v1

import (
	"blog_backend_go/api/v1/admin"
	"blog_backend_go/api/v1/web"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	AdminApiGroup admin.ApiGroup
	WebApiGroup web.ApiGroup
}
