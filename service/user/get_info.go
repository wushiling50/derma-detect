package service

import (
	api "derma/detect/biz/model/api"
	"derma/detect/pkg/utils"

	"derma/detect/dal/db"
)

func (s *UserService) GetInfo(req *api.UserInfoRequest) (*db.User, error) {
	claim, err := utils.CheckToken(req.Token)
	if err != nil {
		return nil, err
	}

	return db.GetUserByID(s.ctx, claim.UserId)
}
