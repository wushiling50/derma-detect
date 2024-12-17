package pack

import (
	api "derma/detect/biz/model/api"

	"derma/detect/dal/db"
)

func User(user *db.User) *api.User {
	return &api.User{
		ID:        user.Id,
		Username:  user.Username,
		Avatar:    user.Avatar,
		Signature: user.Signature,
		Email:     user.Email,
		Phone:     user.Phone,
		Birth:     user.Birth,
	}
}
