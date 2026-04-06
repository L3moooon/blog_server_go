package admin

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

var UserRouterApp = new(UserRouter)

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	user := Router.Group("user")
	{
		user.GET("/getUserList", userApi.GetUserList)
	}
}
