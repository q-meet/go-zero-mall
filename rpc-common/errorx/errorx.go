package errorx

var ParamsError = New(1101001, "参数不正确")

type BizError struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

type ErrorResponse struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

func New(code uint32, msg string) *BizError {
	return &BizError{
		Code: code,
		Msg:  msg,
	}
}

func (e *BizError) GetCode() uint32 {
	return e.Code
}

func (e *BizError) Error() string {
	return e.Msg
}

func (e *BizError) Data() interface{} {
	return &ErrorResponse{
		e.Code,
		e.Msg,
	}
}
