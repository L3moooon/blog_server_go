package admin

import (
	dto "blog_backend_go/internal/dto/admin"
	model "blog_backend_go/internal/model"
)

type MessageService struct {
}

var MessageServiceInstance = new(MessageService)

func (s *MessageService) CreateMessage(u dto.LoginRequest) (userInter model.SysUser, err error) {
	print("111")
	return
}
