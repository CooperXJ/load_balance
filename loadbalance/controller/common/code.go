package common

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota //下面会自增
	CodeNoServer
	CodeNotEnoughParams
	CodeInternalError
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeNoServer:        "没有合适的服务器",
	CodeNotEnoughParams: "缺少参数",
	CodeInternalError:   "内部错误",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeInternalError]
	}
	return msg
}
