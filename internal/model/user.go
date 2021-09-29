package model

import (
	"time"
)

type User struct {
	ID       int64  `gorm:"primarykey"`
	UserName string `gorm:"column:userName"`
	Password string `gorm:"column:password"`
}

func GetAccountUserOne(where string, args ...interface{}) (model User, err error) {

	return

}
func GetUserByNames(names []string) (model []User, e error) {

	return
}
func GetUserByName(names string) (model User, e error) {

	return
}

type UserCoupon struct {
	ID       uint `json:"id"`
	UserID   uint `json:"user_id"`
	CouponID uint `json:"coupon_id"`
	// Status 1:未使用，2：已使用， 3：已过期
	Status     uint8     `json:"status"`
	CreateTime time.Time `json:"create_time"`
	OrderID    uint      `json:"order_id"`
	UpdateTime time.Time `json:"update_time"`
}
