package article

import (
	"derma/detect/biz/model/api"
	"derma/detect/dal/db"
)

func (s *ArticleService) ArticleList(req api.ArticleListRequest) ([]*db.Article, error) {
	articles, err := db.GetArticles(s.ctx, req.PageNum)
	if err != nil {
		return nil, err
	}

	return articles, nil
}
