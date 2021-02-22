package global

/**
*@Author lyer
*@Date 2/20/21 15:13
*@Describe
**/

import (
	"database/sql"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"note-gin/config"
)

var (
	G_VP     *viper.Viper
	G_CONFIG config.Config
	G_DB     *sql.DB
	G_Redis *redis.Client
)
