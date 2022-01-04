package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type CouponType struct {
	ent.Schema
}

func (CouponType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment(""),
		field.Int("code").Comment(""),
		field.String("description").Comment(""),
	}
}
func (CouponType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
