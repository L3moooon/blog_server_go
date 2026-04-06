package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

func (s *UserApi) GetUserList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetUserList"})
}
