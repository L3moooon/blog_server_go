package router

import (
	"blog_backend_go/controllers/web"

	"github.com/gin-gonic/gin"
)

func WebRouterGroup(r *gin.Engine) {
	webGroup := r.Group("/web")
	{
		//首页
		home := webGroup.Group("/home")
		{
			home.GET("/getHomeArticle", web.GetHomeArticle)           //通用首页文章列表
			home.GET("/info", web.Info)                               //获取网站运转信息
			home.GET("/getRecommendArticle", web.GetRecommendArticle) //获取推荐文章
			home.GET("/getTagCloud", web.GetTagCloud)                 //获取标签云
			home.GET("/search", web.Search)                           //搜索文章
		}
		//用户
		user := webGroup.Group("/user")
		{
			user.GET("/visited", web.Visited)       //访问记录
			user.GET("/trackInfo", web.TrackInfo)   //埋点统计
			user.GET("/modifyInfo", web.ModifyInfo) //修改个人信息
		}
		//文章
		article := webGroup.Group("/article")
		{
			article.GET("/getArticle", web.GetArticle)     //获取文章
			article.GET("/getComment", web.GetAllComments) //获取文章评论
			article.PUT("/view", web.View)                 //更新访问量
			article.POST("/comment", web.Comment)          //发表评论
			article.DELETE("/delComment", web.DelComment)  //删除评论
		}
		//留言
		message := webGroup.Group("/message")
		{
			message.GET("/getAllMessage", web.GetAllMessage) //获取留言
			message.POST("/addMessage", web.AddMessage)      //发表留言
		}
		//友链
		friendship := webGroup.Group("/friendship")
		{
			friendship.GET("/getAllLink", web.GetAllLink)      //获取友链
			friendship.POST("/applyForLink", web.ApplyForLink) //发表友链
		}

	}
}
