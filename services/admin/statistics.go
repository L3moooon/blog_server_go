package admin

import (
	dto "blog_backend_go/internal/dto/admin"
	model "blog_backend_go/internal/model"
)

type StatisticsService struct {
}

var StatisticsServiceInstance = new(StatisticsService)

func (s *StatisticsService) CreateStatistics(u dto.LoginRequest) (userInter model.SysUser, err error) {
	print("111")
	return
}
