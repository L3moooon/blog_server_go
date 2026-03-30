package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllLink(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get article list",
	})
}

func ApplyForLink(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get article",
	})
}
