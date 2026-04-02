package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FriendshipApi struct {
}

func (s *FriendshipApi) GetFriendshipList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetFriendshipList"})
}
