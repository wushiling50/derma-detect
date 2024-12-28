package cache

import (
	"derma/detect/pkg/constants"

	"github.com/redis/go-redis/v9"
)

var (
	RedisArticleClient *redis.Client
)

func Init() {
	RedisArticleClient = RCBuilder(constants.RedisDBArticle)
}
