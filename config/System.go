package config

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)
type System struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	RunMode string `yaml:"Mode"`
	PageSize int`yaml:"PageSize"`

}

func (system *System)DefaultSystemConfig(){
	system.Host = "localhost"
	system.Port = "8080"
}

func (system *System) InitSystemConfig(path string)  {
	system.DefaultSystemConfig()
	file,_:=ioutil.ReadFile(path)
	if err:=yaml.Unmarshal(file,system);err!=nil{
		log.Println("ERROR",err)
	}
}

