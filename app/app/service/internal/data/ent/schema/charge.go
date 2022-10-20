package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Charge struct {
	ent.Schema
}

func (Charge) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Comment("user表外键"),
		field.String("amount").Comment("充值金额"),
		field.String("charge_no").Comment("充值单号"),
		field.String("transaction_id").Comment("支付平台流水号"),
		field.Int("pay_way").Comment("支付方式：1微信支付，2支付宝支付"),
		field.Int("client_type").Comment("客户端类型：1安卓，2IOS，3PC"),
	}
}

func (Charge) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
