package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MessageApi struct {
}

func (s *MessageApi) GetMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetMessage"})
}
func (s *MessageApi) DeleteMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteMessage"})
}
