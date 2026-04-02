package web

import (
	dto "blog_backend_go/internal/dto/admin"
	model "blog_backend_go/internal/model"
)

type HomeService struct {
}

var HomeServiceInstance = new(HomeService)

func (s *HomeService) CreateHome(u dto.LoginRequest) (userInter model.SysUser, err error) {
	print("111")
	return
}
