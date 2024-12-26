package article

import (
	"derma/detect/biz/model/api"
	"derma/detect/dal/db"
)

func (s *ArticleService) Publish(req api.PublishRequest, userID int64) (*db.Article, error) {
	article := &db.Article{
		User_id:   userID,
		Title:     req.Title,
		Content:   req.Content,
		Cover_url: req.Cover,
	}

	return db.Publish(s.ctx, article)
}
