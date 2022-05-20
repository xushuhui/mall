package enum
//coupon_type
const (
	FULL_MINUS         = 1 // 满减
	FULL_OFF           = 2 // 满减折扣
	NO_THRESHOLD_MINUS = 3 // 无门槛减除
)
//coupon_status
const (
	AVAILABLE = 1// "未使用"), 
	USED = 2 //, "已使用"),
	 EXPIRED = 3 //"过期");
)
//login_type
const (
	WECHAT = 1
	EMAIL = 2
)
//order_status
const (
	ALL = iota //"全部"), 
    UNPAID // "未支付"),
    PAID  //"已支付"),
    DELIVERED //"已发货")
    FINISHED  //"已完成"),
    CANCELED //"已取消");
)