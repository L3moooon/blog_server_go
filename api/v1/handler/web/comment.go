package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentApi struct {
}

func (s *CommentApi) GetArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get article list",
	})
}

func (s *CommentApi) GetAllComments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get article",
	})
}

func (s *CommentApi) View(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "create article",
	})
}

func (s *CommentApi) Comment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "update article",
	})
}

func (s *CommentApi) DelComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "delete article",
	})
}
