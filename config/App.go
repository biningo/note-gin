package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type App struct {
	PageSize       int    `yaml:"PageSize"`
	Migration      bool   `yaml:"Migration"`
	QiniuAccessKey string `yaml:"QiniuAccessKey"`
	QiniuSecretKey string `yaml:"QiniuSecretKey"`
}

func (app *App) DefaultAppConfig() {
	app.PageSize = 0
	app.QiniuAccessKey = ""
	app.QiniuSecretKey = ""
}

func (app *App) InitAppConfig(path string) {
	app.DefaultAppConfig()
	file, _ := ioutil.ReadFile(path)
	if err := yaml.Unmarshal(file, app); err != nil {
		log.Println("ERROR", err)
	}
}
