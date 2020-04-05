package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type MySql struct {
	Addr         string `yaml:"Addr"`
	Port         string `yaml:"Port"`
	UserName     string `yaml:"UserName"`
	PassWord     string `yaml:"PassWord"`
	DataBaseName string `yaml:"DataBaseName"`
}

func (mysql *MySql) DefaultmySqlConfig() {
	mysql.Addr = "localhost"
	mysql.Port = "3306"
	mysql.DataBaseName = ""
	mysql.UserName = ""
	mysql.PassWord = ""
}

func (mysql *MySql) InitmySqlConfig(path string) {
	mysql.DefaultmySqlConfig()
	file, _ := ioutil.ReadFile(path)
	if err := yaml.Unmarshal(file, mysql); err != nil {
		log.Println("ERROR", err)
	}
}
