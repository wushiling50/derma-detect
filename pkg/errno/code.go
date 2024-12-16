package errno

const (
	// For api-gateway
	SuccessCode = 0
	SuccessMsg  = "ok"

	// Error
	ServiceErrorCode           = 10001 // 未知服务错误
	ParamErrorCode             = 10002 // 参数错误
	AuthorizationFailedErrCode = 10003 // 鉴权失败
	UnexpectedTypeErrorCode    = 10004 // 未知类型
)
