package main

import (
	"derma/detect/biz/model/api"
	"derma/detect/pkg/utils"
	"testing"
)

func testLogin(t *testing.T) {
	resp, err := userService.CheckUser(&api.UserLoginRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	token, err = utils.CreateToken(resp.Id)

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	id = resp.Id
}
