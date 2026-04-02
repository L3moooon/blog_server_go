package config

type System struct {
	Env                string `mapstructure:"env" json:"env" yaml:"env"`
	Port               int    `mapstructure:"port" json:"port" yaml:"port"` // 端口值
	LimitCountIP       int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP        int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
	RouterPrefix       string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	UseStrictAuth      bool   `mapstructure:"use-strict-auth" json:"use-strict-auth" yaml:"use-strict-auth"`                // 使用树形角色分配模式
	DisableAutoMigrate bool   `mapstructure:"disable-auto-migrate" json:"disable-auto-migrate" yaml:"disable-auto-migrate"` // 自动迁移数据库表结构，生产环境建议设为false，手动迁移
}
