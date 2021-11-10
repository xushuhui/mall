package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type BannerItem struct {
	ent.Schema
}

func (BannerItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "banner_item"},
	}
}
func (BannerItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("img").Comment(""),
		field.String("keyword").Comment(""),
		field.Int("type").Comment(""),

		field.Int("banner_id").Comment(""),
		field.String("name").Comment(""),
	}
}
func (BannerItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
