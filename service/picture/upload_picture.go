package picture

import (
	"bytes"
	"derma/detect/dal/db"
	"derma/detect/pkg/constants"
	"derma/detect/pkg/utils"
)

func (s *PictureService) UploadPicture(byteContainer []byte, userID int64) (*db.Picture, []*db.Detection, error) {
	pictureName := utils.GeneratePictureName(userID)
	fileReader := bytes.NewReader(byteContainer)

	objectKey := constants.PictureMainDir + "/" + pictureName

	err := s.bucket.PutObject(objectKey, fileReader)
	if err != nil {
		return nil, nil, err
	}

	url := utils.GetPictureURL(objectKey)
	picture := &db.Picture{
		User_id:     userID,
		Picture_url: url,
	}

	picture, err = db.CreatePicture(s.ctx, picture)
	if err != nil {
		return nil, nil, err
	}

	// TODO :get ai info

	detectionList := []*db.Detection{}
	detectionList, err = db.CreateDetection(s.ctx, detectionList)
	if err != nil {
		return nil, nil, err
	}

	return picture, detectionList, nil
}
