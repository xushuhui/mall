package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Activity struct {
	ent.Schema
}

func (Activity) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment(""),
		field.String("description").Comment(""),
		field.Time("start_time").Comment(""),
		field.Time("end_time").Comment(""),
		field.Time("create_time").Comment(""),
		field.Time("update_time").Comment(""),
		field.Time("delete_time").Comment(""),
		field.String("remark").Comment(""),
		field.Int8("online").Comment(""),
		field.String("entrance_img").Comment(""),
		field.String("internal_top_img").Comment(""),
		field.String("name").Comment(""),
	}
}
