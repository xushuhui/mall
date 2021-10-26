package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type GridCategory struct {
	ent.Schema
}

func (GridCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment(""),
		field.String("img").Comment(""),
		field.String("name").Comment(""),
		field.Time("create_time").Comment(""),
		field.Time("update_time").Comment(""),
		field.Time("delete_time").Comment(""),
		field.Int("category_id").Comment(""),
		field.Int("root_category_id").Comment(""),
	}
}
