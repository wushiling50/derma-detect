package config

var (
	OSS      *oss
	MYSQL    *mysql
	REDIS    *redis
	RABBITMQ *rabbitMQ
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

type redis struct {
	Addr     string
	Password string
}

type rabbitMQ struct {
	Addr     string
	Username string
	Password string
}
