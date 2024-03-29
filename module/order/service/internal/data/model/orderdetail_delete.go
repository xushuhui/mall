// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"fmt"
	"mall-go/module/order/service/internal/data/model/orderdetail"
	"mall-go/module/order/service/internal/data/model/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderDetailDelete is the builder for deleting a OrderDetail entity.
type OrderDetailDelete struct {
	config
	hooks    []Hook
	mutation *OrderDetailMutation
}

// Where appends a list predicates to the OrderDetailDelete builder.
func (odd *OrderDetailDelete) Where(ps ...predicate.OrderDetail) *OrderDetailDelete {
	odd.mutation.Where(ps...)
	return odd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (odd *OrderDetailDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(odd.hooks) == 0 {
		affected, err = odd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			odd.mutation = mutation
			affected, err = odd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(odd.hooks) - 1; i >= 0; i-- {
			if odd.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = odd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, odd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (odd *OrderDetailDelete) ExecX(ctx context.Context) int {
	n, err := odd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (odd *OrderDetailDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: orderdetail.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: orderdetail.FieldID,
			},
		},
	}
	if ps := odd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, odd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// OrderDetailDeleteOne is the builder for deleting a single OrderDetail entity.
type OrderDetailDeleteOne struct {
	odd *OrderDetailDelete
}

// Exec executes the deletion query.
func (oddo *OrderDetailDeleteOne) Exec(ctx context.Context) error {
	n, err := oddo.odd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{orderdetail.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oddo *OrderDetailDeleteOne) ExecX(ctx context.Context) {
	oddo.odd.ExecX(ctx)
}
