package router

import "github.com/gin-gonic/gin"

func PublicRouterGroup(r *gin.Engine) {
	public := r.Group("/public")
	{
		public.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "public",
			})
		})
	}
}
