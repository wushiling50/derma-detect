package user

import (
	"bytes"
	"derma/detect/dal/db"
	"derma/detect/pkg/constants"
	"derma/detect/pkg/utils"
)

func (s *UserService) UploadAvatar(byteContainer []byte, userID int64) (string, error) {
	avatarName := utils.GenerateAvatarName(userID)
	fileReader := bytes.NewReader(byteContainer)

	objectKey := constants.AvatarMainDir + "/" + avatarName

	err := s.bucket.PutObject(objectKey, fileReader)
	if err != nil {
		return "", err
	}

	url := utils.GetAvatarURL(objectKey)
	return url, db.RestAvatar(s.ctx, userID, url)
}
