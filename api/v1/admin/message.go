package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MessageApi struct {
}

func (s *MessageApi) GetMessageList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetMessageList"})
}
