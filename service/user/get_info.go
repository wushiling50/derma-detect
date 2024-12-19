package service

import (
	api "derma/detect/biz/model/api"

	"derma/detect/dal/db"
)

func (s *UserService) GetInfo(req *api.UserInfoRequest, userID int64) (*db.User, error) {
	return db.GetUserByID(s.ctx, userID)
}
