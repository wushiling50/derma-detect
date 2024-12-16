package pack

import (
	"derma/detect/pkg/errno"
	"errors"
)

func BuildBaseResp(err error) (int64, *string) {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}

	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceError.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) (int64, *string) {
	return err.ErrorCode, &err.ErrorMsg
}
