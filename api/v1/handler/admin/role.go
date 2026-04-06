package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleApi struct {
}

func (s *RoleApi) GetRole(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetRole"})
}
func (s *RoleApi) AddRole(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "AddRole"})
}
func (s *RoleApi) UpdateRole(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "UpdateRole"})
}
func (s *RoleApi) DeleteRole(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteRole"})
}
