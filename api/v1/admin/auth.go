package admin

import (
	"blog_backend_go/config"
	dto "blog_backend_go/internal/dto/admin"
	"blog_backend_go/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthApi struct{}

func (a *AuthApi) Login(c *gin.Context) {
	var user dto.LoginRequest
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}
	//根据登录类型进行判断
	switch user.LoginType {
	case 1: //用户名密码登录
		if user.Account == "" || user.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "username and password are required",
			})
			return
		}
		//查询用户信息
		var adminUser model.SysUser
		err = config.DB.Where("account = ?", user.Account).First(&adminUser).Error
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "user not found",
			})
			return
		}
		// 校验密码
		if adminUser.Password != user.Password {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "password error",
			})
			return
		}
	case 2: //邮箱验证码登录
		if user.Email == "" || user.Code == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "email and code are required",
			})
			return
		}
	case 3: //手机验证码登录
		if user.Phone == "" || user.Code == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "phone and code are required",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
	})
}

func (a *AuthApi) Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "register",
	})
}

func (a *AuthApi) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "logout",
	})
}

func (a *AuthApi) ForgetPassword(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ForgetPassword"})
}

func (a *AuthApi) ResetPassword(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ResetPassword"})
}

func (a *AuthApi) GetEmailCaptcha(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetEmailCaptcha"})
}

func (a *AuthApi) GetSmsCaptcha(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetSmsCaptcha"})
}
