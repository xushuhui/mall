package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type SpuDetailImg struct {
	ent.Schema
}

func (SpuDetailImg) Fields() []ent.Field {
	return []ent.Field{
		field.String("img").Comment(""),
		field.Int("spu_id").Comment(""),
		field.Int("index").Comment(""),
		field.Time("create_time").Comment(""),
		field.Time("update_time").Comment(""),
		field.Time("delete_time").Comment(""),
	}
}
