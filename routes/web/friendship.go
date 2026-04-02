package web

import (
	"github.com/gin-gonic/gin"
)

type FriendshipRouter struct{}

var FriendshipRouterApp = new(FriendshipRouter)

func (s *FriendshipRouter) InitFriendshipRouter(Router *gin.RouterGroup) {

		//友链
		friendship := Router.Group("/friendship")
		{
			friendship.GET("/getAllLink", friendshipApi.GetAllLink)      //获取友链
			friendship.POST("/applyForLink", friendshipApi.ApplyForLink) //发表友链
		}
}
