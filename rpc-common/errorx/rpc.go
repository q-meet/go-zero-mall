package errorx

import "fmt"

// 成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码
const SERVER_COMMON_ERROR uint32 = 1000114
const REUQEST_PARAM_ERROR uint32 = 100002
const TOKEN_EXPIRE_ERROR uint32 = 100003
const TOKEN_GENERATE_ERROR uint32 = 100004
const DB_ERROR uint32 = 100005

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
	message[REUQEST_PARAM_ERROR] = "参数错误"
	message[TOKEN_EXPIRE_ERROR] = "token失效，请重新登陆"
	message[TOKEN_GENERATE_ERROR] = "生成token失败"
	message[DB_ERROR] = "数据库繁忙,请稍后再试"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}

/**
常用通用固定错误
*/

type CodeError struct {
	errCode uint32
	errMsg  string
}

// 返回给前端的错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.errCode
}

// GetErrMsg 返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errMsg}
}
func NewErrCode(errCode uint32) *CodeError {
	return &CodeError{errCode: errCode, errMsg: MapErrMsg(errCode)}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{errCode: SERVER_COMMON_ERROR, errMsg: errMsg}
}

var ErrUserAlreadyRegisterError = NewErrCodeMsg(1000021, "该用户已被注册")
