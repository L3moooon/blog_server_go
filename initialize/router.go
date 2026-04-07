package initialize

import (
	"blog_backend_go/global"
	router "blog_backend_go/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化总路由
func InitRouter() *gin.Engine {
	Router := gin.New()
	// 使用自定义的 Recovery 中间件，记录 panic 并入库
	// Router.Use(middleware.GinRecovery(true))
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	adminRouter := router.RouterGroupApp.Admin
	webRouter := router.RouterGroupApp.Web

	PublicGroup := Router.Group(global.CONFIG.System.RouterPrefix)
	PrivateGroup := Router.Group(global.CONFIG.System.RouterPrefix)

	// PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())

	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		adminRouter.InitAuthRouter(PublicGroup)        // 注册基础功能路由 不做鉴权
		adminRouter.InitUserRouter(PrivateGroup)       // 注册用户路由
		adminRouter.InitArticleRouter(PrivateGroup)    // 注册文章路由
		adminRouter.InitCommentRouter(PrivateGroup)    // 注册评论路由
		adminRouter.InitMessageRouter(PrivateGroup)    // 注册留言路由
		adminRouter.InitFriendshipRouter(PrivateGroup) // 注册友情链接路由
		adminRouter.InitScheduleRouter(PrivateGroup)   // 注册定时任务路由
		adminRouter.InitPermissionRouter(PrivateGroup) // 注册权限路由
		adminRouter.InitRoleRouter(PrivateGroup)       // 注册角色路由
		adminRouter.InitStatisticsRouter(PrivateGroup) // 注册统计路由
	}

	{
		webRouter.InitHomeRouter(PublicGroup)       // 注册首页路由
		webRouter.InitUserRouter(PublicGroup)       // 注册用户路由
		webRouter.InitArticleRouter(PublicGroup)    // 注册文章路由
		webRouter.InitCommentRouter(PublicGroup)    // 注册评论路由
		webRouter.InitMessageRouter(PublicGroup)    // 注册留言路由
		webRouter.InitFriendshipRouter(PublicGroup) // 注册友情链接路由
	}
	global.ROUTERS = Router.Routes()

	global.LOG.Info("路由注册成功")
	return Router
}
