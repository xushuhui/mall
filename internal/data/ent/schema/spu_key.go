package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type SpuKey struct {
	ent.Schema
}

func (SpuKey) Fields() []ent.Field {
	return []ent.Field{
		field.Int("spu_id").Comment(""),
		field.Int("spec_key_id").Comment(""),
	}
}
