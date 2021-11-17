package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Theme struct {
	ent.Schema
}

func (Theme) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "theme"},
	}
}
func (Theme) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment(""),
		field.String("description").Comment(""),
		field.String("name").Comment(""),
		field.String("tpl_name").Comment(""),
		field.String("entrance_img").Comment(""),
		field.String("extend").Comment(""),
		field.String("internal_top_img").Comment(""),
		field.String("title_img").Comment(""),
		field.Int8("online").Comment(""),
	}
}
func (Theme) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
func (Theme) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("spu", Spu.Type).StorageKey(
			edge.Table("theme_spu"), edge.Columns("theme_id", "spu_id")),
	}
}
