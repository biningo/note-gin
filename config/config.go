package config

import (
	"flag"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var Conf Config

func SetUp() {
	var ConfigPath string
	flag.StringVar(&ConfigPath, "ConfigPath", "", "配置文件路径")
	Conf = NewConfig(ConfigPath)
}

func NewConfig(path string) Config {
	var config Config
	config.cfp = NewconfigFilePath(path)

	config.RedisConfig.InitRedisConfig(config.cfp.RedisPath)
	config.AppConfig.InitAppConfig(config.cfp.AppPath)
	config.ServerConfig.InitServerConfig(config.cfp.ServerPath)
	config.MySqlConfig.InitmySqlConfig(config.cfp.MySqlPath)

	return config
}

type Config struct {
	cfp          configFilePath
	RedisConfig  Redis
	MySqlConfig  MySql
	ServerConfig Server
	AppConfig    App
}

type configFilePath struct {
	ServerPath string `yaml:"ServerPath"`
	MySqlPath  string `yaml:"MySqlPath"`
	RedisPath  string `yaml:"RedisPath"`
	AppPath    string `yaml:"AppPath"`
}

func NewconfigFilePath(path string) configFilePath {
	cfp := configFilePath{}
	file, _ := ioutil.ReadFile(path)
	if err := yaml.Unmarshal(file, &cfp); err != nil {
		log.Println("ERROR", err)
	}
	return cfp
}
