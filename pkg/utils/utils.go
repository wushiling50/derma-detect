package utils

import (
	"derma/detect/pkg/errno"
	"regexp"
	"strconv"
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

func CheckPhoneNum(phone int64) error {
	s := strconv.FormatInt(phone, 10)
	re, err := regexp.Compile(`^(13[0-9]|14[5|7]|15[0|1|2|3|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\d{8}$`)
	if err != nil {
		return err
	}

	matched := re.MatchString(s)
	if !matched {
		return errno.EamilFormatError
	}

	return nil
}

func CheckEmail(email string) error {
	re, err := regexp.Compile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)
	if err != nil {
		return err
	}

	matched := re.MatchString(email)
	if !matched {
		return errno.EamilFormatError
	}

	return nil
}
