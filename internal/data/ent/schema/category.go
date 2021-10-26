package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Category struct {
	ent.Schema
}

func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment(""),
		field.String("description").Comment(""),
		field.Time("create_time").Comment(""),
		field.Time("update_time").Comment(""),
		field.Time("delete_time").Comment(""),
		field.Int8("is_root").Comment(""),
		field.Int("parent_id").Comment(""),
		field.String("img").Comment(""),
		field.Int("index").Comment(""),
		field.Int("online").Comment(""),
		field.Int("level").Comment(""),
	}
}
