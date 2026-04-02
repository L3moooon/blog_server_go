package web

import (
	"github.com/gin-gonic/gin"
)

type CommentRouter struct{}

var CommentRouterApp = new(CommentRouter)

func (s *CommentRouter) InitCommentRouter(Router *gin.RouterGroup) {
	comment := Router.Group("/comment")
	{
		comment.POST("/comment", commentApi.Comment)          //发表评论
		comment.DELETE("/delComment", commentApi.DelComment)  //删除评论
	}

}
