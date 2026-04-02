package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PermissionApi struct {
}

func (s *PermissionApi) GetPermissionList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetPermissionList"})
}
