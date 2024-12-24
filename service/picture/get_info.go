package picture

import (
	"derma/detect/dal/db"
)

func (s *PictureService) GetInfo(userID int64, pictureID int64) (*db.Picture, []*db.Detection, error) {
	picture, err := db.GetPictureByPictureId(s.ctx, pictureID)
	if err != nil {
		return nil, nil, err
	}

	detections, err := db.GetDetectionById(s.ctx, pictureID, userID)
	if err != nil {
		return nil, nil, err
	}

	return picture, detections, nil
}
