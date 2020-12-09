package model

type Coupon struct {
	BaseModel
	Title       string  `gorm:"column:title;" json:"title"`
	StartTime   string  `gorm:"column:start_time;" json:"start_time"`
	EndTime     string  `gorm:"column:end_time;" json:"end_time"`
	Description string  `gorm:"column:description;" json:"description"`
	FullMoney   float64 `gorm:"column:full_money;" json:"full_money"`
	Minus       float64 `gorm:"column:minus;" json:"minus"`
	Rate        float64 `gorm:"column:rate;" json:"rate"`
	Type        int32   `gorm:"column:type;" json:"type"`

	Valitiy    int32  `gorm:"column:valitiy;" json:"valitiy"`
	ActivityId int32  `gorm:"column:activity_id;" json:"activity_id"`
	Remark     string `gorm:"column:remark;" json:"remark"`
	WholeStore int32  `gorm:"column:whole_store;default:'0'" json:"whole_store"`
}
type CouponTemplate struct {
	BaseModel
	Title       string  `gorm:"column:title;" json:"title"`
	Description string  `gorm:"column:description;" json:"description"`
	FullMoney   float64 `gorm:"column:full_money;" json:"full_money"`
	Minus       float64 `gorm:"column:minus;" json:"minus"`
	Discount    float64 `gorm:"column:discount;" json:"discount"`
	Type        int     `gorm:"column:type;" json:"type"`
}
type CouponType struct {
	BaseModel
	Name        string `gorm:"column:name;" json:"name"`
	Code        int    `gorm:"column:code;" json:"code"`
	Description string `gorm:"column:description;" json:"description"`
}
type Brand struct {
	BaseModel
	Name        string `gorm:"column:name;" json:"name"`
	Description string `gorm:"column:description;" json:"description"`
}
