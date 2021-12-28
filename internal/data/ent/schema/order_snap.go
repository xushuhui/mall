package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type OrderSnap struct {
	ent.Schema
}

func (OrderSnap) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_snap"},
	}
}
func (OrderSnap) Fields() []ent.Field {
	return []ent.Field{
		field.String("snap_img").Comment(""),
		field.String("snap_title").Comment(""),
		field.String("snap_items").Comment(""),
		field.String("snap_address").Comment(""),
		field.Int64("order_id").Optional(),
	}
}
func (OrderSnap) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).Ref("order_snap").Unique().Field("order_id"),
	}
}
