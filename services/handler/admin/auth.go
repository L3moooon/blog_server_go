package admin

import (
	"blog_backend_go/api/v1/dto/admin/auth"
	"blog_backend_go/global"
	"blog_backend_go/model"
	"blog_backend_go/pkg/utils"
	"blog_backend_go/services/dto/admin"
	"context"
	"errors"
	"time"

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
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("查询用户失败：" + err.Error())
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
			UserName: u.Username,
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
	if err == nil {
		//更新用户ip信息
		global.DB.Model(&model.SysUser{}).Where("id = ?", userInter.ID).Updates(map[string]interface{}{
			"ip":    u.Ip,
			"agent": u.UserAgent,
		})
		//生成并下发token

	}
	return userInter, err
}

// 账号密码登录
func (s *AuthService) LoginByAccount(u admin.LoginParam) (userInter *model.SysUser, err error) {
	var user model.SysUser
	var verifyRule = utils.Rules{
		"Account":  {utils.NotEmpty()}, // 账号不能为空
		"Password": {utils.NotEmpty()}, // 密码不能为空
	}
	err = utils.Verify(u, verifyRule)
	if err != nil {
		return nil, err
	}

	err = global.DB.Where("account = ?", u.Account).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
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
	code, err := global.REDIS.Get(context.Background(), u.Email).Result()
	if err != nil {
		return nil, errors.New("验证码已过期")
	}
	if code != u.Code {
		return nil, errors.New("验证码错误")
	}

	err = global.DB.Where("account = ?", u.Email).First(&user).Error
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
	//重新查询
	err = global.DB.Where("account = ?", u.Email).First(&user).Error
	if err != nil {
		return nil, errors.New("注册成功，但查询用户信息失败")
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
	//验证手机验证码
	code, err := global.REDIS.Get(context.Background(), u.Phone).Result()
	if err != nil {
		return nil, errors.New("验证码已过期")
	}
	if code != u.Code {
		return nil, errors.New("验证码错误")
	}
	//判断手机号是否注册
	err = global.DB.Where("account = ?", u.Phone).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//手机号未注册，使用手机号注册用户
		tx := global.DB.Begin() //开启事务
		id := uuid.New()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback() //回滚
			}
		}()
		err = tx.Create(&model.SysUser{
			ID:       id,
			Account:  u.Phone,
			UserName: u.Phone,
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
	//重新查询
	err = global.DB.Where("account = ?", u.Phone).First(&user).Error
	if err != nil {
		return nil, errors.New("注册成功，但查询用户信息失败")
	}
	return &user, err
}

// 获取邮箱验证码
func (s *AuthService) GetEmailCaptcha(u auth.EmailCaptchaRequest) (err error) {
	code := utils.GenerateCode(6)
	// 将验证码存入redis, 设置5分钟过期
	err = global.REDIS.Set(context.Background(), u.Email, code, 5*time.Minute).Err()
	if err != nil {
		return err
	}
	err = utils.SendEmail(u.Email, "验证码", "您的验证码是: "+code)
	return err
}

// 获取手机验证码
func (s *AuthService) GetSmsCaptcha(u auth.SmsCaptchaRequest) (err error) {
	code := utils.GenerateCode(6)
	// 将验证码存入redis, 设置5分钟过期
	err = global.REDIS.Set(context.Background(), u.Phone, code, 5*time.Minute).Err()
	if err != nil {
		return err
	}
	err = utils.SendSms(u.Phone, code)
	return err

}

// 忘记密码-邮件验证
func (s *AuthService) ForgetPassword(u auth.EmailCaptchaRequest) (err error) {
	var user model.SysUser
	err = global.DB.Where("account = ?", u.Email).First(&user).Error
	if err != nil {
		return errors.New("邮箱不存在")
	}
	code := utils.GenerateCode(6)
	global.REDIS.Set(context.Background(), u.Email, code, 5*time.Minute)
	utils.SendEmail(u.Email, "验证码", "您的验证码是: "+code)
	return err

}

// 重置密码
func (s *AuthService) ResetPassword(u admin.LoginParam) {

}
