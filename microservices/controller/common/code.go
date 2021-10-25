package common

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota //下面会自增
	CodeMemError
	CodeCPUError
	CodeInternalError
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:       "success",
	CodeCPUError:      "CPU类型服务出错",
	CodeMemError:      "内存类型服务出错",
	CodeInternalError: "内部错误",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeInternalError]
	}
	return msg
}
