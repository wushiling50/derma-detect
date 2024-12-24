package pack

import (
	"derma/detect/biz/model/api"
	"derma/detect/dal/db"
	"fmt"
)

func Picture(picture *db.Picture) *api.Picture {
	y, m, d := picture.CreatedAt.Date()
	date := fmt.Sprintf("%v-%v-%v", y, m, d)
	return &api.Picture{
		ID:         picture.Id,
		PictureURL: picture.Picture_url,
		CreateDate: date,
	}
}
