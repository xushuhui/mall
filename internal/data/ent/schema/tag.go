package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Tag struct {
	ent.Schema
}

func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment("中文限制6个，英文限制12个，由逻辑层控制"),
		field.String("description").Comment(""),
		field.Time("update_time").Comment(""),
		field.Time("delete_time").Comment(""),
		field.Time("create_time").Comment(""),
		field.Int8("highlight").Comment(""),
		field.Int8("type").Comment(""),
	}
}
