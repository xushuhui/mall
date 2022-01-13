// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"
	"mall-go/app/app/service/internal/data/model"
)

// The ActivityFunc type is an adapter to allow the use of ordinary
// function as Activity mutator.
type ActivityFunc func(context.Context, *model.ActivityMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f ActivityFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.ActivityMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.ActivityMutation", m)
	}
	return f(ctx, mv)
}

// The BannerFunc type is an adapter to allow the use of ordinary
// function as Banner mutator.
type BannerFunc func(context.Context, *model.BannerMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f BannerFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.BannerMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.BannerMutation", m)
	}
	return f(ctx, mv)
}

// The BannerItemFunc type is an adapter to allow the use of ordinary
// function as BannerItem mutator.
type BannerItemFunc func(context.Context, *model.BannerItemMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f BannerItemFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.BannerItemMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.BannerItemMutation", m)
	}
	return f(ctx, mv)
}

// The BrandFunc type is an adapter to allow the use of ordinary
// function as Brand mutator.
type BrandFunc func(context.Context, *model.BrandMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f BrandFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.BrandMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.BrandMutation", m)
	}
	return f(ctx, mv)
}

// The CategoryFunc type is an adapter to allow the use of ordinary
// function as Category mutator.
type CategoryFunc func(context.Context, *model.CategoryMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f CategoryFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.CategoryMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.CategoryMutation", m)
	}
	return f(ctx, mv)
}

// The ChargeFunc type is an adapter to allow the use of ordinary
// function as Charge mutator.
type ChargeFunc func(context.Context, *model.ChargeMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f ChargeFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.ChargeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.ChargeMutation", m)
	}
	return f(ctx, mv)
}

// The CouponFunc type is an adapter to allow the use of ordinary
// function as Coupon mutator.
type CouponFunc func(context.Context, *model.CouponMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f CouponFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.CouponMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.CouponMutation", m)
	}
	return f(ctx, mv)
}

// The CouponTemplateFunc type is an adapter to allow the use of ordinary
// function as CouponTemplate mutator.
type CouponTemplateFunc func(context.Context, *model.CouponTemplateMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f CouponTemplateFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.CouponTemplateMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.CouponTemplateMutation", m)
	}
	return f(ctx, mv)
}

// The CouponTypeFunc type is an adapter to allow the use of ordinary
// function as CouponType mutator.
type CouponTypeFunc func(context.Context, *model.CouponTypeMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f CouponTypeFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.CouponTypeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.CouponTypeMutation", m)
	}
	return f(ctx, mv)
}

// The GridCategoryFunc type is an adapter to allow the use of ordinary
// function as GridCategory mutator.
type GridCategoryFunc func(context.Context, *model.GridCategoryMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f GridCategoryFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.GridCategoryMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.GridCategoryMutation", m)
	}
	return f(ctx, mv)
}

// The OrderFunc type is an adapter to allow the use of ordinary
// function as Order mutator.
type OrderFunc func(context.Context, *model.OrderMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f OrderFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.OrderMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.OrderMutation", m)
	}
	return f(ctx, mv)
}

// The OrderDetailFunc type is an adapter to allow the use of ordinary
// function as OrderDetail mutator.
type OrderDetailFunc func(context.Context, *model.OrderDetailMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f OrderDetailFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.OrderDetailMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.OrderDetailMutation", m)
	}
	return f(ctx, mv)
}

// The OrderSnapFunc type is an adapter to allow the use of ordinary
// function as OrderSnap mutator.
type OrderSnapFunc func(context.Context, *model.OrderSnapMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f OrderSnapFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.OrderSnapMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.OrderSnapMutation", m)
	}
	return f(ctx, mv)
}

// The OrderSubFunc type is an adapter to allow the use of ordinary
// function as OrderSub mutator.
type OrderSubFunc func(context.Context, *model.OrderSubMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f OrderSubFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.OrderSubMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.OrderSubMutation", m)
	}
	return f(ctx, mv)
}

