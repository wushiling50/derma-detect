package mq

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type SyncArticle struct {
}

func (s *SyncArticle) SyncArticleMQ(ctx context.Context) error {
	defer ArticleMQCli.ReleaseRes()

	msgs, err := ArticleMQCli.Consume(ctx)
	if err != nil {
		hlog.Error(err)
		return err
	}

	var forever chan struct{}

	go func() {
		for msg := range msgs {
			hlog.Infof("Resolve msg: %s", msg.Body)

			// TODO: finish insert logic in mq
			// err := ArticleMQCli.ArticleInsert(ctx, msg)
			// if err != nil {
			// 	hlog.Errorf("Insert follow action: %s", err)
			// 	continue
			// }

			err = msg.Ack(false)
			if err != nil {
				hlog.Error(err)
				continue
			}
		}
	}()

	<-forever

	return nil
}
