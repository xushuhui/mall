package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Sku struct {
	ent.Schema
}

func (Sku) Fields() []ent.Field {
	return []ent.Field{
		field.Float("price").Comment(""),
		field.Float("discount_price").Comment(""),
		field.Int8("online").Comment(""),
		field.String("img").Comment(""),
		field.String("title").Comment(""),
		field.Int("spu_id").Comment(""),
		field.Time("create_time").Comment(""),
		field.Time("update_time").Comment(""),
		field.Time("delete_time").Comment(""),
		field.String("specs").Comment(""),
		field.String("code").Comment(""),
		field.Int("stock").Comment(""),
		field.Int("category_id").Comment(""),
		field.Int("root_category_id").Comment(""),
	}
}
