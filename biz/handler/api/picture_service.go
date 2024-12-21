// Code generated by hertz generator.

package api

import (
	"context"
	"io"

	api "derma/detect/biz/model/api"
	"derma/detect/pack"
	"derma/detect/pkg/errno"
	"derma/detect/pkg/utils"
	picture "derma/detect/service/picture"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/h2non/filetype"
)

// UserHistory .
// @router /derma/detect/picture/history/ [GET]
func UserHistory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.HistoryRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(api.HistoryResponse)

	// 对传入的数据做判断
	claim, err := utils.CheckToken(req.Token)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	// 发给业务层
	pictureResp, err := picture.NewPictureService(ctx).GetHistory(claim.UserId)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp.StatusCode, resp.StatusMsg = pack.BuildBaseResp(nil)
	// resp.PictureList = pack.Picture(userResp)
	for _, v := range pictureResp {
		resp.PictureList = append(resp.PictureList, pack.HistoryPicture(v))
	}

	pack.SendResponse(c, resp)
}

// UploadPicture .
// @router /derma/detect/picture/upload/ [POST]
func UploadPicture(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UploadPictureRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	// check avatar file
	file, err := c.FormFile("avatar")
	if err != nil {
		pack.SendFailResponse(c, errno.PictureUploadError)
		return
	}

	if !utils.IsPictureFile(file) {
		pack.SendFailResponse(c, errno.NotPictureFile)
		return
	}

	fileContent, err := file.Open()
	if err != nil {
		pack.SendFailResponse(c, errno.PictureUploadError)
		return
	}

	byteContainer, err := io.ReadAll(fileContent)
	if err != nil {
		pack.SendFailResponse(c, errno.PictureUploadError)
		return
	}

	if !filetype.IsImage(byteContainer) {
		pack.SendFailResponse(c, errno.NotPictureFile)
	}

	// TODO
	resp := new(api.UploadPictureResponse)
	// 对传入的数据做判断
	// claim, err := utils.CheckToken(req.Token)
	// if err != nil {
	// 	pack.SendFailResponse(c, err)
	// 	return
	// }

	pack.SendResponse(c, resp)
}
