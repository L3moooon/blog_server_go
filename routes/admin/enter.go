package admin

import api "blog_backend_go/api/v1"

type RouterGroup struct {
	AuthRouter
	PermissionRouter
	RoleRouter
	UserRouter
	StatisticsRouter
	ArticleRouter
	CommentRouter
	MessageRouter
	FriendshipRouter
	ScheduleRouter
}

var (
	authApi       = api.ApiGroupApp.AdminApiGroup.AuthApi
	permissionApi = api.ApiGroupApp.AdminApiGroup.PermissionApi
	roleApi       = api.ApiGroupApp.AdminApiGroup.RoleApi
	userApi       = api.ApiGroupApp.AdminApiGroup.UserApi
	statisticsApi = api.ApiGroupApp.AdminApiGroup.StatisticsApi
	articleApi    = api.ApiGroupApp.AdminApiGroup.ArticleApi
	commentApi    = api.ApiGroupApp.AdminApiGroup.CommentApi
	messageApi    = api.ApiGroupApp.AdminApiGroup.MessageApi
	friendshipApi = api.ApiGroupApp.AdminApiGroup.FriendshipApi
	scheduleApi   = api.ApiGroupApp.AdminApiGroup.ScheduleApi
)
