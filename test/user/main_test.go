package main

import (
	"context"
	"derma/detect/config"
	"derma/detect/dal"
	"derma/detect/pkg/utils"
	"derma/detect/service/user"
	"testing"
)

var (
	username string
	password string
	token    string
	id       int64

	userService *user.UserService
)

func TestMain(m *testing.M) {
	config.InitForTest()
	dal.Init()

	userService = user.NewUserService(context.Background())

	username = "admin"
	password = "123456"
	id = 10001
	token, _ = utils.CreateToken(id)

	m.Run()
}

func TestMainOrder(t *testing.T) {
	t.Run("register", testRegister)

	t.Run("login", testLogin)

	t.Run("user_info", testGetInfo)

	t.Run("reset_password", testResetPassword)
}