// The RefundFunc type is an adapter to allow the use of ordinary
// function as Refund mutator.
type RefundFunc func(context.Context, *model.RefundMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f RefundFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.RefundMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.RefundMutation", m)
	}
	return f(ctx, mv)
}

// The SaleExplainFunc type is an adapter to allow the use of ordinary
// function as SaleExplain mutator.
type SaleExplainFunc func(context.Context, *model.SaleExplainMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f SaleExplainFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.SaleExplainMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.SaleExplainMutation", m)
	}
	return f(ctx, mv)
}

// The SkuFunc type is an adapter to allow the use of ordinary
// function as Sku mutator.
type SkuFunc func(context.Context, *model.SkuMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f SkuFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.SkuMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.SkuMutation", m)
	}
	return f(ctx, mv)
}

// The SkuSpecFunc type is an adapter to allow the use of ordinary
// function as SkuSpec mutator.
type SkuSpecFunc func(context.Context, *model.SkuSpecMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f SkuSpecFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.SkuSpecMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.SkuSpecMutation", m)
	}
	return f(ctx, mv)
}

// The SpecKeyFunc type is an adapter to allow the use of ordinary
// function as SpecKey mutator.
type SpecKeyFunc func(context.Context, *model.SpecKeyMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f SpecKeyFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.SpecKeyMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.SpecKeyMutation", m)
	}
	return f(ctx, mv)
}

// The SpecValueFunc type is an adapter to allow the use of ordinary
// function as SpecValue mutator.
type SpecValueFunc func(context.Context, *model.SpecValueMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f SpecValueFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.SpecValueMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.SpecValueMutation", m)
	}
	return f(ctx, mv)
}

// The SpuFunc type is an adapter to allow the use of ordinary
// function as Spu mutator.
type SpuFunc func(context.Context, *model.SpuMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f SpuFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.SpuMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.SpuMutation", m)
	}
	return f(ctx, mv)
}

// The SpuDetailImgFunc type is an adapter to allow the use of ordinary
// function as SpuDetailImg mutator.
type SpuDetailImgFunc func(context.Context, *model.SpuDetailImgMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f SpuDetailImgFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.SpuDetailImgMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.SpuDetailImgMutation", m)
	}
	return f(ctx, mv)
}

// The SpuImgFunc type is an adapter to allow the use of ordinary
// function as SpuImg mutator.
type SpuImgFunc func(context.Context, *model.SpuImgMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f SpuImgFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.SpuImgMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.SpuImgMutation", m)
	}
	return f(ctx, mv)
}

// The TagFunc type is an adapter to allow the use of ordinary
// function as Tag mutator.
type TagFunc func(context.Context, *model.TagMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f TagFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.TagMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.TagMutation", m)
	}
	return f(ctx, mv)
}

// The ThemeFunc type is an adapter to allow the use of ordinary
// function as Theme mutator.
type ThemeFunc func(context.Context, *model.ThemeMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f ThemeFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.ThemeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.ThemeMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, model.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m model.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m model.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m model.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op model.Op) Condition {
	return func(_ context.Context, m model.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m model.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m model.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m model.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk model.Hook, cond Condition) model.Hook {
	return func(next model.Mutator) model.Mutator {
		return model.MutateFunc(func(ctx context.Context, m model.Mutation) (model.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, model.Delete|model.Create)
//
func On(hk model.Hook, op model.Op) model.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, model.Update|model.UpdateOne)
//
func Unless(hk model.Hook, op model.Op) model.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) model.Hook {
	return func(model.Mutator) model.Mutator {
		return model.MutateFunc(func(context.Context, model.Mutation) (model.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []model.Hook {
//		return []model.Hook{
//			Reject(model.Delete|model.Update),
//		}
//	}
//
func Reject(op model.Op) model.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []model.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...model.Hook) Chain {
	return Chain{append([]model.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() model.Hook {
	return func(mutator model.Mutator) model.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...model.Hook) Chain {
	newHooks := make([]model.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
