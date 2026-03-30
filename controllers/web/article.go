package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get article list",
	})
}

func GetAllComments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get article",
	})
}

func View(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "create article",
	})
}

func Comment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "update article",
	})
}

func DelComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "delete article",
	})
}
