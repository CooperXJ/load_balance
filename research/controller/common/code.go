package common

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota //下面会自增
	CodeCPUGetFailed
	CodeMemGetFailed
	CodeDiskGetFailed
	CodeNetIOGetFailed
	CodeSumFailed
	CodeOtherFailed
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:        "success",
	CodeCPUGetFailed:   "cpu信息获取失败",
	CodeMemGetFailed:   "内存信息获取失败",
	CodeDiskGetFailed:  "磁盘信息获取失败",
	CodeNetIOGetFailed: "网络IO信息获取失败",
	CodeSumFailed:      "系统报告获取失败",
	CodeOtherFailed:    "内部错误",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeOtherFailed]
	}
	return msg
}
