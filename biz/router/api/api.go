// Code generated by hertz generator. DO NOT EDIT.

package api

import (
	api "derma/detect/biz/handler/api"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_derma := root.Group("/derma", _dermaMw()...)
		{
			_detect := _derma.Group("/detect", _detectMw()...)
			_detect.GET("/ai-doctor", append(_aidoctorMw(), api.AIDoctor)...)
			{
				_picture := _detect.Group("/picture", _pictureMw()...)
				{
					_history := _picture.Group("/history", _historyMw()...)
					_history.GET("/", append(_userhistoryMw(), api.UserHistory)...)
					{
						_info := _history.Group("/info", _infoMw()...)
						_info.GET("/", append(_historyinfoMw(), api.HistoryInfo)...)
					}
				}
				{
					_upload := _picture.Group("/upload", _uploadMw()...)
					_upload.POST("/", append(_uploadpictureMw(), api.UploadPicture)...)
				}
			}
			{
				_user := _detect.Group("/user", _userMw()...)
				{
					_info0 := _user.Group("/info", _info0Mw()...)
					_info0.GET("/", append(_userinfoMw(), api.UserInfo)...)
				}
				{
					_login := _user.Group("/login", _loginMw()...)
					_login.POST("/", append(_userloginMw(), api.UserLogin)...)
				}
				{
					_register := _user.Group("/register", _registerMw()...)
					_register.POST("/", append(_userregisterMw(), api.UserRegister)...)
				}
				{
					_reset_info := _user.Group("/reset-info", _reset_infoMw()...)
					_reset_info.POST("/", append(_resetinfoMw(), api.ResetInfo)...)
				}
				{
					_reset_password := _user.Group("/reset-password", _reset_passwordMw()...)
					_reset_password.POST("/", append(_resetpasswordMw(), api.ResetPassword)...)
				}
				{
					_upload_avatar := _user.Group("/upload-avatar", _upload_avatarMw()...)
					_upload_avatar.POST("/", append(_uploadavatarMw(), api.UploadAvatar)...)
				}
			}
		}
	}
}
