package request


type PlaceOrder struct {
	CouponId        int          `json:"coupon_id"`
	TotalPrice      float64      `json:"total_price"`
	FinalTotalPrice float64      `json:"final_total_price"`
	SkuInfoList     []SkuInfo    `json:"sku_info_list"`
	Address         OrderAddress `json:"address"`
}
type SkuInfo struct {
	Id    int `json:"id"`
	Count int `json:"count"`
}
type OrderAddress struct {
	UserName   string `json:"user_name"`
	Province   string `json:"province"`
	City       string `json:"city"`
	County     string `json:"county"`
	Mobile     string `json:"mobile"`
	NationCode string `json:"nation_code"`
	PostalCode string `json:"postal_code"`
	Detail     string `json:"detail"`
}
