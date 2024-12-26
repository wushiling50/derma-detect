package errno

var (
	// Success
	Success = NewErrNo(SuccessCode, "Success")

	ServiceError             = NewErrNo(ServiceErrorCode, "service is unable to start successfully")
	ServiceInternalError     = NewErrNo(ServiceErrorCode, "service internal error")
	ParamError               = NewErrNo(ParamErrorCode, "parameter error")
	AuthorizationFailedError = NewErrNo(AuthorizationFailedErrCode, "authorization failed")

	// User
	UserExistedError  = NewErrNo(ParamErrorCode, "user existed")
	PasswordDiffError = NewErrNo(ParamErrorCode, "password different")
	PasswordSameError = NewErrNo(ParamErrorCode, "new password same as old password")
	EamilFormatError  = NewErrNo(ParamErrorCode, "invaild email address")
	PhoneFormatError  = NewErrNo(ParamErrorCode, "invaild phone number")
	AvatarUploadError = NewErrNo(FileErrorCode, "avatar upload error")

	// Picture
	PictureUploadError = NewErrNo(FileErrorCode, "image upload error")

	// article
	TitleEmptyError         = NewErrNo(ParamErrorCode, "title is empty")
	ContentTooFewWordsError = NewErrNo(ParamErrorCode, "content too few words")
	CoverEmptyError         = NewErrNo(ParamErrorCode, "cover is empty")

	// other
	UnexpectedTypeError     = NewErrNo(UnexpectedTypeErrorCode, "unexpected type")
	SensitiveWordsHTTPError = NewErrNo(ServiceErrorCode, "sensitive-words api error")
	UserNotFoundError       = NewErrNo(ParamErrorCode, "user not found")
	ArticleNotFoundError    = NewErrNo(ParamErrorCode, "article not found")
	NotPictureFileError     = NewErrNo(ParamErrorCode, "please upload picture")
	FileSizeError           = NewErrNo(ParamErrorCode, "file size too large")
)
