package dal

import (
	"derma/detect/dal/cache"
	"derma/detect/dal/db"
)

func Init() {
	db.InitDB()
	cache.Init()
}
