package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type SpecKey struct {
	ent.Schema
}

func (SpecKey) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "spec_key"},
	}
}
func (SpecKey) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment(""),
		field.String("unit").Comment(""),
		field.Int8("standard").Comment(""),

		field.String("description").Comment(""),
	}
}
func (SpecKey) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
func (SpecKey) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("spu", Spu.Type).Ref("spec_key"),
	}
}
