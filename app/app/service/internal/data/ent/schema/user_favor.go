package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type UserFavor struct {
	ent.Schema
}

func (UserFavor) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_favor"},
	}
}
func (UserFavor) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Comment(""),
		field.Int64("spu_id").Comment(""),
		field.Int("status").Default(1).Comment("1收藏 0取消收藏"),
	}
}
func (UserFavor) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
