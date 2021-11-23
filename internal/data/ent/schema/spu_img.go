package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type SpuImg struct {
	ent.Schema
}

func (SpuImg) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "spu_img"},
	}
}
func (SpuImg) Fields() []ent.Field {
	return []ent.Field{
		field.String("img").Comment(""),
		field.Int64("spu_id").Comment(""),
	}
}
func (SpuImg) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
