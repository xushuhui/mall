package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Banner struct {
	ent.Schema
}

func (Banner) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment(""),
		field.String("description").Comment(""),
		field.Time("create_time").Comment(""),
		field.Time("update_time").Comment(""),
		field.Time("delete_time").Comment(""),
		field.String("title").Comment(""),
		field.String("img").Comment("部分banner可能有标题图片"),
	}
}
