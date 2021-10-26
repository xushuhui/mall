package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type SkuSpec struct {
	ent.Schema
}

func (SkuSpec) Fields() []ent.Field {
	return []ent.Field{
		field.Int("spu_id").Comment(""),
		field.Int("sku_id").Comment(""),
		field.Int("key_id").Comment(""),
		field.Int("value_id").Comment(""),
	}
}
