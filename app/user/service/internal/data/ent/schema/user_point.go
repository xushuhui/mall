package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type UserPoint struct {
	ent.Schema
}

func (UserPoint) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Immutable().Comment("user_id"),
		field.Int("value").Comment(""),
		field.Int("status").Comment(""),
	}
}
func (UserPoint) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
