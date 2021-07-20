package base

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func NewError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func NewResult(code int, data interface{}) *Result {
	return &Result{Code: code, Data: data}
}

func (e *CodeError) Error() string {
	return e.Msg
}
