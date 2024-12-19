package service

import (
	api "derma/detect/biz/model/api"
	"derma/detect/dal/db"
)

func (s *UserService) ResetPassword(req *api.ResetPasswordRequest, userID int64) error {
	return db.ResetPassword(s.ctx, userID, req.NewPassword)
}
