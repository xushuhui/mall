package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type ActivitySpu struct {
	ent.Schema
}

func (ActivitySpu) Fields() []ent.Field {
	return []ent.Field{
		field.Int("activity_id").Comment(""),
		field.Int("spu_id").Comment(""),
		field.Int8("participation").Comment(""),
	}
}
