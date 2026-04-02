package admin

import (
	"blog_backend_go/global"
	dto "blog_backend_go/internal/dto/admin"
	model "blog_backend_go/internal/model"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthService struct {
}

var AuthServiceInstance = new(AuthService)

func (s *AuthService) Login(u dto.LoginRequest) (userInter model.SysUser, err error) {
	var user model.SysUser

	if !errors.Is(global.DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.New()
	err = global.DB.Create(&u).Error
	return u, err
}
