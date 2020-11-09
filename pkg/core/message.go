package core

import (
	"fmt"
	"time"
)

//消息中心，单例
//支持两种类型的回调：
//	1.异步通知：异步调用，没有返回值
//	2.同步回调：同步调用，有返回值
type MsgData struct {
	msgID int
	data  interface{}
}

type Notification func(int, interface{})
type Callback func(int, interface{}) interface{}

type MessageCenter struct {
	callbacks     map[int]map[string]Callback
	notifications map[int]map[string]Notification
	ch            chan MsgData
}

var msgCenter MessageCenter

func initMsg() {
	msgCenter.callbacks = make(map[int]map[string]Callback)
	msgCenter.notifications = make(map[int]map[string]Notification)
	msgCenter.ch = make(chan MsgData, 50000)
	go dispatch()
}

func RegisterCallback(msgID int, name string, callback Callback) (result interface{}) {
	callbacks, ok := msgCenter.callbacks[msgID]
	if !ok {
		callbacks = make(map[string]Callback)
		msgCenter.callbacks[msgID] = callbacks
	}
	callbacks[name] = callback
	return
}

func RegisterNotification(msgID int, n Notification) {
	notifications, ok := msgCenter.notifications[msgID]
	if !ok {
		notifications = make(map[string]Notification)
		msgCenter.notifications[msgID] = notifications
	}
	notifications[fmt.Sprintf("%v", n)] = n
}

func RemoveCallback(msgID int, name string) {
	delete(msgCenter.callbacks[msgID], name)
}
func RemoveNotification(msgID int, notification Notification) {
	delete(msgCenter.notifications[msgID], fmt.Sprintf("%v", notification))
}

func SendMessage(msgID int, data interface{}) (result map[string]interface{}) {
	start := time.Now()
	if len(msgCenter.ch) < 30000 {
		msgCenter.ch <- MsgData{msgID, data}
	}
	// if len(msgCenter.ch)%10 == 0 {
	// 	fmt.Println("len(msgCenter.ch) =", len(msgCenter.ch))
	// }
	PrintDuration(fmt.Sprintf("push to msgCenter %v len=%v", msgID, len(msgCenter.ch)), start, time.Second)
	result = make(map[string]interface{})
	cs, ok := msgCenter.callbacks[msgID]
	if ok {
		for name, c := range cs {
			start := time.Now()
			result[name] = c(msgID, data)
			PrintDuration(fmt.Sprintf("call %v msgid=%v", name, msgID), start, time.Second)
		}
	}
	return
}

func dispatch() {
	for {
		msgData := <-msgCenter.ch
		nf, ok := msgCenter.notifications[msgData.msgID]
		if ok {
			for _, n := range nf {
				go n(msgData.msgID, msgData.data)
			}
		}
	}
}

//打印超过exceed时长的时间
//Prams:
// 	key: 用于识别的关键字
// 	start: 起始时间
// 	exceed: 持续时间超过多久才打印
func PrintDuration(key string, start time.Time, exceed time.Duration) {
	dur := time.Now().Sub(start)
	if dur >= exceed {
		fmt.Println("Duration", key, ":", dur.Seconds())
	}
}
