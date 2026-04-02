package admin

import (
	dto "blog_backend_go/internal/dto/admin"
	model "blog_backend_go/internal/model"
)

type RoleService struct {
}

var RoleServiceInstance = new(RoleService)

func (s *RoleService) CreateRole(u dto.LoginRequest) (userInter model.SysUser, err error) {
	print("111")
	return
}
