package model

import (
	"gorm.io/gorm"
	"mall_go/global"
)

type User struct {
	gorm.Model
	Phone    string `gorm:"column:phone"`
	Password string `gorm:"column:password"`
}

func GetAccountUserOne(where string, args ...interface{}) (model User, err error) {
	err = global.DBEngine.First(&model, where, args).Error
	return

}
