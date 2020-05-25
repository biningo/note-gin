package test

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"note-gin/config"
	"testing"

)
var mySqlConfig = config.Conf.MySqlConfig
func TestOpenSql(t *testing.T){
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		"root", "55555", "118.178.180.115","6379", "note")

	_, err := gorm.Open("mysql", connStr)
	t.Log(err)
	t.Log(connStr)
}
