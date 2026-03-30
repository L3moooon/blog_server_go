package web

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetHomeArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetHomeArticle"})
}

func Info(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Info"})
}

func GetRecommendArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetRecommendArticle"})
}

func GetTagCloud(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetTagCloud"})
}

func Search(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Search"})
}
