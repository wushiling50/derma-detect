package mq

import (
	"context"
	"derma/detect/pkg/constants"
	"derma/detect/pkg/utils"
	"sync"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn  *amqp.Connection
	mqurl string
}

type ArticleMQ struct {
	RabbitMQ
	channel   *amqp.Channel
	exchange  string
	queueName string
}

var (
	Rmq          *RabbitMQ
	ArticleMQCli *ArticleMQ
	Mu           sync.Mutex
)

func Init() {
	Rmq = &RabbitMQ{
		mqurl: utils.GetMQUrl(),
	}

	conn, err := amqp.Dial(Rmq.mqurl)

	if err != nil {
		hlog.Error(err)
		return
	}

	Rmq.conn = conn

	ArticleMQCli, err = ArticleMQInit()

	if err != nil {
		hlog.Error(err)
		return
	}

	ctx := context.Background()
	go run(ctx)
}

func ArticleMQInit() (*ArticleMQ, error) {
	ch, err := Rmq.conn.Channel()
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	return &ArticleMQ{
		RabbitMQ:  *Rmq,
		channel:   ch,
		queueName: constants.ArticleQueueName,
	}, nil
}

func run(ctx context.Context) {
	aSync := new(SyncArticle)
	err := aSync.SyncArticleMQ(ctx)
	if err != nil {
		hlog.Errorf("SyncArticleMQ:%s", err)
	}
}
