package admin

import (
	"github.com/gin-gonic/gin"
)

type TagRouter struct{}

var TagRouterApp = new(TagRouter)

func (s *TagRouter) InitTagRouter(Router *gin.RouterGroup) {
	tag := Router.Group("tag")
	{
		tag.GET("/", tagApi.GetTag)
		tag.POST("/", tagApi.AddTag)
		tag.PUT("/", tagApi.UpdateTag)
		tag.DELETE("/", tagApi.DeleteTag)
	}
}
