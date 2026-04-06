package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FriendshipApi struct {
}

func (s *FriendshipApi) GetFriendship(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetFriendship"})
}
func (s *FriendshipApi) DeleteFriendship(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteFriendship"})
}
