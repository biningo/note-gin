package QiniuClient

import "note-gin/config"

var QiniuAccessKey string
var QiniuSecretKey string

func SetUp() {
	QiniuAccessKey = config.Conf.AppConfig.QiniuAccessKey
	QiniuSecretKey = config.Conf.AppConfig.QiniuSecretKey
}
