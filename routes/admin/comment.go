package admin

import (
	"github.com/gin-gonic/gin"
)

type CommentRouter struct{}

var CommentRouterApp = new(CommentRouter)

func (s *CommentRouter) InitCommentRouter(Router *gin.RouterGroup) {
	comment := Router.Group("/comment")
	{
		comment.GET("/getCommentLists", commentApi.GetCommentList)
		comment.DELETE("/deleteComment", commentApi.DeleteComment)
	}

}
