package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleApi struct {
}

func (s *RoleApi) GetRoleList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetRoleList"})
}
