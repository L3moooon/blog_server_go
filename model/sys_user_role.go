package model

import "github.com/google/uuid"

type SysUserRole struct {
	ID     uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID uuid.UUID `gorm:"type:varchar(36);not null" json:"user_id"`
	RoleID uint64    `gorm:"type:bigint;not null" json:"role_id"`
}

func (SysUserRole) TableName() string {
	return "sys_user_role"
}
