package admin

import (
	"blog_backend_go/api/v1/dto/admin/auth"
	"blog_backend_go/global"
	"blog_backend_go/model"
	"blog_backend_go/services/dto/admin"
	"blog_backend_go/utils"
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
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
func (s *AuthService) Login(u admin.LoginParam) (userInter *auth.LoginResponse, err error) {
	var user *model.SysUser
	switch u.LoginType {
	case "account":
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

	case "email":
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
	case "phone":
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
				tx.Rollback()
				return nil, err
			}
			//分配默认角色 4-游客
			err = tx.Create(&model.SysUserRole{
				UserID: id,
				RoleID: 4,
			}).Error
			if err != nil {
				tx.Rollback()
				return nil, err
			}

			tx.Commit() //提交
		}
	default:
		return nil, errors.New("不支持的登录方式")
	}
	var permissionArr []admin.PermissionList
	//查询权限
	err = global.DB.
		Table("sys_user").
		Joins("JOIN sys_user_role sur ON sur.user_id = sys_user.id"). //用户角色表
		Joins("JOIN sys_role sr ON sr.id = sur.role_id").             //角色表
		Joins("JOIN sys_role_permission srp ON srp.role_id = sr.id"). //角色权限表
		Joins("JOIN sys_permission sp ON sp.id = srp.permission_id"). //权限表
		Select("sys_user.id, sys_user.account, sys_user.user_name, sr.role_name, sp.name, sp.type, sp.path, sp.component, sp.code").
		Where("sys_user.id = ?", user.ID).
		Distinct().
		Find(&permissionArr).Error
	permissions := auth.Permission{}
	permissions.Routes = make([]string, 0) // 初始化为空切片
	permissions.Components = make([]string, 0)
	permissions.Btns = make([]string, 0)
	for _, p := range permissionArr {
		switch p.Type {
		case "1":
			permissions.Routes = append(permissions.Routes, p.Path)
		case "2":
			permissions.Components = append(permissions.Components, p.Component)
		case "3":
			permissions.Btns = append(permissions.Btns, p.Code)
		default:
			global.LOG.Error("未知权限类型: %s", zap.String("type", p.Type))
		}
	}
	log.Println(permissions)
	if err != nil {
		return nil, errors.New("查询权限失败")
	}
	if err != nil {
		return nil, err
	}

	//更新用户ip信息
	global.DB.Model(&model.SysUser{}).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"ip":    u.Ip,
		"agent": u.UserAgent,
	})
	//生成并下发token
	jwt := utils.JWT{}
	claims := jwt.CreateClaims(admin.BaseClaims{
		ID:          user.ID,
		Account:     user.Account,
		UserName:    user.UserName,
		Permissions: permissions,
	})
	token, err := jwt.CreateToken(claims)
	if err != nil {
		global.LOG.Error("生成token失败!", zap.Error(err))
		return nil, err
	}
	//如果
	if !global.CONFIG.System.UseMultiLogin {
		return &auth.LoginResponse{
			Token:       token,
			User:        *user,
			ExpiresTime: time.Now().Unix() + claims.RegisteredClaims.ExpiresAt.Unix()*1000,
			Permissions: permissions,
		}, err
	}
	//TODO 多点登录
	return nil, nil
}

// 获取邮箱验证码
func (s *AuthService) GetEmailCaptcha(u auth.EmailCaptchaRequest) (err error) {
	code := utils.GenerateCode(6)
	// 将验证码存入redis, 设置5分钟过期
	rediskey := utils.LoginEmailCaptchaKey(u.Email)
	err = global.REDIS.Set(context.Background(), rediskey, code, 5*time.Minute).Err()
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
	rediskey := utils.LoginSmsCaptchaKey(u.Phone)
	err = global.REDIS.Set(context.Background(), rediskey, code, 5*time.Minute).Err()
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
	rediskey := utils.ResetEmailCaptchaKey(u.Email)
	global.REDIS.Set(context.Background(), rediskey, code, 5*time.Minute)
	utils.SendEmail(u.Email, "验证码", "您的验证码是: "+code)
	return err
}

// 忘记密码-重置密码
func (s *AuthService) ResetPassword(u auth.ResetPasswordRequest) (err error) {
	//验证验证码
	redisKey := utils.ResetEmailCaptchaKey(u.Email)
	code, err := global.REDIS.Get(context.Background(), redisKey).Result()
	if err != nil {
		return errors.New("验证码已过期")
	}
	if code != u.Code {
		return errors.New("验证码错误")
	}
	//更新密码
	err = global.DB.Model(&model.SysUser{}).Where("account = ?", u.Email).Update("password", utils.BcryptHash(u.Password)).Error
	if err != nil {
		return errors.New("更新密码失败")
	}
	return nil
}
