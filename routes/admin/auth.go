package admin

import (
	"github.com/gin-gonic/gin"
)

type AuthRouter struct{}

var AuthRouterApp = new(AuthRouter)

func (s *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {
	auth := Router.Group("auth")
	{
		auth.POST("/register", authApi.Register)
		auth.POST("/login", authApi.Login)
		auth.POST("/forgetPassword", authApi.ForgetPassword)
		auth.POST("/resetPassword", authApi.ResetPassword)
		auth.POST("/getEmailCaptcha", authApi.GetEmailCaptcha)
		auth.POST("/getSmsCaptcha", authApi.GetSmsCaptcha)
	}

}
