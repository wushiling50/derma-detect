package db

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Picture struct {
	Id          int64
	User_id     int64
	Picture_url string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func CreatePicture(ctx context.Context, picture *Picture) (*Picture, error) {
	picture.Id = SF.NextVal()
	err := DBPicture.WithContext(ctx).Create(picture).Error
	if err != nil {
		// add some logs
		return nil, err
	}

	return picture, nil
}

func GetPictureByUserId(ctx context.Context, userId int64) ([]*Picture, error) {
	resp := make([]*Picture, 0)

	err := DBPicture.WithContext(ctx).Where("user_id = ?", userId).
		Order("created_at desc").Find(&resp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Picture not found")
		}
		return nil, err
	}

	return resp, nil
}

func GetPictureByPictureId(ctx context.Context, pictureId int64) (*Picture, error) {
	resp := new(Picture)

	err := DBPicture.WithContext(ctx).Where("id = ?", pictureId).
		Order("created_at desc").Find(&resp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Picture not found")
		}
		return nil, err
	}

	return resp, nil
}
