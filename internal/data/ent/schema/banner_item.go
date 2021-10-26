package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type BannerItem struct {
	ent.Schema
}

func (BannerItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("img").Comment(""),
		field.String("keyword").Comment(""),
		field.Int("type").Comment(""),
		field.Time("create_time").Comment(""),
		field.Time("update_time").Comment(""),
		field.Time("delete_time").Comment(""),
		field.Int("banner_id").Comment(""),
		field.String("name").Comment(""),
	}
}
