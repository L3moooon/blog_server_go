package admin

import (
	"github.com/gin-gonic/gin"
)

type MessageRouter struct{}

var MessageRouterApp = new(MessageRouter)

func (s *MessageRouter) InitMessageRouter(Router *gin.RouterGroup) {
	message := Router.Group("/message")
	{
		message.GET("/getMessageLists", messageApi.GetMessageList)
		message.DELETE("/deleteMessage", messageApi.DeleteMessage)
	}

}
