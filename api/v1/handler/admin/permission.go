package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PermissionApi struct {
}

func (s *PermissionApi) GetPermission(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetPermission"})
}
func (s *PermissionApi) AddPermission(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "AddPermission"})
}
func (s *PermissionApi) UpdatePermission(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "UpdatePermission"})
}
func (s *PermissionApi) DeletePermission(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeletePermission"})
}
