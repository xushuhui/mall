package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type SpecKey struct {
	ent.Schema
}

func (SpecKey) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment(""),
		field.String("unit").Comment(""),
		field.Int8("standard").Comment(""),
		field.Time("create_time").Comment(""),
		field.Time("update_time").Comment(""),
		field.Time("delete_time").Comment(""),
		field.String("description").Comment(""),
	}
}
