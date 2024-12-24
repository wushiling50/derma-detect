package pack

import (
	"derma/detect/biz/model/api"
	"derma/detect/dal/db"
)

func Detection(detection *db.Detection) *api.Detection {
	return &api.Detection{
		PictureURL: detection.Detection_url,
		Percent:    detection.Percent,
		Describe:   detection.Describe,
	}
}
