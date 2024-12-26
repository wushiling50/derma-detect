package pack

import (
	"derma/detect/biz/model/api"
	"derma/detect/dal/db"
	"fmt"
)

func Article(article *db.Article, author string) *api.Article {
	y, m, d := article.CreatedAt.Date()
	date := fmt.Sprintf("%v-%v-%v", y, int(m), d)
	return &api.Article{
		ID:          article.Id,
		Title:       article.Title,
		Content:     article.Content,
		Cover:       article.Cover_url,
		PublishTime: date,
		Author:      author,
	}
}

func ArticleList(articles []*db.Article) []*api.Article {
	resp := make([]*api.Article, 0)
	for _, article := range articles {
		resp = append(resp, &api.Article{
			ID:    article.Id,
			Title: article.Title,
			Cover: article.Cover_url,
		})
	}
	return resp
}
