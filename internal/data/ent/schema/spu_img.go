package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type SpuImg struct {
	ent.Schema
}

func (SpuImg) Fields() []ent.Field {
	return []ent.Field{
		field.String("img").Comment(""),
		field.Int("spu_id").Comment(""),
		field.Time("delete_time").Comment(""),
		field.Time("update_time").Comment(""),
		field.Time("create_time").Comment(""),
	}
}
