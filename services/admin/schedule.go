package admin

import (
	dto "blog_backend_go/internal/dto/admin"
	model "blog_backend_go/internal/model"
)

type ScheduleService struct {
}

var ScheduleServiceInstance = new(ScheduleService)

func (s *ScheduleService) CreateSchedule(u dto.LoginRequest) (userInter model.SysUser, err error) {
	print("111")
	return
}
