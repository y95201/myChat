package models

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ChatDB *gorm.DB

func InitDB() {
	//初始化
	dsn := viper.GetString(`mysql.dsn`)
	ChatDB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}
