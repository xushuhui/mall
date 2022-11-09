package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Brand struct {
	ent.Schema
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

func (Brand) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("spu", Spu.Type),
	}
}
