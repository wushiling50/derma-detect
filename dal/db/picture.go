package db

import (
	"time"

	"gorm.io/gorm"
)

type Picture struct {
	Id          int64
	User_id     int64
	Picture_url string
	Create_date string
	Percent     string
	Describe    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
