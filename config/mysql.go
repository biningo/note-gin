package config

type MySql struct {
	Addr      string `yaml:"addr"`
	Port      string `yaml:"port"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	DB        string `yaml:"db"`
	Collation string `yaml:"collation"`
}
