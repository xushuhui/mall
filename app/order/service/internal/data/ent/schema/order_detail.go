package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type OrderDetail struct {
	ent.Schema
}

func (OrderDetail) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Comment("user表外键"),

		field.Int("pay_way").Comment("支付方式：1微信支付，2支付宝支付，3余额支付").Default(1),
		field.Int("client_type").Comment("客户端类型：1安卓，2IOS，3PC").Default(1),
		field.String("ship_no").Comment("物流单号"),
	}
}

func (OrderDetail) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
