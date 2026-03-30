package web

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Visited(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Visited"})
}

func TrackInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TrackInfo"})
}

func ModifyInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ModifyInfo"})
}
