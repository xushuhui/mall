package errcode

const (
	SUCCESS = iota
	InvalidParams
	NotFound
	ServerError
	AuthCheckTokenFail
	AuthCheckTokenTimeout
	ErrorAuthToken
	TimeoutAuthToken
	ErrorPassWord
)

var MsgFlags = map[int]string{
	SUCCESS:               "ok",
	InvalidParams:         "请求参数错误",
	NotFound:              "未找到记录",
	ServerError:           "系统异常，请联系管理员！",
	AuthCheckTokenFail:    "Token鉴权失败",
	AuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:        "Token错误",
	ErrorPassWord:         "密码错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[InvalidParams]
}
