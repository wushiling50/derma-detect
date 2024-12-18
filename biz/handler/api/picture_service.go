// Code generated by hertz generator.

package api

import (
	"context"

	api "derma/detect/biz/model/api"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// UserHistory .
// @router /derma/detect/picture/history/ [GET]
func UserHistory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.HistoryRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.HistoryResponse)

	c.JSON(consts.StatusOK, resp)
}

// UploadPicture .
// @router /derma/detect/picture/upload/ [POST]
func UploadPicture(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UploadPictureRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.UploadPictureResponse)

	c.JSON(consts.StatusOK, resp)
}
