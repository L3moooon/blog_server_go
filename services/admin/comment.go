package admin

import (
	dto "blog_backend_go/internal/dto/admin"
	model "blog_backend_go/internal/model"
)

type CommentService struct {
}

var CommentServiceInstance = new(CommentService)

func (s *CommentService) CreateComment(u dto.LoginRequest) (userInter model.SysUser, err error) {
	print("111")
	return
}
