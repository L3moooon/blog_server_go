package admin

import (
	dto "blog_backend_go/internal/dto/admin"
	model "blog_backend_go/internal/model"
)

type UserService struct {
}

var UserServiceInstance = new(UserService)

func (s *UserService) CreateUser(u dto.LoginRequest) (userInter model.SysUser, err error) {
	print("111")
	return
}
