package admin

type ServiceGroup struct {
	AuthService
	ArticleService
	TagService
	CommentService
	MessageService
	FriendshipService
	ScheduleService
	PermissionService
	RoleService
	UserService
	StatisticsService
}
