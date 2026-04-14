package admin

import (
	"blog_backend_go/api/v1/dto/admin/auth"
	"blog_backend_go/api/v1/dto/common/response"
	"blog_backend_go/global"
	"blog_backend_go/services/dto/admin"
	"blog_backend_go/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthApi struct{}

// 注册
func (a *AuthApi) Register(c *gin.Context) {
	var req auth.RegisterRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}

	// 字段校验
	var verifyRule = utils.Rules{
		"Username": {utils.NotEmpty()}, // 用户名不能为空
		"Account":  {utils.NotEmpty()}, // 账号不能为空
		"Password": {utils.NotEmpty()}, // 密码不能为空
	}
	err = utils.Verify(req, verifyRule)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}

	err = authService.Register(req)
	if err != nil {
		global.LOG.Error("用户注册失败!", zap.Error(err))
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithMessage(c, "注册成功")
}

func (a *AuthApi) Login(c *gin.Context) {
	var req auth.LoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	// 字段校验
	var verifyRule = utils.Rules{
		"LoginType": {utils.NotEmpty()}, // 登录类型不能为空
	}
	err = utils.Verify(req, verifyRule)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}

	res, err := authService.Login(admin.LoginParam{
		LoginRequest: req,
		Ip:           c.ClientIP(),
		UserAgent:    c.GetHeader("User-Agent"),
	})
	if err != nil {
		global.LOG.Error("登录失败!", zap.Error(err))
		response.FailWithMessage(c, err.Error())
		return
	}
	if res.User.Status == false {
		global.LOG.Error("登录失败! 用户被禁止登录!")
		response.FailWithMessage(c, "用户被禁止登录")
		return
	}
	response.OkWithDetailed(c, res, "登录成功")
}

func (a *AuthApi) GetEmailCaptcha(c *gin.Context) {
	var req auth.EmailCaptchaRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	// 字段校验
	var verifyRule = utils.Rules{
		"Email": {utils.NotEmpty()}, // 邮箱不能为空
	}
	err = utils.Verify(req, verifyRule)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err = authService.GetEmailCaptcha(req)
	if err != nil {
		global.LOG.Error("获取邮箱验证码失败!", zap.Error(err))
		response.FailWithMessage(c, "获取邮箱验证码失败")
		return
	}
	response.OkWithMessage(c, "获取邮箱验证码成功")
}

func (a *AuthApi) GetSmsCaptcha(c *gin.Context) {
	var req auth.SmsCaptchaRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	// 字段校验
	var verifyRule = utils.Rules{
		"Phone": {utils.NotEmpty()}, // 手机号不能为空
	}
	err = utils.Verify(req, verifyRule)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err = authService.GetSmsCaptcha(req)
	if err != nil {
		global.LOG.Error("获取短信验证码失败!", zap.Error(err))
		response.FailWithMessage(c, "获取短信验证码失败")
		return
	}
	response.OkWithMessage(c, "获取短信验证码成功")
}

func (a *AuthApi) Logout(c *gin.Context) {
}

func (a *AuthApi) ForgetPassword(c *gin.Context) {
	var req auth.EmailCaptchaRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	// 字段校验
	var verifyRule = utils.Rules{
		"Email": {utils.NotEmpty()}, // 邮箱不能为空
	}
	err = utils.Verify(req, verifyRule)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}

	err = authService.ForgetPassword(req)
	if err != nil {
		global.LOG.Error("忘记密码失败!", zap.Error(err))
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Ok(c)
}

func (a *AuthApi) ResetPassword(c *gin.Context) {
	var req auth.ResetPasswordRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	// 字段校验
	var verifyRule = utils.Rules{
		"Email":    {utils.NotEmpty()}, // 邮箱不能为空
		"Code":     {utils.NotEmpty()}, // 验证码不能为空
		"Password": {utils.NotEmpty()}, // 密码不能为空
	}
	err = utils.Verify(req, verifyRule)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}

	err = authService.ResetPassword(req)
	if err != nil {
		global.LOG.Error("重置密码失败!", zap.Error(err))
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Ok(c)
}
