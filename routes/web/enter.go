package web

import api "blog_backend_go/api/v1"

type RouterGroup struct {
	HomeRouter
	UserRouter
	ArticleRouter
	CommentRouter
	MessageRouter
	FriendshipRouter
}

var (
	homeApi       = api.ApiGroupApp.WebApiGroup.HomeApi
	userApi       = api.ApiGroupApp.WebApiGroup.UserApi
	articleApi    = api.ApiGroupApp.WebApiGroup.ArticleApi
	commentApi    = api.ApiGroupApp.WebApiGroup.CommentApi
	messageApi    = api.ApiGroupApp.WebApiGroup.MessageApi
	friendshipApi = api.ApiGroupApp.WebApiGroup.FriendshipApi
)
