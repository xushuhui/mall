package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type UserCoupon struct {
	ent.Schema
}

func (UserCoupon) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_coupon"},
	}
}
func (UserCoupon) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Comment(""),
		field.Int64("coupon_id").Optional().Comment(""),
		field.Int("status").Comment("1:未使用，2：已使用， 3：已过期").Default(1),

		field.Int("order_id").Comment(""),
	}
}
func (UserCoupon) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
func (UserCoupon) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("coupon", Coupon.Type).Ref("user_coupon").Unique().Field("coupon_id"),
	}
}
