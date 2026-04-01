package router

import (
	"blog_backend_go/api/admin"

	"github.com/gin-gonic/gin"
)

func AdminRouterGroup(r *gin.Engine) {
	adminGroup := r.Group("/admin")
	{
		//登录注册
		auth := adminGroup.Group("/auth")
		{
			auth.POST("/login", admin.Login)
			auth.POST("/register", admin.Register)
			auth.POST("/forgetPassword", admin.ForgetPassword)
			auth.POST("/resetPassword", admin.ResetPassword)
			auth.POST("/getEmailCaptcha", admin.GetEmailCaptcha)
			auth.POST("/getSmsCaptcha", admin.GetSmsCaptcha)
		}
		//用户管理
		// user := adminGroup.Group("/user")
		// {
		// 	user.GET("/getUserList", admin.GetUserList)
		// 	user.GET("/getUserById", admin.GetUserById)
		// 	user.GET("/getUserByUsername", admin.GetUserByUsername)
		// 	user.GET("/getUserByEmail", admin.GetUserByEmail)
		// 	user.GET("/getUserByPhone", admin.GetUserByPhone)
		// 	user.GET("/getUserByRole", admin.GetUserByRole)
		// 	user.GET("/getUserByStatus", admin.GetUserByStatus)
		// 	user.GET("/getUserByCreateTime", admin.GetUserByCreateTime)
		// 	user.GET("/getUserByUpdateTime", admin.GetUserByUpdateTime)
		// 	user.GET("/getUserByDeleteTime", admin.GetUserByDeleteTime)
		// 	user.GET("/getUserByDeleteStatus", admin.GetUserByDeleteStatus)
		// 	user.GET("/getUserByDeleteReason", admin.GetUserByDeleteReason)
		// }
		//角色管理
		role := adminGroup.Group("/role")
		{
			role.GET("/getRoleList", admin.GetRoleList)
			role.POST("/addRole", admin.AddRole)
			role.PUT("/updateRole", admin.UpdateRole)
			role.DELETE("/deleteRole", admin.DeleteRole)
		}
		//权限管理
		permission := adminGroup.Group("/permission")
		{
			permission.GET("/getPermissionList", admin.GetPermissionList)
			permission.POST("/addPermission", admin.AddPermission)
			permission.PUT("/updatePermission", admin.UpdatePermission)
			permission.DELETE("/deletePermission", admin.DeletePermission)
		}
		//统计
		statistics := adminGroup.Group("/statistics")
		{
			statistics.GET("/", admin.GetStatistics)
		}
		//文章管理
		article := adminGroup.Group("/article")
		{
			article.GET("/getArticleList", admin.GetArticleList)
			article.POST("/addArticle", admin.AddArticle)
			article.PUT("/updateArticle", admin.UpdateArticle)
			article.DELETE("/deleteArticle", admin.DeleteArticle)
			article.POST("/addTag", admin.AddTag)
			article.DELETE("/deleteTag", admin.DeleteTag)
		}
		//评论管理
		comment := adminGroup.Group("/comment")
		{
			comment.GET("/getCommentLists", admin.GetCommentList)
			comment.DELETE("/deleteComment", admin.DeleteComment)
		}
		//留言管理
		message := adminGroup.Group("/message")
		{
			message.GET("/getMessageLists", admin.GetMessageList)
			message.DELETE("/deleteMessage", admin.DeleteMessage)
		}
		//友链管理
		friendship := adminGroup.Group("/friendship")
		{
			friendship.GET("/getFriendshipLists", admin.GetFriendshipList)
			friendship.DELETE("/deleteFriendship", admin.DeleteFriendship)
		}
		//定时任务
		schedule := adminGroup.Group("/schedule")
		{
			schedule.GET("/getScheduleList", admin.GetScheduleList)
			schedule.POST("/addSchedule", admin.AddSchedule)
			schedule.PUT("/updateSchedule", admin.UpdateSchedule)
			schedule.DELETE("/deleteSchedule", admin.DeleteSchedule)
		}
	}
}
