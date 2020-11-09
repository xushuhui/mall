package msg

import (
	"fmt"
	"mall_go/pkg/core"
)

type User struct {
	Name string
}

func StartMsg() {
	core.RegisterNotification(1, testNotify)
}
func testNotify(msgid int, data interface{}) {
	switch v := data.(type) {
	case User:
		fmt.Println("_________", v.Name)
		return
	}
}
