package config

import "github.com/spf13/viper"

var (
	OSS   *oss
	MYSQL *mysql
)

type mysql struct {
	DriverName string
	Host       string
	Port       string
	DbName     string
	Username   string
	Password   string
	Charset    string
}

type oss struct {
	Endpoint        string
	AccessKeyID     string `mapstructure:"accessKey-id"`
	AccessKeySecret string `mapstructure:"accessKey-secret"`
	BucketName      string
}

func OssInit() {
	OSS = &oss{}
	OSS.Endpoint = viper.GetString("oss.endpoint")
	OSS.AccessKeyID = viper.GetString("oss.accessKey-id")
	OSS.AccessKeySecret = viper.GetString("oss.accessKey-secret")
	OSS.BucketName = viper.GetString("oss.bucketname")
}

func MysqlInit() {
	MYSQL = &mysql{}
	MYSQL.Host = viper.GetString("mysql.host")
	MYSQL.Port = viper.GetString("mysql.port")
	MYSQL.DbName = viper.GetString("mysql.dbName")
	MYSQL.Username = viper.GetString("mysql.username")
	MYSQL.Password = viper.GetString("mysql.password")
	MYSQL.Charset = viper.GetString("mysql.charset")
}
