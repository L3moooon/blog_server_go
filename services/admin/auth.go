package service

import (
	dto "blog_backend_go/internal/dto/admin"
	model "blog_backend_go/internal/model"
)

type AuthService struct {
}

var AuthServiceInstance = new(AuthService)

func (s *AuthService) Login(u dto.LoginRequest) (userInter model.SysUser, err error) {

}
