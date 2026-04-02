package web

import (
	"github.com/gin-gonic/gin"
)

type HomeRouter struct{}

var HomeRouterApp = new(HomeRouter)

func (s *HomeRouter) InitHomeRouter(Router *gin.RouterGroup) {

	home := Router.Group("home")
	{
		home.GET("/getHomeArticle", homeApi.GetHomeArticle)           //通用首页文章列表
		home.GET("/info", homeApi.Info)                               //获取网站运转信息
		home.GET("/getRecommendArticle", homeApi.GetRecommendArticle) //获取推荐文章
		home.GET("/getTagCloud", homeApi.GetTagCloud)                 //获取标签云
		home.GET("/search", homeApi.Search)                           //搜索文章
	}

}
