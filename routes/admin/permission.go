package admin

import (
	"github.com/gin-gonic/gin"
)

type PermissionRouter struct{}

var PermissionRouterApp = new(PermissionRouter)

func (s *PermissionRouter) InitPermissionRouter(Router *gin.RouterGroup) {
	permission := Router.Group("/permission")
	{
		permission.GET("/getPermissionList", permissionApi.GetPermissionList)
		permission.POST("/addPermission", permissionApi.AddPermission)
		permission.PUT("/updatePermission", permissionApi.UpdatePermission)
		permission.DELETE("/deletePermission", permissionApi.DeletePermission)
	}

}
