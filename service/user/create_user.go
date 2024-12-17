package service

import (
	api "derma/detect/biz/model/api"
	"derma/detect/dal/db"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) CreateUser(req *api.UserRegisterRequest) (*db.User, error) {
	hashBytes, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	userModel := &db.User{
		Username: req.Username,
		Password: string(hashBytes),
		Birth:    time.Now().Format("2006.01.02"),
	}

	return db.CreateUser(s.ctx, userModel)
}
