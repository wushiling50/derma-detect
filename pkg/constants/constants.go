package constants

import "time"

const (
	// auth
	JWTValue = "MT8xNTkwM0g1Mw=="
	StartID  = 10000

	// snowflake
	SnowflakeWorkerID     = 0
	SnowflakeDatacenterID = 0

	// limit
	MaxIdleConns    = 10
	MaxConnections  = 1000
	ConnMaxLifetime = 10 * time.Second
	MaxFileSize     = 10 << 20 // 10 MB
	MinContent      = 50

	// db table name
	UserTableName      = "user"
	PictureTableName   = "picture"
	DetectionTableName = "detection"
	ArticleTableName   = "article"

	// oss
	AvatarMainDir    = "avatar"
	PictureMainDir   = "picture"
	DetectionMainDir = "detection"
	ArticleMainDir   = "article"

	// redis
	RedisDBArticle = 1
)
