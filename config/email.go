package config

type Email struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Secure   bool   `mapstructure:"secure" json:"secure" yaml:"secure"`
	AuthUser string `mapstructure:"auth-user" json:"auth-user" yaml:"auth-user"`
	AuthPass string `mapstructure:"auth-pass" json:"auth-pass" yaml:"auth-pass"`
}
