package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type UserWallet struct {
	ent.Schema
}

func (UserWallet) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Immutable().Comment("用户ID"),
		field.Int("value").Comment("余额"),
	}
}

func (UserWallet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
