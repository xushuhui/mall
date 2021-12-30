package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type SaleExplain struct {
	ent.Schema
}

func (SaleExplain) Fields() []ent.Field {
	return []ent.Field{
		field.Int8("fixed").Comment(""),
		field.String("text").Comment(""),
		field.Int64("spu_id").Optional().Comment(""),
		field.Int("index").Comment(""),
		field.Int64("replace_id").Comment(""),
	}
}
func (SaleExplain) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (SaleExplain) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("spu", Spu.Type).Ref("sale_explain").Unique().Field("spu_id"),
	}
}
