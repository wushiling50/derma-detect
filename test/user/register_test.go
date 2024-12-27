package main

import (
	"context"
	"derma/detect/biz/model/api"
	"derma/detect/dal/db"
	"testing"

	"bou.ke/monkey"
)

func testRegister(t *testing.T) {
	monkey.Patch(db.CreateUser, func(ctx context.Context, user *db.User) (*db.User, error) {
		return &db.User{Id: id}, nil
	})

	defer monkey.UnpatchAll()

	_, err := userService.CreateUser(&api.UserRegisterRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}
