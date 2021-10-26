package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Order struct {
	ent.Schema
}

func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.String("order_no").Comment(""),
		field.Int("user_id").Comment("user表外键"),
		field.Float("total_price").Comment(""),
		field.Int("total_count").Comment(""),
		field.Time("create_time").Comment(""),
		field.Time("delete_time").Comment(""),
		field.Time("update_time").Comment(""),
		field.String("snap_img").Comment(""),
		field.String("snap_title").Comment(""),
		field.String("snap_items").Comment(""),
		field.String("snap_address").Comment(""),
		field.String("prepay_id").Comment(""),
		field.Float("final_total_price").Comment(""),
		field.Int8("status").Comment(""),
	}
}
