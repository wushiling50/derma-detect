package main

import (
	"context"
	"derma/detect/biz/model/api"
	"derma/detect/dal/db"
	"testing"

	"bou.ke/monkey"
)

func testResetPassword(t *testing.T) {
	monkey.Patch(db.ResetPassword, func(ctx context.Context, id int64, password string) error {
		return nil
	})

	defer monkey.UnpatchAll()

	err := userService.ResetPassword(&api.ResetPasswordRequest{
		Token:         token,
		OldPassword:   password,
		NewPassword:   "654321",
		CheckPassword: "654321",
	}, id)

	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}
