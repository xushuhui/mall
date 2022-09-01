// Code generated by entc, DO NOT EDIT.

package model

import (
	"time"

	"mall-go/app/payment/service/internal/data/ent/schema"
	"mall-go/app/payment/service/internal/data/model/coupon"
	"mall-go/app/payment/service/internal/data/model/coupontemplate"
	"mall-go/app/payment/service/internal/data/model/coupontype"
	"mall-go/app/payment/service/internal/data/model/usercoupon"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	couponMixin := schema.Coupon{}.Mixin()
	couponMixinFields0 := couponMixin[0].Fields()
	_ = couponMixinFields0
	couponFields := schema.Coupon{}.Fields()
	_ = couponFields
	// couponDescCreateTime is the schema descriptor for create_time field.
	couponDescCreateTime := couponMixinFields0[0].Descriptor()
	// coupon.DefaultCreateTime holds the default value on creation for the create_time field.
	coupon.DefaultCreateTime = couponDescCreateTime.Default.(func() time.Time)
	// couponDescUpdateTime is the schema descriptor for update_time field.
	couponDescUpdateTime := couponMixinFields0[1].Descriptor()
	// coupon.DefaultUpdateTime holds the default value on creation for the update_time field.
	coupon.DefaultUpdateTime = couponDescUpdateTime.Default.(func() time.Time)
	// coupon.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	coupon.UpdateDefaultUpdateTime = couponDescUpdateTime.UpdateDefault.(func() time.Time)
	coupontemplateMixin := schema.CouponTemplate{}.Mixin()
	coupontemplateMixinFields0 := coupontemplateMixin[0].Fields()
	_ = coupontemplateMixinFields0
	coupontemplateFields := schema.CouponTemplate{}.Fields()
	_ = coupontemplateFields
	// coupontemplateDescCreateTime is the schema descriptor for create_time field.
	coupontemplateDescCreateTime := coupontemplateMixinFields0[0].Descriptor()
	// coupontemplate.DefaultCreateTime holds the default value on creation for the create_time field.
	coupontemplate.DefaultCreateTime = coupontemplateDescCreateTime.Default.(func() time.Time)
	// coupontemplateDescUpdateTime is the schema descriptor for update_time field.
	coupontemplateDescUpdateTime := coupontemplateMixinFields0[1].Descriptor()
	// coupontemplate.DefaultUpdateTime holds the default value on creation for the update_time field.
	coupontemplate.DefaultUpdateTime = coupontemplateDescUpdateTime.Default.(func() time.Time)
	// coupontemplate.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	coupontemplate.UpdateDefaultUpdateTime = coupontemplateDescUpdateTime.UpdateDefault.(func() time.Time)
	coupontypeMixin := schema.CouponType{}.Mixin()
	coupontypeMixinFields0 := coupontypeMixin[0].Fields()
	_ = coupontypeMixinFields0
	coupontypeFields := schema.CouponType{}.Fields()
	_ = coupontypeFields
	// coupontypeDescCreateTime is the schema descriptor for create_time field.
	coupontypeDescCreateTime := coupontypeMixinFields0[0].Descriptor()
	// coupontype.DefaultCreateTime holds the default value on creation for the create_time field.
	coupontype.DefaultCreateTime = coupontypeDescCreateTime.Default.(func() time.Time)
	// coupontypeDescUpdateTime is the schema descriptor for update_time field.
	coupontypeDescUpdateTime := coupontypeMixinFields0[1].Descriptor()
	// coupontype.DefaultUpdateTime holds the default value on creation for the update_time field.
	coupontype.DefaultUpdateTime = coupontypeDescUpdateTime.Default.(func() time.Time)
	// coupontype.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	coupontype.UpdateDefaultUpdateTime = coupontypeDescUpdateTime.UpdateDefault.(func() time.Time)
	usercouponMixin := schema.UserCoupon{}.Mixin()
	usercouponMixinFields0 := usercouponMixin[0].Fields()
	_ = usercouponMixinFields0
	usercouponFields := schema.UserCoupon{}.Fields()
	_ = usercouponFields
	// usercouponDescCreateTime is the schema descriptor for create_time field.
	usercouponDescCreateTime := usercouponMixinFields0[0].Descriptor()
	// usercoupon.DefaultCreateTime holds the default value on creation for the create_time field.
	usercoupon.DefaultCreateTime = usercouponDescCreateTime.Default.(func() time.Time)
	// usercouponDescUpdateTime is the schema descriptor for update_time field.
	usercouponDescUpdateTime := usercouponMixinFields0[1].Descriptor()
	// usercoupon.DefaultUpdateTime holds the default value on creation for the update_time field.
	usercoupon.DefaultUpdateTime = usercouponDescUpdateTime.Default.(func() time.Time)
	// usercoupon.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	usercoupon.UpdateDefaultUpdateTime = usercouponDescUpdateTime.UpdateDefault.(func() time.Time)
	// usercouponDescStatus is the schema descriptor for status field.
	usercouponDescStatus := usercouponFields[2].Descriptor()
	// usercoupon.DefaultStatus holds the default value on creation for the status field.
	usercoupon.DefaultStatus = usercouponDescStatus.Default.(int)
}
