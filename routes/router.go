package router

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use()

	// WebRouterGroup(r)
	// PublicRouterGroup(r)
	AdminRouterGroup(r)

	return r
}
