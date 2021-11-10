package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Banner struct {
	ent.Schema
}

func (Banner) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "banner"},
	}
}
func (Banner) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment(""),
		field.String("description").Comment(""),

		field.String("title").Comment(""),
		field.String("img").Comment("部分banner可能有标题图片"),
	}
}
func (Banner) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
func (Banner) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("item", BannerItem.Type),
	}
}
