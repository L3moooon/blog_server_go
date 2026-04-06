package model

import (
	"gorm.io/gorm"
)

type WebVisit struct {
	gorm.Model
	ID           uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Identify     string `gorm:"type:varchar(45);not null" json:"identify"`
	Name         string `gorm:"type:varchar(45);not null" json:"name"`
	Avatar       string `gorm:"type:varchar(255)" json:"avatar"`
	Ip           string `gorm:"type:varchar(45);not null" json:"ip"`
	Location     string `gorm:"type:varchar(255);not null" json:"location"`
	Agent        string `gorm:"type:varchar(255);not null" json:"agent"`
	VisitedCount int64  `gorm:"type:bigint;not null" json:"visited_count"`
}

func (WebVisit) TableName() string {
	return "web_visit"
}
