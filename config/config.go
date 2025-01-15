package config

import (
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	Init()
}

func Init() {
	mysqlInit()
	ossInit()
	redisInit()
	rabbitMQInit()
	aiServerInit()
}

func ossInit() {
	OSS = &oss{}
	OSS.Endpoint = viper.GetString("oss.endpoint")
	OSS.AccessKeyID = viper.GetString("oss.accessKey-id")
	OSS.AccessKeySecret = viper.GetString("oss.accessKey-secret")
	OSS.BucketName = viper.GetString("oss.bucketname")
}

func mysqlInit() {
	MYSQL = &mysql{}
	MYSQL.Host = viper.GetString("mysql.host")
	MYSQL.Port = viper.GetString("mysql.port")
	MYSQL.DbName = viper.GetString("mysql.dbName")
	MYSQL.Username = viper.GetString("mysql.username")
	MYSQL.Password = viper.GetString("mysql.password")
	MYSQL.Charset = viper.GetString("mysql.charset")
}

func redisInit() {
	REDIS = &redis{}
	REDIS.Addr = viper.GetString("redis.addr")
	REDIS.Password = viper.GetString("redis.password")
}

func rabbitMQInit() {
	RABBITMQ = &rabbitMQ{}
	RABBITMQ.Addr = viper.GetString("rabbitmq.addr")
	RABBITMQ.Username = viper.GetString("rabbitmq.username")
	RABBITMQ.Password = viper.GetString("rabbitmq.password")
}

func aiServerInit() {
	AISERVER = &aiserver{}
	AISERVER.Addr = viper.GetString("aiserver.addr")
}

func InitForTest() {
	MYSQL = &mysql{
		DriverName: "mysql",
		Host:       "127.0.0.1",
		Port:       "3306",
		DbName:     "derma",
		Username:   "derma",
		Password:   "derma",
		Charset:    "utf8mb4",
	}

	RABBITMQ = &rabbitMQ{
		Addr:     "127.0.0.1:5672",
		Username: "tiktok",
		Password: "tiktok",
	}

	REDIS = &redis{
		Addr:     "127.0.0.1:6379",
		Password: "tiktok",
	}
}
