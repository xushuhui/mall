package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("openid").Comment(""),
		field.String("nickname").Comment(""),
		field.Int("unify_uid").Comment(""),
		field.String("email").Comment(""),
		field.String("password").Comment(""),
		field.String("mobile").Comment(""),
		field.String("wx_profile").Comment(""),
		field.Time("create_time").Comment(""),
		field.Time("update_time").Comment(""),
		field.Time("delete_time").Comment(""),
	}
}
