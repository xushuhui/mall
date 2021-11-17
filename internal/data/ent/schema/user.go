package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user"},
	}
}
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("openid").Comment(""),
		field.String("nickname").Comment(""),
		field.Int("unify_uid").Comment(""),
		field.String("email").Comment(""),
		field.String("password").Comment(""),
		field.String("mobile").Comment(""),
		field.String("wx_profile").Comment(""),
	}
}
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_coupon", UserCoupon.Type),
	}
}
