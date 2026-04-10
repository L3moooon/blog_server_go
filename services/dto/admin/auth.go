package admin

import (
	"blog_backend_go/api/v1/dto/admin/auth"
	"blog_backend_go/model"
)

type LoginParam struct {
	auth.LoginRequest
	Ip        string `json:"ip"`
	UserAgent string `json:"user_agent"`
}

type PermissionList struct {
	// 用户表字段
	model.SysUser

	// 角色ID（来自用户角色表）
	RoleName string `gorm:"column:role_name"`

	// 权限表字段（来自角色权限表 + 权限表）
	Name      string `gorm:"column:name"`
	Type      string `gorm:"column:type"`      // 权限标识如 user:list
	Path      string `gorm:"column:path"`      // 路由
	Component string `gorm:"column:component"` // 组件
	Code      string `gorm:"column:code"`      // 按钮标识

}


