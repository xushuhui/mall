// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"fmt"
	"mall-go/module/coupon/service/internal/data/model/coupontemplate"
	"mall-go/module/coupon/service/internal/data/model/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CouponTemplateDelete is the builder for deleting a CouponTemplate entity.
type CouponTemplateDelete struct {
	config
	hooks    []Hook
	mutation *CouponTemplateMutation
}

// Where appends a list predicates to the CouponTemplateDelete builder.
func (ctd *CouponTemplateDelete) Where(ps ...predicate.CouponTemplate) *CouponTemplateDelete {
	ctd.mutation.Where(ps...)
	return ctd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ctd *CouponTemplateDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ctd.hooks) == 0 {
		affected, err = ctd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CouponTemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ctd.mutation = mutation
			affected, err = ctd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ctd.hooks) - 1; i >= 0; i-- {
			if ctd.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = ctd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctd *CouponTemplateDelete) ExecX(ctx context.Context) int {
	n, err := ctd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ctd *CouponTemplateDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: coupontemplate.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: coupontemplate.FieldID,
			},
		},
	}
	if ps := ctd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ctd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// CouponTemplateDeleteOne is the builder for deleting a single CouponTemplate entity.
type CouponTemplateDeleteOne struct {
	ctd *CouponTemplateDelete
}

// Exec executes the deletion query.
func (ctdo *CouponTemplateDeleteOne) Exec(ctx context.Context) error {
	n, err := ctdo.ctd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{coupontemplate.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ctdo *CouponTemplateDeleteOne) ExecX(ctx context.Context) {
	ctdo.ctd.ExecX(ctx)
}
