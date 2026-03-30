package admin

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Statistics
func GetStatistics(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetStatistics"})
}

// Schedule management
func GetScheduleList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetScheduleList"})
}

func AddSchedule(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "AddSchedule"})
}

func UpdateSchedule(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "UpdateSchedule"})
}

func DeleteSchedule(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteSchedule"})
}
