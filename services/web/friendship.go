package web

import (
	dto "blog_backend_go/internal/dto/admin"
	model "blog_backend_go/internal/model"
)

type FriendshipService struct {
}

var FriendshipServiceInstance = new(FriendshipService)

func (s *FriendshipService) CreateFriendship(u dto.LoginRequest) (userInter model.SysUser, err error) {
	print("111")
	return
}
