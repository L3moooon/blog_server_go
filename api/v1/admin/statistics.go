package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatisticsApi struct {
}

func (s *StatisticsApi) GetStatistics(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetStatistics"})
}

func (s *StatisticsApi) GetScheduleList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetScheduleList"})
}

func (s *StatisticsApi) AddSchedule(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "AddSchedule"})
}

func (s *StatisticsApi) UpdateSchedule(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "UpdateSchedule"})
}

func (s *StatisticsApi) DeleteSchedule(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteSchedule"})
}
