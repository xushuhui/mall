package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type UserWalletDetail struct {
	ent.Schema
}

func (UserWalletDetail) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Comment(""),
		field.Int8("op").Comment("1 增加 2减少").Default(1),
		field.Int("current").Comment("当前余额"),
		field.Int("value").Comment("余额"),
		field.Int8("type").Comment("1充值，2消费").Default(1),
	}
}
func (UserWalletDetail) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
