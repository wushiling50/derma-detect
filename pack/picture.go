package pack

import (
	"derma/detect/biz/model/api"
	"derma/detect/dal/db"
)

func HistoryPicture(picture *db.Picture) *api.Picture {
	return &api.Picture{
		ID:         picture.Id,
		PictureURL: picture.Picture_url,
		CreateDate: picture.Create_date,
	}
}
