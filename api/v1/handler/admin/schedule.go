package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ScheduleApi struct {
}

func (s *ScheduleApi) GetSchedule(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetSchedule"})
}
func (s *ScheduleApi) AddSchedule(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "AddSchedule"})
}
func (s *ScheduleApi) UpdateSchedule(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "UpdateSchedule"})
}
func (s *ScheduleApi) DeleteSchedule(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteSchedule"})
}
