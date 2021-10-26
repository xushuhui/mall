package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Theme struct {
	ent.Schema
}

func (Theme) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment(""),
		field.String("description").Comment(""),
		field.String("name").Comment(""),
		field.Time("create_time").Comment(""),
		field.String("tpl_name").Comment(""),
		field.Time("update_time").Comment(""),
		field.Time("delete_time").Comment(""),
		field.String("entrance_img").Comment(""),
		field.String("extend").Comment(""),
		field.String("internal_top_img").Comment(""),
		field.String("title_img").Comment(""),
		field.Int8("online").Comment(""),
	}
}
