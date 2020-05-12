package config

import (
	"github.com/spf13/viper"
	"os"
)


func Load() {
	dir, _ := os.Getwd()
	viper.SetConfigName("app")
	viper.SetConfigType("json")
	viper.AddConfigPath(dir+"/config")
	if err := viper.ReadInConfig(); err != nil {
		panic("配置错误，" + err.Error())
	}
	viper.SetDefault("db.driver", "mysql")
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "3306")
	viper.SetDefault("db.database", "vortex")
	viper.SetDefault("db.username", "root")
	viper.SetDefault("db.password", "")
	viper.SetDefault("db.charset", "utf8")
}