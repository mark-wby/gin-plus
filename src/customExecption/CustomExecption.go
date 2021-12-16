package customExecption


//自定义异常类
type CustomExecption struct {
	Code int
	Msg string
}

func NewCustomExecption(code int, msg string) *CustomExecption {
	return &CustomExecption{Code: code, Msg: msg}
}


