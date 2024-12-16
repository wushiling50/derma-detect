package errno

var (
	// Success
	Success = NewErrNo(SuccessCode, "Success")

	ServiceError             = NewErrNo(ServiceErrorCode, "service is unable to start successfully")
	ServiceInternalError     = NewErrNo(ServiceErrorCode, "service internal error")
	ParamError               = NewErrNo(ParamErrorCode, "parameter error")
	AuthorizationFailedError = NewErrNo(AuthorizationFailedErrCode, "authorization failed")

	// User
	UserExistedError = NewErrNo(ParamErrorCode, "user existed")

	// other
	UnexpectedTypeError     = NewErrNo(UnexpectedTypeErrorCode, "unexpected type")
	SensitiveWordsHTTPError = NewErrNo(ServiceErrorCode, "sensitive-words api error")
	UserNotFoundError       = NewErrNo(ParamErrorCode, "user not found")
)
