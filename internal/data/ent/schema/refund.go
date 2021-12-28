package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Refund struct {
	ent.Schema
}

func (Refund) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "refund"},
	}
}
func (Refund) Fields() []ent.Field {
	return []ent.Field{
		field.String("refund_no").Comment(""),
		field.Int64("user_id").Optional().Comment("user表外键"),
		field.String("reason").Comment(""),
		field.Int64("order_id").Optional(),
		field.Int64("order_sub_id").Optional(),
		field.Int8("status").Comment(""),
	}
}
func (Refund) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
func (Refund) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("refund").Unique().Field("user_id"),
	}
}
