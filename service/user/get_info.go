package user

import (
	api "derma/detect/biz/model/api"

	"derma/detect/dal/db"
)

func (s *UserService) GetInfo(req *api.UserInfoRequest, userID int64) (*db.User, error) {
	user, err := db.GetUserByID(s.ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
