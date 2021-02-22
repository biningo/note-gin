package config

type App struct {
	PageSize       int    `yaml:"page-size"`
	QiniuAccessKey string `yaml:"qiniu-access-key"`
	QiniuSecretKey string `yaml:"qiniu-secret-key"`
}
