package model

type Order struct {
	BaseModel
	OrderNo string `json:"order_no"`
	// UserID user表外键
	UserID     uint    `json:"user_id"`
	TotalPrice float32 `json:"total_price"`
	TotalCount uint    `json:"total_count"`

	SnapImg         string  `json:"snap_img"`
	SnapTitle       string  `json:"snap_title"`
	SnapItems       string  `json:"snap_items"`
	SnapAddress     string  `json:"snap_address"`
	PrepayID        string  `json:"prepay_id"`
	FinalTotalPrice float32 `json:"final_total_price"`
	Status          uint8   `json:"status"`
}
