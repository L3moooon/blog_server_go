package web

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

var UserRouterApp = new(UserRouter)

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	user := Router.Group("user")
	{
		user.GET("/getHomeArticle", userApi.GetHomeArticle)           //通用首页文章列表
		user.GET("/info", userApi.Info)                               //获取网站运转信息
		user.GET("/getRecommendArticle", userApi.GetRecommendArticle) //获取推荐文章
		user.GET("/getTagCloud", userApi.GetTagCloud)                 //获取标签云
		user.GET("/search", userApi.Search)                           //搜索文章
	}

}
