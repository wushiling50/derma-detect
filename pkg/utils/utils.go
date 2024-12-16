package utils

import (
	"strings"

	"github.com/spf13/viper"
)

func GetMysqlDSN() string {
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	dbName := viper.GetString("mysql.dbName")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	charset := viper.GetString("mysql.charset")

	dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", dbName, "?charset=" + charset + "&parseTime=true"}, "")

	return dsn
}
