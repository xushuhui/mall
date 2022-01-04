package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
		field.Int64("banner_id").Optional(),

		field.String("name").Comment(""),
	}
}
func (BannerItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (BannerItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("banner", Banner.Type).
			Ref("banner_item").
			Unique().Field("banner_id"),
	}
}
