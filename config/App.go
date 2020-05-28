package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type App struct {
	PageSize       int    `yaml:"PageSize"`
	MakeMigration      bool   `yaml:"MakeMigration"` //是否进行数据库迁移
	QiniuAccessKey string `yaml:"QiniuAccessKey"`
	QiniuSecretKey string `yaml:"QiniuSecretKey"`
	LogFilePath    string `yaml:"LogFilePath"`
	JwtSecretKey    string `yaml:"JwtSecretKey"`
}

func (app *App) DefaultAppConfig() {
	app.PageSize = 13
	app.QiniuAccessKey = ""
	app.QiniuSecretKey = ""
	app.MakeMigration=false
	app.LogFilePath = "panic.log"
	app.JwtSecretKey = "note-gin"
}

func (app *App) InitAppConfig(path string) {
	app.DefaultAppConfig()
	file, _ := ioutil.ReadFile(path)
	if err := yaml.Unmarshal(file, app); err != nil {
		log.Println("ERROR", err)
	}
}
