package model

import (
	"gorm.io/gorm"
)

type SysRole struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	RoleName    string `gorm:"type:varchar(45);not null" json:"role_name"`
	RoleCode    string `gorm:"type:varchar(45);not null" json:"role_code"`
	Description string `gorm:"type:varchar(255);not null" json:"description"`
	Status      bool   `gorm:"type:tinyint;not null" json:"status"`
}

func (SysRole) TableName() string {
	return "sys_role"
}
