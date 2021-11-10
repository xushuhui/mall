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
		field.Int("spu_id").Comment(""),

		field.String("specs").Comment(""),
		field.String("code").Comment(""),
		field.Int("stock").Comment(""),
		field.Int("category_id").Comment(""),
		field.Int("root_category_id").Comment(""),
	}
}
func (Sku) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
