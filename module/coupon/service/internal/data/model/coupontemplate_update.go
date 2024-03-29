// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/module/coupon/service/internal/data/model/coupontemplate"
	"mall-go/module/coupon/service/internal/data/model/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CouponTemplateUpdate is the builder for updating CouponTemplate entities.
type CouponTemplateUpdate struct {
	config
	hooks    []Hook
	mutation *CouponTemplateMutation
}

// Where appends a list predicates to the CouponTemplateUpdate builder.
func (ctu *CouponTemplateUpdate) Where(ps ...predicate.CouponTemplate) *CouponTemplateUpdate {
	ctu.mutation.Where(ps...)
	return ctu
}

// SetUpdateTime sets the "update_time" field.
func (ctu *CouponTemplateUpdate) SetUpdateTime(t time.Time) *CouponTemplateUpdate {
	ctu.mutation.SetUpdateTime(t)
	return ctu
}

// SetDeleteTime sets the "delete_time" field.
func (ctu *CouponTemplateUpdate) SetDeleteTime(t time.Time) *CouponTemplateUpdate {
	ctu.mutation.SetDeleteTime(t)
	return ctu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ctu *CouponTemplateUpdate) SetNillableDeleteTime(t *time.Time) *CouponTemplateUpdate {
	if t != nil {
		ctu.SetDeleteTime(*t)
	}
	return ctu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (ctu *CouponTemplateUpdate) ClearDeleteTime() *CouponTemplateUpdate {
	ctu.mutation.ClearDeleteTime()
	return ctu
}

// SetTitle sets the "title" field.
func (ctu *CouponTemplateUpdate) SetTitle(s string) *CouponTemplateUpdate {
	ctu.mutation.SetTitle(s)
	return ctu
}

// SetDescription sets the "description" field.
func (ctu *CouponTemplateUpdate) SetDescription(s string) *CouponTemplateUpdate {
	ctu.mutation.SetDescription(s)
	return ctu
}

// SetFullMoney sets the "full_money" field.
func (ctu *CouponTemplateUpdate) SetFullMoney(f float64) *CouponTemplateUpdate {
	ctu.mutation.ResetFullMoney()
	ctu.mutation.SetFullMoney(f)
	return ctu
}

// AddFullMoney adds f to the "full_money" field.
func (ctu *CouponTemplateUpdate) AddFullMoney(f float64) *CouponTemplateUpdate {
	ctu.mutation.AddFullMoney(f)
	return ctu
}

// SetMinus sets the "minus" field.
func (ctu *CouponTemplateUpdate) SetMinus(f float64) *CouponTemplateUpdate {
	ctu.mutation.ResetMinus()
	ctu.mutation.SetMinus(f)
	return ctu
}

// AddMinus adds f to the "minus" field.
func (ctu *CouponTemplateUpdate) AddMinus(f float64) *CouponTemplateUpdate {
	ctu.mutation.AddMinus(f)
	return ctu
}

// SetDiscount sets the "discount" field.
func (ctu *CouponTemplateUpdate) SetDiscount(f float64) *CouponTemplateUpdate {
	ctu.mutation.ResetDiscount()
	ctu.mutation.SetDiscount(f)
	return ctu
}

// AddDiscount adds f to the "discount" field.
func (ctu *CouponTemplateUpdate) AddDiscount(f float64) *CouponTemplateUpdate {
	ctu.mutation.AddDiscount(f)
	return ctu
}

// SetType sets the "type" field.
func (ctu *CouponTemplateUpdate) SetType(i int) *CouponTemplateUpdate {
	ctu.mutation.ResetType()
	ctu.mutation.SetType(i)
	return ctu
}

// AddType adds i to the "type" field.
func (ctu *CouponTemplateUpdate) AddType(i int) *CouponTemplateUpdate {
	ctu.mutation.AddType(i)
	return ctu
}

// Mutation returns the CouponTemplateMutation object of the builder.
func (ctu *CouponTemplateUpdate) Mutation() *CouponTemplateMutation {
	return ctu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ctu *CouponTemplateUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ctu.defaults()
	if len(ctu.hooks) == 0 {
		affected, err = ctu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CouponTemplateMutation)
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
func (ctu *CouponTemplateUpdate) SaveX(ctx context.Context) int {
	affected, err := ctu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ctu *CouponTemplateUpdate) Exec(ctx context.Context) error {
	_, err := ctu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctu *CouponTemplateUpdate) ExecX(ctx context.Context) {
	if err := ctu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ctu *CouponTemplateUpdate) defaults() {
	if _, ok := ctu.mutation.UpdateTime(); !ok {
		v := coupontemplate.UpdateDefaultUpdateTime()
		ctu.mutation.SetUpdateTime(v)
	}
}

func (ctu *CouponTemplateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coupontemplate.Table,
			Columns: coupontemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: coupontemplate.FieldID,
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
			Column: coupontemplate.FieldUpdateTime,
		})
	}
	if value, ok := ctu.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coupontemplate.FieldDeleteTime,
		})
	}
	if ctu.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: coupontemplate.FieldDeleteTime,
		})
	}
	if value, ok := ctu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coupontemplate.FieldTitle,
		})
	}
	if value, ok := ctu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coupontemplate.FieldDescription,
		})
	}
	if value, ok := ctu.mutation.FullMoney(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupontemplate.FieldFullMoney,
		})
	}
	if value, ok := ctu.mutation.AddedFullMoney(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupontemplate.FieldFullMoney,
		})
	}
	if value, ok := ctu.mutation.Minus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupontemplate.FieldMinus,
		})
	}
	if value, ok := ctu.mutation.AddedMinus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupontemplate.FieldMinus,
		})
	}
	if value, ok := ctu.mutation.Discount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupontemplate.FieldDiscount,
		})
	}
	if value, ok := ctu.mutation.AddedDiscount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupontemplate.FieldDiscount,
		})
	}
	if value, ok := ctu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coupontemplate.FieldType,
		})
	}
	if value, ok := ctu.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coupontemplate.FieldType,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ctu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coupontemplate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CouponTemplateUpdateOne is the builder for updating a single CouponTemplate entity.
type CouponTemplateUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CouponTemplateMutation
}

// SetUpdateTime sets the "update_time" field.
func (ctuo *CouponTemplateUpdateOne) SetUpdateTime(t time.Time) *CouponTemplateUpdateOne {
	ctuo.mutation.SetUpdateTime(t)
	return ctuo
}

// SetDeleteTime sets the "delete_time" field.
func (ctuo *CouponTemplateUpdateOne) SetDeleteTime(t time.Time) *CouponTemplateUpdateOne {
	ctuo.mutation.SetDeleteTime(t)
	return ctuo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ctuo *CouponTemplateUpdateOne) SetNillableDeleteTime(t *time.Time) *CouponTemplateUpdateOne {
	if t != nil {
		ctuo.SetDeleteTime(*t)
	}
	return ctuo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (ctuo *CouponTemplateUpdateOne) ClearDeleteTime() *CouponTemplateUpdateOne {
	ctuo.mutation.ClearDeleteTime()
	return ctuo
}

// SetTitle sets the "title" field.
func (ctuo *CouponTemplateUpdateOne) SetTitle(s string) *CouponTemplateUpdateOne {
	ctuo.mutation.SetTitle(s)
	return ctuo
}

// SetDescription sets the "description" field.
func (ctuo *CouponTemplateUpdateOne) SetDescription(s string) *CouponTemplateUpdateOne {
	ctuo.mutation.SetDescription(s)
	return ctuo
}

// SetFullMoney sets the "full_money" field.
func (ctuo *CouponTemplateUpdateOne) SetFullMoney(f float64) *CouponTemplateUpdateOne {
	ctuo.mutation.ResetFullMoney()
	ctuo.mutation.SetFullMoney(f)
	return ctuo
}

// AddFullMoney adds f to the "full_money" field.
func (ctuo *CouponTemplateUpdateOne) AddFullMoney(f float64) *CouponTemplateUpdateOne {
	ctuo.mutation.AddFullMoney(f)
	return ctuo
}

// SetMinus sets the "minus" field.
func (ctuo *CouponTemplateUpdateOne) SetMinus(f float64) *CouponTemplateUpdateOne {
	ctuo.mutation.ResetMinus()
	ctuo.mutation.SetMinus(f)
	return ctuo
}

// AddMinus adds f to the "minus" field.
func (ctuo *CouponTemplateUpdateOne) AddMinus(f float64) *CouponTemplateUpdateOne {
	ctuo.mutation.AddMinus(f)
	return ctuo
}

// SetDiscount sets the "discount" field.
func (ctuo *CouponTemplateUpdateOne) SetDiscount(f float64) *CouponTemplateUpdateOne {
	ctuo.mutation.ResetDiscount()
	ctuo.mutation.SetDiscount(f)
	return ctuo
}

// AddDiscount adds f to the "discount" field.
func (ctuo *CouponTemplateUpdateOne) AddDiscount(f float64) *CouponTemplateUpdateOne {
	ctuo.mutation.AddDiscount(f)
	return ctuo
}

// SetType sets the "type" field.
func (ctuo *CouponTemplateUpdateOne) SetType(i int) *CouponTemplateUpdateOne {
	ctuo.mutation.ResetType()
	ctuo.mutation.SetType(i)
	return ctuo
}

// AddType adds i to the "type" field.
func (ctuo *CouponTemplateUpdateOne) AddType(i int) *CouponTemplateUpdateOne {
	ctuo.mutation.AddType(i)
	return ctuo
}

// Mutation returns the CouponTemplateMutation object of the builder.
func (ctuo *CouponTemplateUpdateOne) Mutation() *CouponTemplateMutation {
	return ctuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ctuo *CouponTemplateUpdateOne) Select(field string, fields ...string) *CouponTemplateUpdateOne {
	ctuo.fields = append([]string{field}, fields...)
	return ctuo
}

// Save executes the query and returns the updated CouponTemplate entity.
func (ctuo *CouponTemplateUpdateOne) Save(ctx context.Context) (*CouponTemplate, error) {
	var (
		err  error
		node *CouponTemplate
	)
	ctuo.defaults()
	if len(ctuo.hooks) == 0 {
		node, err = ctuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CouponTemplateMutation)
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
		nv, ok := v.(*CouponTemplate)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CouponTemplateMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ctuo *CouponTemplateUpdateOne) SaveX(ctx context.Context) *CouponTemplate {
	node, err := ctuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ctuo *CouponTemplateUpdateOne) Exec(ctx context.Context) error {
	_, err := ctuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctuo *CouponTemplateUpdateOne) ExecX(ctx context.Context) {
	if err := ctuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ctuo *CouponTemplateUpdateOne) defaults() {
	if _, ok := ctuo.mutation.UpdateTime(); !ok {
		v := coupontemplate.UpdateDefaultUpdateTime()
		ctuo.mutation.SetUpdateTime(v)
	}
}

func (ctuo *CouponTemplateUpdateOne) sqlSave(ctx context.Context) (_node *CouponTemplate, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coupontemplate.Table,
			Columns: coupontemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: coupontemplate.FieldID,
			},
		},
	}
	id, ok := ctuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "CouponTemplate.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ctuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coupontemplate.FieldID)
		for _, f := range fields {
			if !coupontemplate.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != coupontemplate.FieldID {
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
			Column: coupontemplate.FieldUpdateTime,
		})
	}
	if value, ok := ctuo.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coupontemplate.FieldDeleteTime,
		})
	}
	if ctuo.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: coupontemplate.FieldDeleteTime,
		})
	}
	if value, ok := ctuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coupontemplate.FieldTitle,
		})
	}
	if value, ok := ctuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coupontemplate.FieldDescription,
		})
	}
	if value, ok := ctuo.mutation.FullMoney(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupontemplate.FieldFullMoney,
		})
	}
	if value, ok := ctuo.mutation.AddedFullMoney(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupontemplate.FieldFullMoney,
		})
	}
	if value, ok := ctuo.mutation.Minus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupontemplate.FieldMinus,
		})
	}
	if value, ok := ctuo.mutation.AddedMinus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupontemplate.FieldMinus,
		})
	}
	if value, ok := ctuo.mutation.Discount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupontemplate.FieldDiscount,
		})
	}
	if value, ok := ctuo.mutation.AddedDiscount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupontemplate.FieldDiscount,
		})
	}
	if value, ok := ctuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coupontemplate.FieldType,
		})
	}
	if value, ok := ctuo.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coupontemplate.FieldType,
		})
	}
	_node = &CouponTemplate{config: ctuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ctuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coupontemplate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
