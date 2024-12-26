package db

import (
	"context"
	"derma/detect/pkg/errno"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Article struct {
	Id        int64
	User_id   int64
	Title     string
	Content   string
	Cover_url string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func Publish(ctx context.Context, article *Article) (*Article, error) {
	article.Id = SF.NextVal()

	if err := DBArticle.WithContext(ctx).Create(article).Error; err != nil {
		// add some logs
		return nil, err
	}

	return article, nil
}

func GetArticleById(ctx context.Context, articleID int64, userID int64) (*Article, error) {
	articleResp := new(Article)

	err := DBArticle.WithContext(ctx).Where("id = ? AND user_id = ?", articleID, userID).First(&articleResp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ArticleNotFoundError
		}
		return nil, err
	}

	return articleResp, nil
}

func ArticleDelete(ctx context.Context, article *Article, userID int64) error {
	err := DBArticle.WithContext(ctx).Where("user_id = ?", userID).Delete(article).Error
	if err != nil {
		return err
	}

	return nil
}

func GetArticles(ctx context.Context, page_num int64) ([]*Article, error) {
	article_list := make([]*Article, 0)

	err := DBArticle.WithContext(ctx).Limit(20).Offset(int((page_num - 1) * 20)).
		Find(&article_list).Error
	if err != nil {
		return nil, err
	}

	return article_list, nil
}
