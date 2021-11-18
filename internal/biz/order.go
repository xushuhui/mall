package biz

import (
	"mall-go/internal/logic"
	"mall-go/internal/request"

	"github.com/xushuhui/goal/core"
)

func PlaceOrder(userId int, req request.PlaceOrder, oc logic.OrderChecker) (orderId int, err error) {
	return
}

func OrderIsOk(userId int, orderDto request.PlaceOrder) (oc logic.OrderChecker, err error) {
	if orderDto.FinalTotalPrice <= 0 {
		err = core.ParamsError(core.InvalidParams)
		return
	}
	// TODO
	return
}
