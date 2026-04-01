package model

import "time"

type Login interface {
	GetID() uint64
	GetAccount() string
	GetUserName() string
	GetAvatar() string
	GetIp() string
	GetLocation() string
	GetCreateTime() time.Time
	GetLastLoginTime() time.Time
	GetStatus() bool
}

type SysUser struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement;comment:id" json:"id"`
	Account       string    `gorm:"type:varchar(45);not null;comment:账号" json:"account"`
	UserName      string    `gorm:"type:varchar(45);not null;comment:用户名称" json:"user_name"`
	Password      string    `gorm:"type:varchar(45);not null;comment:密码" json:"-"`
	Avatar        string    `gorm:"type:varchar(255);comment:头像" json:"avatar"`
	Ip            string    `gorm:"type:varchar(45);not null;comment:ip地址" json:"ip"`
	Location      string    `gorm:"type:varchar(255);not null;comment:地址" json:"location"`
	CreateTime    time.Time `gorm:"type:datetime;not null;comment:创建时间" json:"create_time"`
	LastLoginTime time.Time `gorm:"type:datetime;not null;comment:最后登录时间" json:"last_login_time"`
	Status        bool      `gorm:"type:tinyint;not null;comment:状态 1启用 0禁用" json:"status"`
}

func (SysUser) TableName() string {
	return "sys_user"
}

func (s *SysUser) GetID() uint64 {
	return s.ID
}

func (s *SysUser) GetAccount() string {
	return s.Account
}

func (s *SysUser) GetUserName() string {
	return s.UserName
}

func (s *SysUser) GetAvatar() string {
	return s.Avatar
}

func (s *SysUser) GetIp() string {
	return s.Ip
}

func (s *SysUser) GetLocation() string {
	return s.Location
}

func (s *SysUser) GetCreateTime() time.Time {
	return s.CreateTime
}

func (s *SysUser) GetLastLoginTime() time.Time {
	return s.LastLoginTime
}

func (s *SysUser) GetStatus() bool {
	return s.Status
}
