package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type SpuImg struct {
	ent.Schema
}

func (SpuImg) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "spu_img"},
	}
}
func (SpuImg) Fields() []ent.Field {
	return []ent.Field{
		field.String("img").Comment(""),
		field.Int64("spu_id").Optional().Comment(""),
	}
}
func (SpuImg) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
func (SpuImg) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("spu", Spu.Type).
			Ref("spu_img").
			Unique().Field("spu_id"),
	}
}
