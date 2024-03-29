// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/module/coupon/service/internal/data/model/coupontype"
	"mall-go/module/coupon/service/internal/data/model/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CouponTypeUpdate is the builder for updating CouponType entities.
type CouponTypeUpdate struct {
	config
	hooks    []Hook
	mutation *CouponTypeMutation
}

// Where appends a list predicates to the CouponTypeUpdate builder.
func (ctu *CouponTypeUpdate) Where(ps ...predicate.CouponType) *CouponTypeUpdate {
	ctu.mutation.Where(ps...)
	return ctu
}

// SetUpdateTime sets the "update_time" field.
func (ctu *CouponTypeUpdate) SetUpdateTime(t time.Time) *CouponTypeUpdate {
	ctu.mutation.SetUpdateTime(t)
	return ctu
}

// SetDeleteTime sets the "delete_time" field.
func (ctu *CouponTypeUpdate) SetDeleteTime(t time.Time) *CouponTypeUpdate {
	ctu.mutation.SetDeleteTime(t)
	return ctu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ctu *CouponTypeUpdate) SetNillableDeleteTime(t *time.Time) *CouponTypeUpdate {
	if t != nil {
		ctu.SetDeleteTime(*t)
	}
	return ctu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (ctu *CouponTypeUpdate) ClearDeleteTime() *CouponTypeUpdate {
	ctu.mutation.ClearDeleteTime()
	return ctu
}

// SetName sets the "name" field.
func (ctu *CouponTypeUpdate) SetName(s string) *CouponTypeUpdate {
	ctu.mutation.SetName(s)
	return ctu
}

// SetCode sets the "code" field.
func (ctu *CouponTypeUpdate) SetCode(i int) *CouponTypeUpdate {
	ctu.mutation.ResetCode()
	ctu.mutation.SetCode(i)
	return ctu
}

// AddCode adds i to the "code" field.
func (ctu *CouponTypeUpdate) AddCode(i int) *CouponTypeUpdate {
	ctu.mutation.AddCode(i)
	return ctu
}

// SetDescription sets the "description" field.
func (ctu *CouponTypeUpdate) SetDescription(s string) *CouponTypeUpdate {
	ctu.mutation.SetDescription(s)
	return ctu
}

// Mutation returns the CouponTypeMutation object of the builder.
func (ctu *CouponTypeUpdate) Mutation() *CouponTypeMutation {
	return ctu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ctu *CouponTypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ctu.defaults()
	if len(ctu.hooks) == 0 {
		affected, err = ctu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CouponTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ctu.mutation = mutation
			affected, err = ctu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ctu.hooks) - 1; i >= 0; i-- {
			if ctu.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = ctu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ctu *CouponTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := ctu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ctu *CouponTypeUpdate) Exec(ctx context.Context) error {
	_, err := ctu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctu *CouponTypeUpdate) ExecX(ctx context.Context) {
	if err := ctu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ctu *CouponTypeUpdate) defaults() {
	if _, ok := ctu.mutation.UpdateTime(); !ok {
		v := coupontype.UpdateDefaultUpdateTime()
		ctu.mutation.SetUpdateTime(v)
	}
}

func (ctu *CouponTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coupontype.Table,
			Columns: coupontype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: coupontype.FieldID,
			},
		},
	}
	if ps := ctu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coupontype.FieldUpdateTime,
		})
	}
	if value, ok := ctu.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coupontype.FieldDeleteTime,
		})
	}
	if ctu.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: coupontype.FieldDeleteTime,
		})
	}
	if value, ok := ctu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coupontype.FieldName,
		})
	}
	if value, ok := ctu.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coupontype.FieldCode,
		})
	}
	if value, ok := ctu.mutation.AddedCode(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coupontype.FieldCode,
		})
	}
	if value, ok := ctu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coupontype.FieldDescription,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ctu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coupontype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CouponTypeUpdateOne is the builder for updating a single CouponType entity.
type CouponTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CouponTypeMutation
}

