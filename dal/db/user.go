package db

import (
	"context"
	"derma/detect/pkg/errno"
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        int64
	Username  string
	Password  string
	Avatar    string `gorm:"default:https://s2.loli.net/2024/01/02/bhezZY4OPAFDHd2.jpg"`
	Signature string `gorm:"default:NOT NULL BUT SEEMS NULL"`
	Email     string
	Phone     int64
	Birth     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func CreateUser(ctx context.Context, user *User) (*User, error) {
	userResp := new(User)

	user.Id = SF.NextVal()
	err := DB.WithContext(ctx).Where("username = ?", user.Username).First(&userResp).Error

	if err == nil {
		return nil, errno.UserExistedError
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err := DB.WithContext(ctx).Create(user).Error; err != nil {
		// add some logs
		return nil, err
	}

	return user, nil
}

func GetUserByUsername(ctx context.Context, username string) (*User, error) {
	userResp := new(User)

	err := DB.WithContext(ctx).Where("username = ?", username).First(&userResp).Error

	if err != nil {
		// add some logs

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("User not found")
		}
		return nil, err
	}

	return userResp, nil
}

func GetUserByID(ctx context.Context, userid int64) (*User, error) {
	userResp := new(User)

	err := DB.WithContext(ctx).Where("id = ?", userid).First(&userResp).Error

	if err != nil {
		// add some logs

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.UserNotFoundError
		}
		return nil, err
	}

	return userResp, nil
}

func ResetInfo(ctx context.Context, userModel *User) error {
	err := DB.Save(userModel).Error

	if err != nil {
		// add some logs

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errno.UserNotFoundError
		}
		return err
	}
	return nil
}

func ResetPassword(ctx context.Context, id int64, password string) error {
	err := DB.WithContext(ctx).Model(&User{}).Where("id= ?", id).
		Update("password", password).Error

	if err != nil {
		// add some logs

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errno.UserNotFoundError
		}
		return err
	}
	return nil
}
