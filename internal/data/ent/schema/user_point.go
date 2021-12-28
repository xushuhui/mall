package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UserPoint struct {
	ent.Schema
}
func (UserPoint) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_point"},
	}
}
func (UserPoint) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Immutable().Comment(""),
		field.Int("value").Comment(""),
		field.Int8("status").Comment(""),
	}
}
func (UserPoint) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
