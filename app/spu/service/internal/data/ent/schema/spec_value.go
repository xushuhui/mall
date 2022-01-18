package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type SpecValue struct {
	ent.Schema
}

func (SpecValue) Fields() []ent.Field {
	return []ent.Field{
		field.String("value").Comment(""),
		field.Int64("spec_id").Comment(""),

		field.String("extend").Comment(""),
	}
}
func (SpecValue) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
