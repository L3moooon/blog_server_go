package admin

import (
	"github.com/gin-gonic/gin"
)

type ArticleRouter struct{}

var ArticleRouterApp = new(ArticleRouter)

func (s *ArticleRouter) InitArticleRouter(Router *gin.RouterGroup) {
	article := Router.Group("/article")
	{
		article.GET("/getArticleList", articleApi.GetArticleList)
		article.POST("/addArticle", articleApi.AddArticle)
		article.PUT("/updateArticle", articleApi.UpdateArticle)
		article.DELETE("/deleteArticle", articleApi.DeleteArticle)
		article.POST("/addTag", articleApi.AddTag)
		article.DELETE("/deleteTag", articleApi.DeleteTag)
	}

}
