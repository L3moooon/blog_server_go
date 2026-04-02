package config

type Server struct {
	JWT   JWT   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap   Zap   `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	// Mongo     Mongo   `mapstructure:"mongo" json:"mongo" yaml:"mongo"`
	// Email     Email   `mapstructure:"email" json:"email" yaml:"email"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	// Captcha   Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// oss
	AliyunOSS AliyunOSS `mapstructure:"aliyun-oss" json:"aliyun-oss" yaml:"aliyun-oss"`

	// Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`

	// DiskList []DiskList `mapstructure:"disk-list" json:"disk-list" yaml:"disk-list"`

	// // 跨域配置
	// Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`

	// // MCP配置
	// MCP MCP `mapstructure:"mcp" json:"mcp" yaml:"mcp"`
}
