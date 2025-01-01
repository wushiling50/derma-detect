package mq

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (a *ArticleMQ) Publish(ctx context.Context, body string) error {
	_, err := a.channel.QueueDeclare(
		a.queueName,
		true, // 持久化
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		hlog.Error(err)
		return err
	}

	err = a.channel.PublishWithContext(ctx,
		a.exchange,
		a.queueName,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, // 持久化模式
			ContentType:  "application/json",
			Body:         []byte(body),
		},
	)
	if err != nil {
		hlog.Error(err)
		return err
	}

	hlog.Info("send msg success")
	return nil
}

func (a *ArticleMQ) Consume(ctx context.Context) (<-chan amqp.Delivery, error) {
	_, err := a.channel.QueueDeclare(
		a.queueName,
		true, // 持久化
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	msgs, err := a.channel.Consume(
		a.queueName,
		"",
		false, // 关闭自动应答
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	hlog.Info("msg consumer setup")
	return msgs, nil
}

// 释放资源,建议获取实例后配合defer使用
func (a *ArticleMQ) ReleaseRes() {
	a.conn.Close()
	a.channel.Close()
}
