package admin

import (
	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

var RoleRouterApp = new(RoleRouter)

func (s *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	role := Router.Group("role")
	{
		role.GET("/", roleApi.GetRole)
		role.POST("/", roleApi.AddRole)
		role.PUT("/", roleApi.UpdateRole)
		role.DELETE("/", roleApi.DeleteRole)
	}

}
