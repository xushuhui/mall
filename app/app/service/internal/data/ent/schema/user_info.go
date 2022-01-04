package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UserInfo struct {
	ent.Schema
}

func (UserInfo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_info"},
	}
}
func (UserInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Immutable().Comment(""),
		field.String("nickname").Comment(""),
		field.Int("unify_uid").Comment(""),
		field.String("email").Comment(""),
		field.String("avatar").Comment(""),
		field.String("wx_profile").Comment(""),
	}
}
func (UserInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
