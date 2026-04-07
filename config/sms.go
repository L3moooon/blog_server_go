package config

type SMS struct {
	AccessKeyId     string `mapstructure:"access-key-id" json:"access-key-id" yaml:"access-key-id"`
	AccessKeySecret string `mapstructure:"access-key-secret" json:"access-key-secret" yaml:"access-key-secret"`
	SignName        string `mapstructure:"sign-name" json:"sign-name" yaml:"sign-name"`
	TemplateCode    string `mapstructure:"template-code" json:"template-code" yaml:"template-code"`
	TemplateParam   string `mapstructure:"template-param" json:"template-param" yaml:"template-param"`
}
