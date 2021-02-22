package config

type Server struct {
	Addr string `yaml:"addr"`
	Port string `yaml:"port"`
	Mode string `yaml:"mode"`
}
