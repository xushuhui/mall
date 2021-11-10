package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type SaleExplain struct {
	ent.Schema
}

func (SaleExplain) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sale_explain"},
	}
}
func (SaleExplain) Fields() []ent.Field {
	return []ent.Field{
		field.Int8("fixed").Comment(""),
		field.String("text").Comment(""),
		field.Int("spu_id").Comment(""),
		field.Int("index").Comment(""),
		field.Int("replace_id").Comment(""),
	}
}
func (SaleExplain) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
