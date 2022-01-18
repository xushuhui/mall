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
		field.Int64("spu_id").Comment(""),
		field.Int64("sku_id").Comment(""),
		field.Int64("key_id").Comment(""),
		field.Int64("value_id").Comment(""),
	}
}
