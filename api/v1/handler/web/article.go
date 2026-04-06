package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleApi struct {
}

func (s *ArticleApi) GetArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get article list",
	})
}

func (s *ArticleApi) GetAllComments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get article",
	})
}

func (s *ArticleApi) View(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "create article",
	})
}

func (s *ArticleApi) Comment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "update article",
	})
}

func (s *ArticleApi) DelComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "delete article",
	})
}
