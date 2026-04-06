package admin

import service "blog_backend_go/services"

type ApiGroup struct {
	AuthApi
	ArticleApi
	TagApi
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
	authService       = service.ServiceGroupApp.AdminServiceGroup.AuthService
	articleService    = service.ServiceGroupApp.AdminServiceGroup.ArticleService
	tagService        = service.ServiceGroupApp.AdminServiceGroup.TagService
	commentService    = service.ServiceGroupApp.AdminServiceGroup.CommentService
	messageService    = service.ServiceGroupApp.AdminServiceGroup.MessageService
	friendshipService = service.ServiceGroupApp.AdminServiceGroup.FriendshipService
	scheduleService   = service.ServiceGroupApp.AdminServiceGroup.ScheduleService
	permissionService = service.ServiceGroupApp.AdminServiceGroup.PermissionService
	roleService       = service.ServiceGroupApp.AdminServiceGroup.RoleService
	userService       = service.ServiceGroupApp.AdminServiceGroup.UserService
	statisticsService = service.ServiceGroupApp.AdminServiceGroup.StatisticsService
)
