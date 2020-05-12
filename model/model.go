package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)
var DB *gorm.DB


func Load() {
	driver := viper.GetString("db.driver")
	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	database := viper.GetString("db.database")
	username := viper.GetString("db.username")
	password := viper.GetString("db.password")
	charset := viper.GetString("db.charset")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
			username,
			password,
			host,
			port,
			database,
			charset)
	var err error
	fmt.Println(dsn)
	DB, err = gorm.Open(driver, dsn)
	if err != nil {
		panic("数据库连接失败！" + err.Error())
	}
	DB.LogMode(true)
}