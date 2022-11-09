package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Refund struct {
	ent.Schema
}

func (Refund) Fields() []ent.Field {
	return []ent.Field{
		field.String("refund_no").Comment(""),
		field.String("transaction_id").Comment("支付平台流水号"),
		field.Int64("user_id").Optional().Comment("user表外键"),
		field.String("reason").Comment(""),
		field.Int64("order_id").Optional(),
		field.Int64("order_sub_id").Optional(),
		field.Int("status").Comment(""),
	}
}

func (Refund) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
