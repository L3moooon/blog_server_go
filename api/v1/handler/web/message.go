package web

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type MessageApi struct {
}

func (s *MessageApi) GetAllMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetAllMessage"})
}

func (s *MessageApi) AddMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "AddMessage"})
}
