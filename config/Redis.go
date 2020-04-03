package config

type Redis struct {
	Addr string	`yaml:"Addr"`
	Port string		`yaml:"Port"`
	UserName string	`yaml:"UserName"`
	PassWord string `yaml:"PassWord"`
	DataBaseName string `yaml:"DataBaseName"`
}
