package e

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004

	ERROR_EXIST_USERNAME     = 30001
	ERROR_DIFF_PASSWORD      = 30002
	ERROR_NOT_EXIST_USERNAME = 30003
	ERROR_INCORRECT_PASSWORD = 30004
	ERROR_EDIT_PASSWORD      = 30005
	ERROR_NOT_EXIST_USERID   = 30006
	// 继续定义自己的错误码
)
