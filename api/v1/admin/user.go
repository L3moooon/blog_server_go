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

func (s *UserApi) GetUserById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetUserById"})
}

func GetUserByUsername(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetUserByUsername"})
}

func GetUserByEmail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetUserByEmail"})
}

func GetUserByPhone(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetUserByPhone"})
}

func GetUserByRole(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetUserByRole"})
}

func GetUserByStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetUserByStatus"})
}

func GetUserByCreateTime(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetUserByCreateTime"})
}

func GetUserByUpdateTime(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetUserByUpdateTime"})
}

func GetUserByDeleteTime(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetUserByDeleteTime"})
}

func GetUserByDeleteStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetUserByDeleteStatus"})
}

func GetUserByDeleteReason(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetUserByDeleteReason"})
}
