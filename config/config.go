package config
type Config struct {
	MySql  MySql  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
	Server Server `yaml:"server"`
	App    App    `yaml:"app"`
}
