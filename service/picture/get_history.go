package picture

import "derma/detect/dal/db"

func (s *PictureService) GetHistory(userID int64) ([]*db.Picture, error) {
	return db.GetPictureByUserId(s.ctx, userID)
}
