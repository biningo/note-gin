package test

import (
	"io/ioutil"
	"note-gin/config"
	"testing"
)

func TestConfigPath(t *testing.T) {
	cfp := config.NewconfigFilePath("../config/file/BootLoader.yaml") //路径是从当前目录出发  不是从NewconfigFilePath的目录出发
	t.Log(cfp.AppPath)
	cfg := config.Config{}

	cfg.AppConfig.InitAppConfig(cfp.AppPath)
	t.Log(cfg.AppConfig)
}

func TestIoutil(t *testing.T) {
	b, e := ioutil.ReadFile("../config/file/BootLoader.yaml")
	t.Log(string(b), e)
}

func TestApp(t *testing.T) {
	var cfg config.Config //所有属性都会初始化
	t.Log(cfg.AppConfig)
}
