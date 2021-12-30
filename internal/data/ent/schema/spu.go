package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Spu struct {
	ent.Schema
}

func (Spu) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment(""),
		field.String("subtitle").Comment(""),
		field.Int64("category_id").Comment(""),
		field.Int64("root_category_id").Comment(""),

		field.Int8("online").Comment(""),

		field.String("price").Comment("文本型价格，有时候SPU需要展示的是一个范围，或者自定义平均价格"),
		field.Int("sketch_spec_id").Comment("某种规格可以直接附加单品图片"),
		field.Int("default_sku_id").Comment("默认选中的sku"),
		field.String("img").Comment(""),
		field.String("discount_price").Comment(""),
		field.String("description").Comment(""),
		field.String("tags").Comment(""),

		field.String("spu_theme_img").Comment(""),
		field.String("for_theme_img").Comment(""),
	}
}
func (Spu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
func (Spu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sale_explain", SaleExplain.Type),
		edge.To("spu_img", SpuImg.Type),
		edge.To("spu_detail_img", SpuDetailImg.Type),
		edge.To("spec_key", SpecKey.Type).StorageKey(
			edge.Table("spu_key"), edge.Columns("spu_id", "spec_key_id")),
		edge.To("tag", Tag.Type).StorageKey(
			edge.Table("spu_tag"), edge.Columns("spu_id", "tag_id")),
		edge.From("theme", Theme.Type).Ref("spu"),
		edge.From("activity", Activity.Type).Ref("spu"),
		edge.From("brand", Brand.Type).Ref("spu"),
	}
}
