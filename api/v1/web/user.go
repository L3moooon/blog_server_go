package web

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

func (s *UserApi) Visited(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Visited"})
}

func (s *UserApi) TrackInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TrackInfo"})
}

func (s *UserApi) ModifyInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ModifyInfo"})
}
