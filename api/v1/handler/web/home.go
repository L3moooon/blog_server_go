package web

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type HomeApi struct {
}

func (s *HomeApi) GetHomeArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetHomeArticle"})
}

func (s *HomeApi) Info(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Info"})
}

func (s *HomeApi) GetRecommendArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetRecommendArticle"})
}

func (s *HomeApi) GetTagCloud(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetTagCloud"})
}

func (s *HomeApi) Search(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Search"})
}
