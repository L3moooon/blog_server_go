package web

import (
	"github.com/gin-gonic/gin"
)

type ArticleRouter struct{}

var ArticleRouterApp = new(ArticleRouter)

func (s *ArticleRouter) InitArticleRouter(Router *gin.RouterGroup) {
	article := Router.Group("/article")
	{
		article.GET("/getArticle", articleApi.GetArticle)     //获取文章
		article.GET("/getComment", articleApi.GetAllComments) //获取文章评论
		article.PUT("/view", articleApi.View)                 //更新访问量
	}

}
