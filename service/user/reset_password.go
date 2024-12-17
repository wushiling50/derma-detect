package service

import (
	api "derma/detect/biz/model/api"
	"derma/detect/dal/db"
	"derma/detect/pkg/utils"
)

func (s *UserService) ResetPassword(req *api.ResetPasswordRequest) error {
	claim, err := utils.CheckToken(req.Token)
	if err != nil {
		return err
	}

	return db.ResetPassword(s.ctx, claim.UserId, req.NewPassword)
}
