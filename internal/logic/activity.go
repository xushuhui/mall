package logic

type ActivityItem struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	EntranceImg string      `json:"entrance_img"`
	Online      bool        `json:"online"`
	Remark      string      `json:"remark"`
	StartTime   interface{} `json:"start_time"`
	EndTime     interface{} `json:"end_time"`
}
type ActivityWithCoupon struct {
	ActivityItem
	Coupons []CouponItem `json:"coupons"`
}
type CouponItem struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	StartTime   int64   `json:"start_time"`
	EndTime     int64   `json:"end_time"`
	Description string  `json:"description"`
	FullMoney   int     `json:"full_money"`
	Minus       float64 `json:"minus"`
	Rate        float64 `json:"rate"`
	Type        int     `json:"type"`
	Remark      string  `json:"remark"`
	WholeStore  bool    `json:"whole_store"`
}
