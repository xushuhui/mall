package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Order struct {
	ent.Schema
}

func (Order) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order"},
	}
}
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.String("order_no").Comment(""),
		field.Int("user_id").Comment("user表外键"),
		field.Float("total_price").Comment(""),
		field.Int("total_count").Comment(""),

		field.String("snap_img").Comment(""),
		field.String("snap_title").Comment(""),
		field.String("snap_items").Comment(""),
		field.String("snap_address").Comment(""),
		field.String("prepay_id").Comment(""),
		field.Float("final_total_price").Comment(""),
		field.Int8("status").Comment(""),
	}
}
func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
