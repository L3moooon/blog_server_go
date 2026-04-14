package admin

import (
	"github.com/gin-gonic/gin"
)

type ArticleRouter struct{}

var ArticleRouterApp = new(ArticleRouter)

func (s *ArticleRouter) InitArticleRouter(Router *gin.RouterGroup) {
	article := Router.Group("/article")
	{
		article.GET("/", articleApi.GetArticleList)
		article.POST("/", articleApi.AddArticle)
		article.PUT("/", articleApi.UpdateArticle)
		article.DELETE("/", articleApi.DeleteArticle)
	}

}
