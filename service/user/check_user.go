package service

import (
	api "derma/detect/biz/model/api"
	"derma/detect/dal/db"
	"derma/detect/pkg/errno"

	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) CheckUser(req *api.UserLoginRequest) (*db.User, error) {
	userModel, err := db.GetUserByUsername(s.ctx, req.Username)

	if err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(req.Password)) != nil {
		return nil, errno.AuthorizationFailedError
	}

	return userModel, nil
}
