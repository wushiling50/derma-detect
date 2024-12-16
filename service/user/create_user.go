package service

import (
	api "derma/detect/biz/model/api"
	"derma/detect/dal/db"

	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) CreateUser(req *api.UserRegisterRequest) (*db.User, error) {
	hashBytes, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	userModel := &db.User{
		Username: req.Username,
		Password: string(hashBytes),
	}

	return db.CreateUser(s.ctx, userModel)
}