// SetUpdateTime sets the "update_time" field.
func (ctuo *CouponTypeUpdateOne) SetUpdateTime(t time.Time) *CouponTypeUpdateOne {
	ctuo.mutation.SetUpdateTime(t)
	return ctuo
}

// SetDeleteTime sets the "delete_time" field.
func (ctuo *CouponTypeUpdateOne) SetDeleteTime(t time.Time) *CouponTypeUpdateOne {
	ctuo.mutation.SetDeleteTime(t)
	return ctuo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ctuo *CouponTypeUpdateOne) SetNillableDeleteTime(t *time.Time) *CouponTypeUpdateOne {
	if t != nil {
		ctuo.SetDeleteTime(*t)
	}
	return ctuo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (ctuo *CouponTypeUpdateOne) ClearDeleteTime() *CouponTypeUpdateOne {
	ctuo.mutation.ClearDeleteTime()
	return ctuo
}

// SetName sets the "name" field.
func (ctuo *CouponTypeUpdateOne) SetName(s string) *CouponTypeUpdateOne {
	ctuo.mutation.SetName(s)
	return ctuo
}

// SetCode sets the "code" field.
func (ctuo *CouponTypeUpdateOne) SetCode(i int) *CouponTypeUpdateOne {
	ctuo.mutation.ResetCode()
	ctuo.mutation.SetCode(i)
	return ctuo
}

// AddCode adds i to the "code" field.
func (ctuo *CouponTypeUpdateOne) AddCode(i int) *CouponTypeUpdateOne {
	ctuo.mutation.AddCode(i)
	return ctuo
}

// SetDescription sets the "description" field.
func (ctuo *CouponTypeUpdateOne) SetDescription(s string) *CouponTypeUpdateOne {
	ctuo.mutation.SetDescription(s)
	return ctuo
}

// Mutation returns the CouponTypeMutation object of the builder.
func (ctuo *CouponTypeUpdateOne) Mutation() *CouponTypeMutation {
	return ctuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ctuo *CouponTypeUpdateOne) Select(field string, fields ...string) *CouponTypeUpdateOne {
	ctuo.fields = append([]string{field}, fields...)
	return ctuo
}

// Save executes the query and returns the updated CouponType entity.
func (ctuo *CouponTypeUpdateOne) Save(ctx context.Context) (*CouponType, error) {
	var (
		err  error
		node *CouponType
	)
	ctuo.defaults()
	if len(ctuo.hooks) == 0 {
		node, err = ctuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CouponTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ctuo.mutation = mutation
			node, err = ctuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ctuo.hooks) - 1; i >= 0; i-- {
			if ctuo.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = ctuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ctuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CouponType)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CouponTypeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ctuo *CouponTypeUpdateOne) SaveX(ctx context.Context) *CouponType {
	node, err := ctuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ctuo *CouponTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := ctuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctuo *CouponTypeUpdateOne) ExecX(ctx context.Context) {
	if err := ctuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ctuo *CouponTypeUpdateOne) defaults() {
	if _, ok := ctuo.mutation.UpdateTime(); !ok {
		v := coupontype.UpdateDefaultUpdateTime()
		ctuo.mutation.SetUpdateTime(v)
	}
}

func (ctuo *CouponTypeUpdateOne) sqlSave(ctx context.Context) (_node *CouponType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coupontype.Table,
			Columns: coupontype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: coupontype.FieldID,
			},
		},
	}
	id, ok := ctuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "CouponType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ctuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coupontype.FieldID)
		for _, f := range fields {
			if !coupontype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != coupontype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ctuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coupontype.FieldUpdateTime,
		})
	}
	if value, ok := ctuo.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coupontype.FieldDeleteTime,
		})
	}
	if ctuo.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: coupontype.FieldDeleteTime,
		})
	}
	if value, ok := ctuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coupontype.FieldName,
		})
	}
	if value, ok := ctuo.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coupontype.FieldCode,
		})
	}
	if value, ok := ctuo.mutation.AddedCode(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coupontype.FieldCode,
		})
	}
	if value, ok := ctuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coupontype.FieldDescription,
		})
	}
	_node = &CouponType{config: ctuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ctuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coupontype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
