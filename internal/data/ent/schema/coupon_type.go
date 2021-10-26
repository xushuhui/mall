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
		field.Time("create_time").Comment(""),
		field.Time("update_time").Comment(""),
		field.Time("delete_time").Comment(""),
	}
}
