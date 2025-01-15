package aireq

import (
	"github.com/cloudwego/hertz/pkg/app/client"
)

var AIClinet *client.Client

func Init() {
	c, err := client.NewClient()
	if err != nil {
		panic(err)
	}
	AIClinet = c
}
