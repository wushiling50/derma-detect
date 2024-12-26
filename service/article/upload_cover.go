package article

import (
	"bytes"
	"derma/detect/pkg/constants"
	"derma/detect/pkg/utils"
)

func (s *ArticleService) UploadCover(byteContainer []byte, userID int64) (string, error) {
	coverName := utils.GenerateCoverName(userID)
	fileReader := bytes.NewReader(byteContainer)

	objectKey := constants.ArticleMainDir + "/" + coverName

	err := s.bucket.PutObject(objectKey, fileReader)
	if err != nil {
		return "", err
	}

	return utils.GetCoverURL(objectKey), nil
}
