package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type SpecValue struct {
	ent.Schema
}

func (SpecValue) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "spec_value"},
	}
}
func (SpecValue) Fields() []ent.Field {
	return []ent.Field{
		field.String("value").Comment(""),
		field.Int("spec_id").Comment(""),

		field.String("extend").Comment(""),
	}
}
func (SpecValue) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
