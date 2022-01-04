package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type GridCategory struct {
	ent.Schema
}

func (GridCategory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "grid_category"},
	}
}
func (GridCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment(""),
		field.String("img").Comment(""),
		field.String("name").Comment(""),

		field.Int("category_id").Comment(""),
		field.Int("root_category_id").Comment(""),
	}
}
func (GridCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
