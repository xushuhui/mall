package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Order struct {
	ent.Schema
}

func (Order) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order"},
	}
}
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.String("order_no").Comment(""),
		field.Int64("user_id").Optional().Comment("user表外键"),
		field.Float("total_price").Comment(""),
		field.Int("total_count").Comment(""),
	
		field.Float("final_total_price").Comment(""),
		field.Int8("status").Comment(""),
	}
}
func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("order").Unique().Field("user_id"),
		edge.To("order_snap", OrderSnap.Type),
		edge.To("order_sub", OrderSub.Type),
	}
}
