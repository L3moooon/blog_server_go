package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleApi struct {
}

func (s *ArticleApi) GetArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetArticle"})
}
func (s *ArticleApi) AddArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "AddArticle"})
}
func (s *ArticleApi) UpdateArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "UpdateArticle"})
}
func (s *ArticleApi) DeleteArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteArticle"})
}
