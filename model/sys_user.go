package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SysUser struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:char(36);primaryKey;comment:id" json:"id"`
	Account  string    `gorm:"type:varchar(45);not null;comment:账号" json:"account"`
	UserName string    `gorm:"type:varchar(45);not null;comment:用户名称" json:"user_name"`
	Password string    `gorm:"type:char(60);comment:密码" json:"-"`
	Avatar   string    `gorm:"type:varchar(255);comment:头像" json:"avatar"`
	Ip       string    `gorm:"type:varchar(20);not null;comment:ip地址" json:"ip"`
	Agent    string    `gorm:"type:varchar(255);not null;comment:浏览器信息" json:"agent"`
	Location string    `gorm:"type:varchar(255);not null;comment:地址" json:"location"`
	Status   bool      `gorm:"type:tinyint(1);default:1;not null;comment:状态 1启用 0禁用" json:"status"`
}

func (SysUser) TableName() string {
	return "sys_user"
}
