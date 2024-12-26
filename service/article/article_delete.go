package article

import (
	"derma/detect/biz/model/api"
	"derma/detect/dal/db"
)

func (s *ArticleService) ArticleDelete(req api.ArticleDeleteRequest, userID int64) error {
	article := &db.Article{
		Id: req.ArticleID,
	}

	err := db.ArticleDelete(s.ctx, article, userID)
	if err != nil {
		return err
	}

	return nil
}
