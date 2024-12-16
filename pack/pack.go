package pack

import (
	"derma/detect/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Response struct {
	Code int64  `json:"status_code"`
	Msg  string `json:"status_msg"`
}

func SendFailResponse(c *app.RequestContext, err error) {
	if err == nil {
		c.JSON(consts.StatusOK, Response{
			Code: errno.SuccessCode,
			Msg:  errno.SuccessMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, Response{
		Code: -1,
		Msg:  errno.ConvertErr(err).Error(),
	})
}

func SendResponse(c *app.RequestContext, data interface{}) {
	c.JSON(consts.StatusOK, data)
}
