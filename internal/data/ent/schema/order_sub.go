package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type OrderSub struct {
	ent.Schema
}

func (OrderSub) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_sub"},
	}
}
func (OrderSub) Fields() []ent.Field {
	return []ent.Field{
		field.String("order_no").Comment(""),
		field.Int64("user_id").Optional().Comment("user表外键"),
		field.Float("price").Comment(""),
		field.Int("count").Comment(""),
	
		field.Float("final_price").Comment(""),
		field.Int8("status").Comment(""),
		field.Int64("order_id").Optional(),
	}
}
func (OrderSub) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
func (OrderSub) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).Ref("order_sub").Unique().Field("order_id"),
	}
}
