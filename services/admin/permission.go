package admin

import (
	dto "blog_backend_go/internal/dto/admin"
	model "blog_backend_go/internal/model"
)

type PermissionService struct {
}

var PermissionServiceInstance = new(PermissionService)

func (s *PermissionService) CreatePermission(u dto.LoginRequest) (userInter model.SysUser, err error) {
	print("111")
	return
}
