package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentApi struct {
}

func (s *CommentApi) GetComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetCommentList"})
}
func (s *CommentApi) DeleteComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteComment"})
}
