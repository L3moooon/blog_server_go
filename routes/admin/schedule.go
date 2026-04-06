package admin

import (
	"github.com/gin-gonic/gin"
)

type ScheduleRouter struct{}

var ScheduleRouterApp = new(ScheduleRouter)

func (s *ScheduleRouter) InitScheduleRouter(Router *gin.RouterGroup) {
	schedule := Router.Group("schedule")
	{
		schedule.GET("/", scheduleApi.GetSchedule)
		schedule.POST("/", scheduleApi.AddSchedule)
		schedule.PUT("/", scheduleApi.UpdateSchedule)
		schedule.DELETE("/", scheduleApi.DeleteSchedule)
	}

}
