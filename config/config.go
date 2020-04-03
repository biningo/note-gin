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
	Conf=NewConfig(ConfigPath)
}
func NewConfig(path string) Config{
	var config Config
	config.cfp=NewconfigFilePath(path)
	config.SystemConfig.InitSystemConfig(config.cfp.SystemPath)
	return config
}

type Config struct {
	cfp          configFilePath
	SystemConfig System
	MySqlConfig  MySql
	RedisConfig  Redis
}

type configFilePath struct {
	SystemPath string `yaml:"SystemPath"`
	MySqlPath  string `yaml:"MySqlPath"`
	RedisPath  string `yaml:"RedisPath"`
}

func NewconfigFilePath(path string) configFilePath{
	cfp:=configFilePath{}
	file,_:=ioutil.ReadFile(path)
	if err:=yaml.Unmarshal(file,&cfp);err!=nil{
		log.Println("ERROR",err)
	}
	return cfp
}






