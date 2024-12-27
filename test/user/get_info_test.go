package main

import (
	"derma/detect/biz/model/api"
	"testing"
)

func testGetInfo(t *testing.T) {
	_, err := userService.GetInfo(&api.UserInfoRequest{
		Token: token,
	}, id)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
