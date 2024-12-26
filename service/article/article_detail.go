package article

import (
	"derma/detect/biz/model/api"
	"derma/detect/dal/db"
)

func (s *ArticleService) ArticleDetail(req api.ArticleDetailRequest, userID int64) (*db.Article, string, error) {
	article, err := db.GetArticleById(s.ctx, req.ArticleID, userID)
	if err != nil {
		return nil, "", err
	}

	user, err := db.GetUserByID(s.ctx, userID)
	if err != nil {
		return nil, "", err
	}

	return article, user.Username, nil
}
