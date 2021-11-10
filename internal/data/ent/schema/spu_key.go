package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type SpuKey struct {
	ent.Schema
}

func (SpuKey) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "spu_key"},
	}
}
func (SpuKey) Fields() []ent.Field {
	return []ent.Field{
		field.Int("spu_id").Comment(""),
		field.Int("spec_key_id").Comment(""),
	}
}
