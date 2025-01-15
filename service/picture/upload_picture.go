package picture

import (
	"bytes"
	"derma/detect/dal/db"
	aireq "derma/detect/pkg/ai_req"
	"derma/detect/pkg/constants"
	"derma/detect/pkg/utils"
	"os"
	"strconv"

	"github.com/cloudwego/hertz/pkg/common/hlog"
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

	// get ai info
	filePath := "../../temp/" + pictureName
	err = os.WriteFile(filePath, byteContainer, 0644)
	if err != nil {
		return nil, nil, err
	}

	defer func() {
		err := os.RemoveAll(filePath)
		if err != nil {
			hlog.Infof("recover panic : %v", err)
			return
		}
	}()

	data, err := aireq.PictureReq(filePath)
	if err != nil {
		return nil, nil, err
	}

	detectionList := []*db.Detection{}
	for _, v := range data.Predictions {
		d := db.Detection{}
		d.Name = v.Class
		d.Percent = strconv.FormatFloat(v.Probability, 'f', 2, 64)
		detectionList = append(detectionList, &d)
	}

	detectionList, err = db.CreateDetection(s.ctx, detectionList)
	if err != nil {
		return nil, nil, err
	}

	return picture, detectionList, nil
}
