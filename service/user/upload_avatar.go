package service

import (
	"bytes"
	api "derma/detect/biz/model/api"
	"derma/detect/dal/db"
	"derma/detect/pkg/constants"
	"derma/detect/pkg/utils"
)

func (s *UserService) UploadAvatar(req *api.UploadAvatarRequest, userID int64) (string, error) {
	avatarName := utils.GenerateAvatarName(userID)
	fileReader := bytes.NewReader(req.Avatar)
	url := "https://" + constants.AvatarMainDir + "/" + avatarName

	err := s.bucket.PutObject(constants.AvatarMainDir+"/"+avatarName, fileReader)
	if err != nil {
		return "", err
	}

	return url, db.RestAvatar(s.ctx, userID, url)
}
