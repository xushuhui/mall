package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type CouponTemplate struct {
	ent.Schema
}

func (CouponTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment(""),
		field.String("description").Comment(""),
		field.Float("full_money").Comment(""),
		field.Float("minus").Comment(""),
		field.Float("discount").Comment("国内多是打折，国外多是百分比 off"),
		field.Int("type").Comment("1. 满减券 2.折扣券 3.无门槛券 4.满金额折扣券"),
	}
}
func (CouponTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
