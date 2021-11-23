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
		field.Int64("activity_id").Comment(""),
		field.Int64("spu_id").Comment(""),
		field.Int8("participation").Comment(""),
	}
}
