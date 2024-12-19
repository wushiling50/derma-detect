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

	// db table name
	UserTableName = "user"

	// oss
	AvatarMainDir  = "avatar"
	PictureMainDir = "picture"
)
