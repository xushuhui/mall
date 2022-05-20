package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UserIdentiy struct {
	ent.Schema
}

func (UserIdentiy) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "lin_user_identiy"},
	}
}
func (UserIdentiy) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Comment("用户id"),
		field.String("identity_type").Comment("phone,weapp"),
		field.String("identifier").Comment(""),
		field.String("credential").Comment(""),
	}
}

func (UserIdentiy) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
