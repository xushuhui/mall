package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
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
		field.String("title").Comment("中文限制6个，英文限制12个，由逻辑层控制"),
		field.String("description").Comment(""),

		field.Int8("highlight").Comment(""),
		field.Int8("type").Comment(""),
	}
}
func (Tag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("spu", Spu.Type).Ref("tag"),
	}
}
