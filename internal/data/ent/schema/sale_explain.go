package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type SaleExplain struct {
	ent.Schema
}

func (SaleExplain) Fields() []ent.Field {
	return []ent.Field{
		field.Int8("fixed").Comment(""),
		field.String("text").Comment(""),
		field.Int("spu_id").Comment(""),
		field.Int("index").Comment(""),
		field.Int("replace_id").Comment(""),
		field.Time("create_time").Comment(""),
		field.Time("delete_time").Comment(""),
		field.Time("update_time").Comment(""),
	}
}
