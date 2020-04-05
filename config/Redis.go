package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Redis struct {
	Addr           string `yaml:"Addr"`
	Port           string `yaml:"Port"`
	PassWord       string `yaml:"PassWord"`
	DataBaseNumber int    `yaml:"DataBaseNumber"`
}

func (redis *Redis) DefaultRedisConfig() {
	redis.Addr = "localhost"
	redis.Port = "6379"
	redis.DataBaseNumber = 0
	redis.PassWord = ""
}

func (redis *Redis) InitRedisConfig(path string) {
	redis.DefaultRedisConfig()
	file, _ := ioutil.ReadFile(path)
	if err := yaml.Unmarshal(file, redis); err != nil {
		log.Println("ERROR", err)
	}
}
