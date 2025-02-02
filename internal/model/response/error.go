package response

type ErrorCode struct {
	Code int
	Msg  string
}

func (ec ErrorCode) Error() string {
	return ec.Msg
}

func newErrorCode(code int, msg string) ErrorCode {
	return ErrorCode{
		Code: code,
		Msg:  msg,
	}
}

var (
	NotLoginOrInvalidAccess = newErrorCode(7001, "未登录或非法访问")
	AuthExpired             = newErrorCode(7002, "授权已过期")
	UsernameOrPWDError      = newErrorCode(7003, "用户名或密码错误")
)
