package exception

// MsgFlags this is MsgFlags
var MsgFlags = map[int]string{
	SUCCESS:                    "ok",
	ERROR:                      "服务异常",
	ServerError:                "服务器错误",
	InvalidParams:              "请求参数错误",
	ErrotAuthCheckTokenFail:    "Token鉴权失败",
	ErrotAuthCheckTokenTimeout: "Token已超时",
	ErrotAuthToken:             "Token生成失败",
	ErrotAuth:                  "Token错误",
	ErrorSign:                  "验证签名失败",
}

// GetMsg this is GetMsg
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
