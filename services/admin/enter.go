package admin

type ServiceGroup struct {
	AuthService
	ArticleService
	CommentService
	MessageService
	FriendshipService
	ScheduleService
	PermissionService
	RoleService
	UserService
	StatisticsService
}
