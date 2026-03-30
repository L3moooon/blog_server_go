package web

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAllMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetAllMessage"})
}

func AddMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "AddMessage"})
}
