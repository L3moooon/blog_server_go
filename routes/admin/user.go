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
		user.GET("/getUserById", userApi.GetUserById)
		user.GET("/getUserByUsername", userApi.GetUserByUsername)
		user.GET("/getUserByEmail", userApi.GetUserByEmail)
		user.GET("/getUserByPhone", userApi.GetUserByPhone)
		user.GET("/getUserByRole", userApi.GetUserByRole)
		user.GET("/getUserByStatus", userApi.GetUserByStatus)
		user.GET("/getUserByCreateTime", userApi.GetUserByCreateTime)
		user.GET("/getUserByUpdateTime", userApi.GetUserByUpdateTime)
		user.GET("/getUserByDeleteTime", userApi.GetUserByDeleteTime)
		user.GET("/getUserByDeleteStatus", userApi.GetUserByDeleteStatus)
		user.GET("/getUserByDeleteReason", userApi.GetUserByDeleteReason)
	}
}
