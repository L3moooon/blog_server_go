package admin

import (
	"blog_backend_go/api/v1/dto/admin/auth"
	"blog_backend_go/global"
	"blog_backend_go/model"
	"blog_backend_go/pkg/utils"
	"blog_backend_go/services/dto/admin"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthService struct {
}

var AuthServiceInstance = new(AuthService)

// 账号密码注册
func (s *AuthService) Register(u auth.RegisterRequest) (err error) {
	var user model.SysUser
	err = global.DB.Where("account = ?", u.Account).First(&user).Error
	if err == nil {
		return errors.New("该用户已注册")
	}
	// 开启事务
	{
		tx := global.DB.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback() //回滚
			}
		}()
		id := uuid.New()
		//创建用户
		err = tx.Create(&model.SysUser{
			ID:       id,
			Account:  u.Account,
			Password: utils.BcryptHash(u.Password),
		}).Error
		if err != nil {
			tx.Rollback() //回滚
			return err
		}
		//分配默认角色 4-游客
		err = tx.Create(&model.SysUserRole{
			UserID: id,
			RoleID: 4,
		}).Error
		if err != nil {
			tx.Rollback() //回滚
			return err
		}

		tx.Commit() //提交
	}

	return err
}

// 登录
func (s *AuthService) Login(u admin.LoginParam) (userInter *model.SysUser, err error) {
	switch u.LoginType {
	case "account":
		userInter, err = s.LoginByAccount(u)
	case "email":
		userInter, err = s.LoginByEmail(u)
	case "phone":
		userInter, err = s.LoginByPhone(u)
	default:
		return nil, errors.New("不支持的登录方式")
	}
	if err != nil {
		//下发token
	}
	return userInter, err
}

// 账号密码登录
func (s *AuthService) LoginByAccount(u admin.LoginParam) (userInter *model.SysUser, err error) {
	var user model.SysUser
	var verifyRule = utils.Rules{
		"Account":  {utils.NotEmpty(), utils.RegexpMatch("^[a-zA-Z0-9_]{3,15}$")}, // 账号不能为空
		"Password": {utils.NotEmpty(), utils.RegexpMatch("^[a-zA-Z0-9_]{6,18}$")}, // 密码不能为空
	}
	err = utils.Verify(u, verifyRule)
	if err != nil {
		return nil, err
	}

	err = global.DB.Where("account = ?", u.Account).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(user.Password, u.Password); !ok {
			return nil, errors.New("密码错误")
		}
	} else {
		return nil, errors.New("账号不存在")
	}

	return &user, err
}

// 邮箱验证码登录
func (s *AuthService) LoginByEmail(u admin.LoginParam) (userInter *model.SysUser, err error) {
	var user model.SysUser

	var verifyRule = utils.Rules{
		"Email": {utils.NotEmpty()}, // 邮箱不能为空
		"Code":  {utils.NotEmpty()}, // 验证码不能为空
	}
	err = utils.Verify(u, verifyRule)
	if err != nil {
		return nil, err
	}
	//从redis验证邮箱验证码

	err = global.DB.Where("email = ?", u.Email).First(&user).Error
	//邮箱不存在，注册邮箱
	if err != nil {
		tx := global.DB.Begin() //开启事务
		id := uuid.New()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback() //回滚
			}
		}()
		err = tx.Create(&model.SysUser{
			ID:      id,
			Account: u.Email,
		}).Error
		if err != nil {
			tx.Rollback() //回滚
			return nil, err
		}
		//分配默认角色 4-游客
		err = tx.Create(&model.SysUserRole{
			UserID: id,
			RoleID: 4,
		}).Error
		if err != nil {
			tx.Rollback() //回滚
			return nil, err
		}
		tx.Commit() //提交
	}

	return &user, err
}

// 手机验证码登录
func (s *AuthService) LoginByPhone(u admin.LoginParam) (userInter *model.SysUser, err error) {
	var user model.SysUser

	var verifyRule = utils.Rules{
		"Phone": {utils.NotEmpty()}, // 手机号不能为空
		"Code":  {utils.NotEmpty()}, // 验证码不能为空
	}
	err = utils.Verify(u, verifyRule)
	if err != nil {
		return nil, err
	}

	if !errors.Is(global.DB.Where("phone = ?", u.Phone).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		//手机号未注册，使用手机号注册用户
		err = global.DB.Create(&model.SysUser{
			ID:      uuid.New(),
			Account: u.Phone,
		}).Error
		if err != nil {
			return nil, err
		}
	}
	//验证手机验证码

	return &user, err
}

// 获取邮箱验证码
func (s *AuthService) GetEmailCaptcha(u auth.EmailCaptchaRequest) (userInter *auth.EmailCaptchaResponse, err error) {
	code := utils.GenerateCode(6)
	utils.SendEmail(u.Email, "验证码", "您的验证码是: "+code)
	return &auth.EmailCaptchaResponse{Code: code}, err
}

// 获取手机验证码
func (s *AuthService) GetSmsCaptcha(u admin.LoginParam) {

}

// 忘记密码-发送邮件验证码
func (s *AuthService) ForgetPassword(u admin.LoginParam) {

}

// 重置密码
func (s *AuthService) ResetPassword(u admin.LoginParam) {

}
