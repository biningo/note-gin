package setup

import (
	"database/sql"
	"note-gin/global"
)

/**
*@Author lyer
*@Date 2/20/21 15:23
*@Describe
**/

import (
	"github.com/go-sql-driver/mysql"
)

func InitMySql() (*sql.DB, error) {
	connector, _ := mysql.NewConnector(&mysql.Config{
		Addr:      global.G_CONFIG.MySql.Addr + ":" + global.G_CONFIG.MySql.Port,
		User:      global.G_CONFIG.MySql.User,
		Passwd:    global.G_CONFIG.MySql.Password,
		DBName:    global.G_CONFIG.MySql.DB,
		Collation: global.G_CONFIG.MySql.Collation,
	})
	db := sql.OpenDB(connector)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}