package config

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Server struct {
	Host    string `yaml:"Host"`
	Port    string `yaml:"Port"`
	RunMode string `yaml:"RunMode"`
}

func (server *Server) DefaultServerConfig() {
	server.Host = "localhost"
	server.Port = "9000"
	server.RunMode = "release"
}

func (server *Server) InitServerConfig(path string) {
	server.DefaultServerConfig()
	file, _ := ioutil.ReadFile(path)
	if err := yaml.Unmarshal(file, server); err != nil {
		log.Println("ERROR", err)
	}
}
