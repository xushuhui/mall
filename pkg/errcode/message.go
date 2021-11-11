package errcode

import "github.com/xushuhui/goal/core"

const (
	AuthCheckTokenFail = 10000
)

func init() {
	core.CodeMapping = map[int]string{
		AuthCheckTokenFail: "Token鉴权失败",
	}
}
