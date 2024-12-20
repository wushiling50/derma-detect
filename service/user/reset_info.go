package user

import (
	"context"
	api "derma/detect/biz/model/api"

	"derma/detect/dal/db"
)

func (s *UserService) ResetInfo(req *api.ResetInfoRequest, userID int64) error {
	userModel, err := compareInfo(s.ctx, userID, req)
	if err != nil {
		return err
	}

	return db.ResetInfo(s.ctx, userID, userModel)
}

func compareInfo(ctx context.Context, id int64, req *api.ResetInfoRequest) (*db.User, error) {
	userModel := &db.User{}

	oldInfo, err := db.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Username != nil && req.Username != &oldInfo.Username {
		userModel.Username = *req.Username
	}

	if req.Email != nil && req.Email != &oldInfo.Email {
		userModel.Email = *req.Email
	}

	if req.Phone != nil && req.Phone != &oldInfo.Phone {
		userModel.Phone = *req.Phone
	}

	if req.Birth != nil && req.Birth != &oldInfo.Birth {
		userModel.Birth = *req.Birth
	}

	return userModel, nil
}
