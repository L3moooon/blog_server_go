package admin

import (
	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

var RoleRouterApp = new(RoleRouter)

func (s *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	role := Router.Group("role")
	{
		role.GET("/getRoleList", roleApi.GetRoleList)
		role.POST("/addRole", roleApi.AddRole)
		role.PUT("/updateRole", roleApi.UpdateRole)
		role.DELETE("/deleteRole", roleApi.DeleteRole)
	}

}
