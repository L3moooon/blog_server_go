package admin

import (
	"github.com/gin-gonic/gin"
)

type FriendshipRouter struct{}

var FriendshipRouterApp = new(FriendshipRouter)

func (s *FriendshipRouter) InitFriendshipRouter(Router *gin.RouterGroup) {
	friendship := Router.Group("friendship")
	{
		friendship.GET("/", friendshipApi.GetFriendship)
		friendship.DELETE("/", friendshipApi.DeleteFriendship)
	}

}
