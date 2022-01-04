package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user"},
	}
}
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("openid").Comment(""),
		field.String("email").Comment(""),
		field.String("password").Comment(""),
		field.String("mobile").Comment(""),
		field.Int("status").Default(1).Comment(""),
	}
}
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
func (User) Edges() []ent.Edge {
	return []ent.Edge{

		edge.To("order", Order.Type),
		edge.To("refund", Refund.Type),
	}
}
