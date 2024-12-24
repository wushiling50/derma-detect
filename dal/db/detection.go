package db

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Detection struct {
	Id            int64
	User_id       int64
	Picture_id    int64
	Detection_url string
	Percent       string
	Describe      string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func CreateDetection(ctx context.Context, detection []*Detection) ([]*Detection, error) {
	for _, v := range detection {
		v.Id = SF.NextVal()
	}

	err := DBDetection.WithContext(ctx).Create(&detection).Error
	if err != nil {
		return nil, err
	}

	return detection, nil
}

func GetDetectionById(ctx context.Context, pictureId int64, userId int64) ([]*Detection, error) {
	resp := make([]*Detection, 0)

	err := DBDetection.WithContext(ctx).Where("id = ? AND picture_id = ?", userId, pictureId).
		Order("created_at desc").Find(&resp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Picture not found")
		}
		return nil, err
	}
	return resp, nil
}
