package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Spu struct {
	ent.Schema
}

func (Spu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "spu"},
	}
}
func (Spu) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment(""),
		field.String("subtitle").Comment(""),
		field.Int("category_id").Comment(""),
		field.Int("root_category_id").Comment(""),
		field.Int8("online").Comment(""),

		field.String("price").Comment("文本型价格，有时候SPU需要展示的是一个范围，或者自定义平均价格"),
		field.Int("sketch_spec_id").Comment("某种规格可以直接附加单品图片"),
		field.Int("default_sku_id").Comment("默认选中的sku"),
		field.String("img").Comment(""),
		field.String("discount_price").Comment(""),
		field.String("description").Comment(""),
		field.String("tags").Comment(""),
		field.Int8("is_test").Comment(""),
		field.String("spu_theme_img").Comment(""),
		field.String("for_theme_img").Comment(""),
	}
}
func (Spu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
