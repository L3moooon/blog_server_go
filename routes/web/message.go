package web

import (
	"github.com/gin-gonic/gin"
)

type MessageRouter struct{}

var MessageRouterApp = new(MessageRouter)

func (s *MessageRouter) InitMessageRouter(Router *gin.RouterGroup) {

	message := Router.Group("message")
	{
		message.GET("/getAllMessage", messageApi.GetAllMessage) //获取留言
		message.POST("/addMessage", messageApi.AddMessage)      //发表留言
	}

}
