package e

// CustomError 业务自定义错误
type CustomError struct {
	Code int
	Msg  string
}

func (e *CustomError) Error() string {
	return e.Msg
}

func New(code int, msg string) *CustomError {
	return &CustomError{Code: code, Msg: msg}
}

// 常用业务错误码
var (
	ErrInvalidParam = New(400, "请求参数错误！")
	ErrNotFound     = New(204, "资源不存在！")
	ErrInternal     = New(500, "服务器内部错误！")
)
