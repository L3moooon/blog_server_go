package web

import service "blog_backend_go/services"

type ApiGroup struct {
	HomeApi
	UserApi
	ArticleApi
	CommentApi
	MessageApi
	FriendshipApi
}

var (
	homeApi       = service.ServiceGroupApp.WebServiceGroup.HomeService
	userApi       = service.ServiceGroupApp.WebServiceGroup.UserService
	articleApi    = service.ServiceGroupApp.WebServiceGroup.ArticleService
	commentApi    = service.ServiceGroupApp.WebServiceGroup.CommentService
	messageApi    = service.ServiceGroupApp.WebServiceGroup.MessageService
	friendshipApi = service.ServiceGroupApp.WebServiceGroup.FriendshipService
)
