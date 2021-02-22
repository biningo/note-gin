package setup

import (
	"flag"
	"github.com/spf13/viper"
	"note-gin/global"
)

/**
*@Author lyer
*@Date 2/20/21 15:21
*@Describe
**/

func InitViper() (*viper.Viper, error) {
	path := ""
	flag.StringVar(&path, "c", "config.yaml", "choose config file")
	flag.Parse()
	v := viper.New()
	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := v.Unmarshal(&global.G_CONFIG); err != nil {
		return nil, err
	}
	return v, nil
}
