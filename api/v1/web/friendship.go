package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FriendshipApi struct {
}

func (s *FriendshipApi) GetAllLink(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get article list",
	})
}

func (s *FriendshipApi) ApplyForLink(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get article",
	})
}
