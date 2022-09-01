// Code generated by entc, DO NOT EDIT.

package coupontemplate

import (
	"time"

	"mall-go/app/payment/service/internal/data/model/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v time.Time) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteTime), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// FullMoney applies equality check predicate on the "full_money" field. It's identical to FullMoneyEQ.
func FullMoney(v float64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFullMoney), v))
	})
}

// Minus applies equality check predicate on the "minus" field. It's identical to MinusEQ.
func Minus(v float64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMinus), v))
	})
}

// Discount applies equality check predicate on the "discount" field. It's identical to DiscountEQ.
func Discount(v float64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscount), v))
	})
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v int) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.CouponTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.CouponTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.CouponTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.CouponTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v time.Time) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v time.Time) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...time.Time) predicate.CouponTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeleteTime), v...))
	})
}

// DeleteTimeNotIn applies the NotIn predicate on the "delete_time" field.
func DeleteTimeNotIn(vs ...time.Time) predicate.CouponTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeleteTime), v...))
	})
}

// DeleteTimeGT applies the GT predicate on the "delete_time" field.
func DeleteTimeGT(v time.Time) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v time.Time) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v time.Time) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v time.Time) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeleteTime)))
	})
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeleteTime)))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.CouponTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.CouponTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.CouponTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.CouponTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// FullMoneyEQ applies the EQ predicate on the "full_money" field.
func FullMoneyEQ(v float64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFullMoney), v))
	})
}

// FullMoneyNEQ applies the NEQ predicate on the "full_money" field.
func FullMoneyNEQ(v float64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFullMoney), v))
	})
}

// FullMoneyIn applies the In predicate on the "full_money" field.
func FullMoneyIn(vs ...float64) predicate.CouponTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFullMoney), v...))
	})
}

// FullMoneyNotIn applies the NotIn predicate on the "full_money" field.
func FullMoneyNotIn(vs ...float64) predicate.CouponTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFullMoney), v...))
	})
}

// FullMoneyGT applies the GT predicate on the "full_money" field.
func FullMoneyGT(v float64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFullMoney), v))
	})
}

// FullMoneyGTE applies the GTE predicate on the "full_money" field.
func FullMoneyGTE(v float64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFullMoney), v))
	})
}

// FullMoneyLT applies the LT predicate on the "full_money" field.
func FullMoneyLT(v float64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFullMoney), v))
	})
}

// FullMoneyLTE applies the LTE predicate on the "full_money" field.
func FullMoneyLTE(v float64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFullMoney), v))
	})
}

// MinusEQ applies the EQ predicate on the "minus" field.
func MinusEQ(v float64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMinus), v))
	})
}

// MinusNEQ applies the NEQ predicate on the "minus" field.
func MinusNEQ(v float64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMinus), v))
	})
}

// MinusIn applies the In predicate on the "minus" field.
func MinusIn(vs ...float64) predicate.CouponTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMinus), v...))
	})
}

// MinusNotIn applies the NotIn predicate on the "minus" field.
func MinusNotIn(vs ...float64) predicate.CouponTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMinus), v...))
	})
}

// MinusGT applies the GT predicate on the "minus" field.
func MinusGT(v float64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMinus), v))
	})
}

// MinusGTE applies the GTE predicate on the "minus" field.
func MinusGTE(v float64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMinus), v))
	})
}

// MinusLT applies the LT predicate on the "minus" field.
func MinusLT(v float64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMinus), v))
	})
}

// MinusLTE applies the LTE predicate on the "minus" field.
func MinusLTE(v float64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMinus), v))
	})
}

// DiscountEQ applies the EQ predicate on the "discount" field.
func DiscountEQ(v float64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscount), v))
	})
}

// DiscountNEQ applies the NEQ predicate on the "discount" field.
func DiscountNEQ(v float64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDiscount), v))
	})
}

// DiscountIn applies the In predicate on the "discount" field.
func DiscountIn(vs ...float64) predicate.CouponTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDiscount), v...))
	})
}

// DiscountNotIn applies the NotIn predicate on the "discount" field.
func DiscountNotIn(vs ...float64) predicate.CouponTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDiscount), v...))
	})
}

// DiscountGT applies the GT predicate on the "discount" field.
func DiscountGT(v float64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDiscount), v))
	})
}

// DiscountGTE applies the GTE predicate on the "discount" field.
func DiscountGTE(v float64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDiscount), v))
	})
}

// DiscountLT applies the LT predicate on the "discount" field.
func DiscountLT(v float64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDiscount), v))
	})
}

// DiscountLTE applies the LTE predicate on the "discount" field.
func DiscountLTE(v float64) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDiscount), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v int) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v int) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...int) predicate.CouponTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...int) predicate.CouponTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v int) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	})
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v int) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	})
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v int) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	})
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v int) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CouponTemplate) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CouponTemplate) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CouponTemplate) predicate.CouponTemplate {
	return predicate.CouponTemplate(func(s *sql.Selector) {
		p(s.Not())
	})
}
