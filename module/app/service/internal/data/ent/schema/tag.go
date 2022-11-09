package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Tag struct {
	ent.Schema
}

func (Tag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tag"},
	}
}

func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment(""),
		field.String("description").Comment(""),
		field.Int("highlight").Comment(""),
		field.Int("type").Comment(""),
	}
}

func (Tag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
