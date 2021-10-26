package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type SpecValue struct {
	ent.Schema
}

func (SpecValue) Fields() []ent.Field {
	return []ent.Field{
		field.String("value").Comment(""),
		field.Int("spec_id").Comment(""),
		field.Time("create_time").Comment(""),
		field.Time("update_time").Comment(""),
		field.Time("delete_time").Comment(""),
		field.String("extend").Comment(""),
	}
}
