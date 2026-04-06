package admin

import (
	"github.com/gin-gonic/gin"
)

type PermissionRouter struct{}

var PermissionRouterApp = new(PermissionRouter)

func (s *PermissionRouter) InitPermissionRouter(Router *gin.RouterGroup) {
	permission := Router.Group("permission")
	{
		permission.GET("/", permissionApi.GetPermission)
		permission.POST("/", permissionApi.AddPermission)
		permission.PUT("/", permissionApi.UpdatePermission)
		permission.DELETE("/", permissionApi.DeletePermission)
	}

}
