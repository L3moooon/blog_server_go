package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentApi struct {
}

func (s *CommentApi) GetCommentList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetCommentList"})
}
