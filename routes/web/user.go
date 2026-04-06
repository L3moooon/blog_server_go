package web

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

var UserRouterApp = new(UserRouter)

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	user := Router.Group("user")
	{
		user.GET("/getHomeArticle", userApi.TrackInfo) //通用首页文章列表
	}

}
