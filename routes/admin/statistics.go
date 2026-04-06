package admin

import (
	"github.com/gin-gonic/gin"
)

type StatisticsRouter struct{}

var StatisticsRouterApp = new(StatisticsRouter)

func (s *StatisticsRouter) InitStatisticsRouter(Router *gin.RouterGroup) {
	statistics := Router.Group("statistics")
	{
		statistics.GET("/", statisticsApi.GetStatistics)
	}

}
