package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Sku struct {
	ent.Schema
}

func (Sku) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sku"},
	}
}
func (Sku) Fields() []ent.Field {
	return []ent.Field{
		field.Float("price").Comment(""),
		field.Float("discount_price").Comment(""),
		field.Int8("online").Comment(""),
		field.String("img").Comment(""),
		field.String("title").Comment(""),
		field.Int64("spu_id").Optional().Comment(""),

		field.JSON("specs", []Spec{}).Comment(""),
		field.String("code").Comment(""),
		field.Int("stock").Comment(""),
		field.Int64("category_id").Comment(""),
		field.Int64("root_category_id").Comment(""),
	}
}
func (Sku) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
