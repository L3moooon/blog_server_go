package config

type Server struct {
	System    System    `mapstructure:"system" json:"system" yaml:"system"`
	Mysql     Mysql     `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis     Redis     `mapstructure:"redis" json:"redis" yaml:"redis"`
	AliyunOSS AliyunOSS `mapstructure:"aliyun-oss" json:"aliyun-oss" yaml:"aliyun-oss"`
	JWT       JWT       `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap       Zap       `mapstructure:"zap" json:"zap" yaml:"zap"`
	Email     Email     `mapstructure:"email" json:"email" yaml:"email"`
	SMS       SMS       `mapstructure:"sms" json:"sms" yaml:"sms"`
	// Captcha   Captcha   `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// Mongo     Mongo   `mapstructure:"mongo" json:"mongo" yaml:"mongo"`
	// Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`
	// DiskList []DiskList `mapstructure:"disk-list" json:"disk-list" yaml:"disk-list"`
	// // 跨域配置
	// Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
	// // MCP配置
	// MCP MCP `mapstructure:"mcp" json:"mcp" yaml:"mcp"`
}
