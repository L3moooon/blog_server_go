package model

import "gorm.io/gorm"

type SysPermission struct {
	gorm.Model
	ID             uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	ParentID       uint64    `gorm:"type:bigint;not null" json:"parent_id"`
	PermissionName string    `gorm:"type:varchar(45);not null" json:"permission_name"`
	PermissionType string    `gorm:"type:varchar(45);not null" json:"permission_type"`
	Path           string    `gorm:"type:varchar(45)" json:"path"`
	Component      string    `gorm:"type:varchar(45)" json:"component"`
	PermissionCode string    `gorm:"type:varchar(45)" json:"permission_code"`
	Sort           int64     `gorm:"type:bigint" json:"sort"`
	Status         bool      `gorm:"type:tinyint;not null" json:"status"`
}

func (SysPermission) TableName() string {
	return "sys_permission"
}
