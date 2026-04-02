package admin

import service "blog_backend_go/services"

type ApiGroup struct {
	AuthApi
	ArticleApi
	CommentApi
	MessageApi
	FriendshipApi
	ScheduleApi
	PermissionApi
	RoleApi
	UserApi
	StatisticsApi
}

var (
	authApi       = service.ServiceGroupApp.AdminServiceGroup.AuthService
	articleApi    = service.ServiceGroupApp.AdminServiceGroup.ArticleService
	commentApi    = service.ServiceGroupApp.AdminServiceGroup.CommentService
	messageApi    = service.ServiceGroupApp.AdminServiceGroup.MessageService
	friendshipApi = service.ServiceGroupApp.AdminServiceGroup.FriendshipService
	scheduleApi   = service.ServiceGroupApp.AdminServiceGroup.ScheduleService
	permissionApi = service.ServiceGroupApp.AdminServiceGroup.PermissionService
	roleApi       = service.ServiceGroupApp.AdminServiceGroup.RoleService
	userApi       = service.ServiceGroupApp.AdminServiceGroup.UserService
	statisticsApi = service.ServiceGroupApp.AdminServiceGroup.StatisticsService
)
