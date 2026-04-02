package admin

import (
	"github.com/gin-gonic/gin"
)

type ScheduleRouter struct{}

var ScheduleRouterApp = new(ScheduleRouter)

func (s *ScheduleRouter) InitScheduleRouter(Router *gin.RouterGroup) {
	schedule := Router.Group("/schedule")
	{
		schedule.GET("/getScheduleList", scheduleApi.GetScheduleList)
		schedule.POST("/addSchedule", scheduleApi.AddSchedule)
		schedule.PUT("/updateSchedule", scheduleApi.UpdateSchedule)
		schedule.DELETE("/deleteSchedule", scheduleApi.DeleteSchedule)
	}

}
