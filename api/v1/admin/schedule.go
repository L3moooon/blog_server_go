package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ScheduleApi struct {
}

func (s *ScheduleApi) GetScheduleList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetScheduleList"})
}
