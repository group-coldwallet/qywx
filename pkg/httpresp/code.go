package httpresp

type HttpCode int

//常规错误：10000 - 19999
//HTTP API接口错误：20000 - 29999
// 定义错误信息
const (
	FAIL    = 10000
	SUCCESS = 200
)

var MsgFlags = map[int]string{
	SUCCESS: "ok",
	FAIL:    "fail",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[FAIL]
}
