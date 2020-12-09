package model

import "time"

type Activity struct {
	BaseModel
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`

	Remark         string `json:"remark"`
	Online         uint8  `json:"online"`
	EntranceImg    string `json:"entrance_img"`
	InternalTopImg string `json:"internal_top_img"`
	Name           string `json:"name"`
}

type ActivityCategory struct {
	ID         uint   `json:"id"`
	CategoryID uint   `json:"category_id"`
	ActivityID uint32 `json:"activity_id"`
}
type ActivityCoupon struct {
	ID         uint `json:"id"`
	CouponID   uint `json:"coupon_id"`
	ActivityID uint `json:"activity_id"`
}
type ActivitySpu struct {
	ID            uint  `json:"id"`
	ActivityID    uint  `json:"activity_id"`
	SpuID         uint  `json:"spu_id"`
	Participation uint8 `json:"participation"`
}
