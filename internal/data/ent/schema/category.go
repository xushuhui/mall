package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Category struct {
	ent.Schema
}

func (Category) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "category"},
	}
}
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment(""),
		field.String("description").Comment(""),

		field.Int8("is_root").Comment(""),
		field.Int("parent_id").Comment(""),
		field.String("img").Comment(""),
		field.Int("index").Comment(""),
		field.Int("online").Comment(""),
		field.Int("level").Comment(""),
	}
}
func (Category) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
