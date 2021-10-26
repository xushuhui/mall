package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type UserCoupon struct {
	ent.Schema
}

func (UserCoupon) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Comment("").Unique(),
		field.Int("coupon_id").Comment(""),
		field.Int8("status").Comment("1:未使用，2：已使用， 3：已过期"),
		field.Time("create_time").Comment(""),
		field.Int("order_id").Comment(""),
		field.Time("update_time").Comment(""),
	}
}
