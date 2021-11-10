package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Brand struct {
	ent.Schema
}

func (Brand) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "brand"},
	}
}
func (Brand) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment(""),
		field.String("description").Comment(""),
	}
}
func (Brand) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
