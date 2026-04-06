package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TagApi struct {
}

func (s *TagApi) GetTag(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetTag"})
}

func (s *TagApi) AddTag(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "AddTag"})
}

func (s *TagApi) UpdateTag(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "UpdateTag"})
}

func (s *TagApi) DeleteTag(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteTag"})
}
