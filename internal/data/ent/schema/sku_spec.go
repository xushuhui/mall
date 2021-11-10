package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type SkuSpec struct {
	ent.Schema
}

func (SkuSpec) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sku_spec"},
	}
}
func (SkuSpec) Fields() []ent.Field {
	return []ent.Field{
		field.Int("spu_id").Comment(""),
		field.Int("sku_id").Comment(""),
		field.Int("key_id").Comment(""),
		field.Int("value_id").Comment(""),
	}
}
