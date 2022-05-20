package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UserPointDetail struct {
	ent.Schema
}

func (UserPointDetail) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_point_detail"},
	}
}
func (UserPointDetail) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Comment(""),
		field.Int("value").Comment(""),
	}
}
func (UserPointDetail) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
