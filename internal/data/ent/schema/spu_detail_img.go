package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type SpuDetailImg struct {
	ent.Schema
}

func (SpuDetailImg) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "spu_detail_img"},
	}
}
func (SpuDetailImg) Fields() []ent.Field {
	return []ent.Field{
		field.String("img").Comment(""),
		field.Int("spu_id").Comment(""),
		field.Int("index").Comment(""),
	}
}
func (SpuDetailImg) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
