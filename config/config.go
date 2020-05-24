package config

import (
	"flag"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var Conf Config

//启动  go run main.go -c config/file/BootLoader.yaml
func SetUp() {
	var ConfigPath string
	flag.StringVar(&ConfigPath, "c", "", "配置文件路径")
	flag.Parse()
	//log.Println(ConfigPath)
	Conf = NewConfig(ConfigPath)
}

//总配置对象
type Config struct {
	//存放各个配置文件的路径 Path
	Cfp configFilePath

	//各个配置模块
	RedisConfig  Redis
	MySqlConfig  MySql
	ServerConfig Server
	AppConfig    App
}

//根据命令行输入的BootLoader.yaml路径  创建Config  同时根据BootLoader.yaml配置的其他文件路径初始化其他模块配置
//Config像一个抓手 找到其他文件路径并初始化配置
func NewConfig(path string) Config {
	var config Config
	config.Cfp = NewconfigFilePath(path)

	config.RedisConfig.InitRedisConfig(config.Cfp.RedisPath)
	config.AppConfig.InitAppConfig(config.Cfp.AppPath)
	config.ServerConfig.InitServerConfig(config.Cfp.ServerPath)
	config.MySqlConfig.InitmySqlConfig(config.Cfp.MySqlPath)

	return config
}

//该struct保存了各个模块的配置文件路径
type configFilePath struct {
	ServerPath string `yaml:"ServerPath"`
	MySqlPath  string `yaml:"MySqlPath"`
	RedisPath  string `yaml:"RedisPath"`
	AppPath    string `yaml:"AppPath"`
}

//创建 configFilePath
func NewconfigFilePath(path string) configFilePath {
	Cfp := configFilePath{}
	file, _ := ioutil.ReadFile(path)
	if err := yaml.Unmarshal(file, &Cfp); err != nil {
		log.Println("ERROR", err)
	}
	return Cfp
}
