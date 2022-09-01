// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"mall-go/app/payment/service/internal/data/model/coupon"
	"mall-go/app/payment/service/internal/data/model/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CouponUpdate is the builder for updating Coupon entities.
type CouponUpdate struct {
	config
	hooks    []Hook
	mutation *CouponMutation
}

// Where appends a list predicates to the CouponUpdate builder.
func (cu *CouponUpdate) Where(ps ...predicate.Coupon) *CouponUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdateTime sets the "update_time" field.
func (cu *CouponUpdate) SetUpdateTime(t time.Time) *CouponUpdate {
	cu.mutation.SetUpdateTime(t)
	return cu
}

// SetDeleteTime sets the "delete_time" field.
func (cu *CouponUpdate) SetDeleteTime(t time.Time) *CouponUpdate {
	cu.mutation.SetDeleteTime(t)
	return cu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (cu *CouponUpdate) SetNillableDeleteTime(t *time.Time) *CouponUpdate {
	if t != nil {
		cu.SetDeleteTime(*t)
	}
	return cu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (cu *CouponUpdate) ClearDeleteTime() *CouponUpdate {
	cu.mutation.ClearDeleteTime()
	return cu
}

// SetTitle sets the "title" field.
func (cu *CouponUpdate) SetTitle(s string) *CouponUpdate {
	cu.mutation.SetTitle(s)
	return cu
}

// SetStartTime sets the "start_time" field.
func (cu *CouponUpdate) SetStartTime(t time.Time) *CouponUpdate {
	cu.mutation.SetStartTime(t)
	return cu
}

// SetEndTime sets the "end_time" field.
func (cu *CouponUpdate) SetEndTime(t time.Time) *CouponUpdate {
	cu.mutation.SetEndTime(t)
	return cu
}

// SetDescription sets the "description" field.
func (cu *CouponUpdate) SetDescription(s string) *CouponUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetFullMoney sets the "full_money" field.
func (cu *CouponUpdate) SetFullMoney(f float64) *CouponUpdate {
	cu.mutation.ResetFullMoney()
	cu.mutation.SetFullMoney(f)
	return cu
}

// AddFullMoney adds f to the "full_money" field.
func (cu *CouponUpdate) AddFullMoney(f float64) *CouponUpdate {
	cu.mutation.AddFullMoney(f)
	return cu
}

// SetMinus sets the "minus" field.
func (cu *CouponUpdate) SetMinus(f float64) *CouponUpdate {
	cu.mutation.ResetMinus()
	cu.mutation.SetMinus(f)
	return cu
}

// AddMinus adds f to the "minus" field.
func (cu *CouponUpdate) AddMinus(f float64) *CouponUpdate {
	cu.mutation.AddMinus(f)
	return cu
}

// SetRate sets the "rate" field.
func (cu *CouponUpdate) SetRate(f float64) *CouponUpdate {
	cu.mutation.ResetRate()
	cu.mutation.SetRate(f)
	return cu
}

// AddRate adds f to the "rate" field.
func (cu *CouponUpdate) AddRate(f float64) *CouponUpdate {
	cu.mutation.AddRate(f)
	return cu
}

// SetType sets the "type" field.
func (cu *CouponUpdate) SetType(i int) *CouponUpdate {
	cu.mutation.ResetType()
	cu.mutation.SetType(i)
	return cu
}

// AddType adds i to the "type" field.
func (cu *CouponUpdate) AddType(i int) *CouponUpdate {
	cu.mutation.AddType(i)
	return cu
}

// SetValitiy sets the "valitiy" field.
func (cu *CouponUpdate) SetValitiy(i int) *CouponUpdate {
	cu.mutation.ResetValitiy()
	cu.mutation.SetValitiy(i)
	return cu
}

// AddValitiy adds i to the "valitiy" field.
func (cu *CouponUpdate) AddValitiy(i int) *CouponUpdate {
	cu.mutation.AddValitiy(i)
	return cu
}

// SetActivityID sets the "activity_id" field.
func (cu *CouponUpdate) SetActivityID(i int64) *CouponUpdate {
	cu.mutation.ResetActivityID()
	cu.mutation.SetActivityID(i)
	return cu
}

// SetNillableActivityID sets the "activity_id" field if the given value is not nil.
func (cu *CouponUpdate) SetNillableActivityID(i *int64) *CouponUpdate {
	if i != nil {
		cu.SetActivityID(*i)
	}
	return cu
}

// AddActivityID adds i to the "activity_id" field.
func (cu *CouponUpdate) AddActivityID(i int64) *CouponUpdate {
	cu.mutation.AddActivityID(i)
	return cu
}

// ClearActivityID clears the value of the "activity_id" field.
func (cu *CouponUpdate) ClearActivityID() *CouponUpdate {
	cu.mutation.ClearActivityID()
	return cu
}

// SetRemark sets the "remark" field.
func (cu *CouponUpdate) SetRemark(s string) *CouponUpdate {
	cu.mutation.SetRemark(s)
	return cu
}

// SetWholeStore sets the "whole_store" field.
func (cu *CouponUpdate) SetWholeStore(i int) *CouponUpdate {
	cu.mutation.ResetWholeStore()
	cu.mutation.SetWholeStore(i)
	return cu
}

// AddWholeStore adds i to the "whole_store" field.
func (cu *CouponUpdate) AddWholeStore(i int) *CouponUpdate {
	cu.mutation.AddWholeStore(i)
	return cu
}

// Mutation returns the CouponMutation object of the builder.
func (cu *CouponUpdate) Mutation() *CouponMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CouponUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CouponMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CouponUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CouponUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CouponUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CouponUpdate) defaults() {
	if _, ok := cu.mutation.UpdateTime(); !ok {
		v := coupon.UpdateDefaultUpdateTime()
		cu.mutation.SetUpdateTime(v)
	}
}

func (cu *CouponUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coupon.Table,
			Columns: coupon.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: coupon.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coupon.FieldUpdateTime,
		})
	}
	if value, ok := cu.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coupon.FieldDeleteTime,
		})
	}
	if cu.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: coupon.FieldDeleteTime,
		})
	}
	if value, ok := cu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coupon.FieldTitle,
		})
	}
	if value, ok := cu.mutation.StartTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coupon.FieldStartTime,
		})
	}
	if value, ok := cu.mutation.EndTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coupon.FieldEndTime,
		})
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coupon.FieldDescription,
		})
	}
	if value, ok := cu.mutation.FullMoney(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupon.FieldFullMoney,
		})
	}
	if value, ok := cu.mutation.AddedFullMoney(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupon.FieldFullMoney,
		})
	}
	if value, ok := cu.mutation.Minus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupon.FieldMinus,
		})
	}
	if value, ok := cu.mutation.AddedMinus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupon.FieldMinus,
		})
	}
	if value, ok := cu.mutation.Rate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupon.FieldRate,
		})
	}
	if value, ok := cu.mutation.AddedRate(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupon.FieldRate,
		})
	}
	if value, ok := cu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coupon.FieldType,
		})
	}
	if value, ok := cu.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coupon.FieldType,
		})
	}
	if value, ok := cu.mutation.Valitiy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coupon.FieldValitiy,
		})
	}
	if value, ok := cu.mutation.AddedValitiy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coupon.FieldValitiy,
		})
	}
	if value, ok := cu.mutation.ActivityID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: coupon.FieldActivityID,
		})
	}
	if value, ok := cu.mutation.AddedActivityID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: coupon.FieldActivityID,
		})
	}
	if cu.mutation.ActivityIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: coupon.FieldActivityID,
		})
	}
	if value, ok := cu.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coupon.FieldRemark,
		})
	}
	if value, ok := cu.mutation.WholeStore(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coupon.FieldWholeStore,
		})
	}
	if value, ok := cu.mutation.AddedWholeStore(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coupon.FieldWholeStore,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coupon.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CouponUpdateOne is the builder for updating a single Coupon entity.
type CouponUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CouponMutation
}

// SetUpdateTime sets the "update_time" field.
func (cuo *CouponUpdateOne) SetUpdateTime(t time.Time) *CouponUpdateOne {
	cuo.mutation.SetUpdateTime(t)
	return cuo
}

// SetDeleteTime sets the "delete_time" field.
func (cuo *CouponUpdateOne) SetDeleteTime(t time.Time) *CouponUpdateOne {
	cuo.mutation.SetDeleteTime(t)
	return cuo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (cuo *CouponUpdateOne) SetNillableDeleteTime(t *time.Time) *CouponUpdateOne {
	if t != nil {
		cuo.SetDeleteTime(*t)
	}
	return cuo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (cuo *CouponUpdateOne) ClearDeleteTime() *CouponUpdateOne {
	cuo.mutation.ClearDeleteTime()
	return cuo
}

// SetTitle sets the "title" field.
func (cuo *CouponUpdateOne) SetTitle(s string) *CouponUpdateOne {
	cuo.mutation.SetTitle(s)
	return cuo
}

// SetStartTime sets the "start_time" field.
func (cuo *CouponUpdateOne) SetStartTime(t time.Time) *CouponUpdateOne {
	cuo.mutation.SetStartTime(t)
	return cuo
}

// SetEndTime sets the "end_time" field.
func (cuo *CouponUpdateOne) SetEndTime(t time.Time) *CouponUpdateOne {
	cuo.mutation.SetEndTime(t)
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *CouponUpdateOne) SetDescription(s string) *CouponUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetFullMoney sets the "full_money" field.
func (cuo *CouponUpdateOne) SetFullMoney(f float64) *CouponUpdateOne {
	cuo.mutation.ResetFullMoney()
	cuo.mutation.SetFullMoney(f)
	return cuo
}

// AddFullMoney adds f to the "full_money" field.
func (cuo *CouponUpdateOne) AddFullMoney(f float64) *CouponUpdateOne {
	cuo.mutation.AddFullMoney(f)
	return cuo
}

// SetMinus sets the "minus" field.
func (cuo *CouponUpdateOne) SetMinus(f float64) *CouponUpdateOne {
	cuo.mutation.ResetMinus()
	cuo.mutation.SetMinus(f)
	return cuo
}

// AddMinus adds f to the "minus" field.
func (cuo *CouponUpdateOne) AddMinus(f float64) *CouponUpdateOne {
	cuo.mutation.AddMinus(f)
	return cuo
}

// SetRate sets the "rate" field.
func (cuo *CouponUpdateOne) SetRate(f float64) *CouponUpdateOne {
	cuo.mutation.ResetRate()
	cuo.mutation.SetRate(f)
	return cuo
}

// AddRate adds f to the "rate" field.
func (cuo *CouponUpdateOne) AddRate(f float64) *CouponUpdateOne {
	cuo.mutation.AddRate(f)
	return cuo
}

// SetType sets the "type" field.
func (cuo *CouponUpdateOne) SetType(i int) *CouponUpdateOne {
	cuo.mutation.ResetType()
	cuo.mutation.SetType(i)
	return cuo
}

// AddType adds i to the "type" field.
func (cuo *CouponUpdateOne) AddType(i int) *CouponUpdateOne {
	cuo.mutation.AddType(i)
	return cuo
}

// SetValitiy sets the "valitiy" field.
func (cuo *CouponUpdateOne) SetValitiy(i int) *CouponUpdateOne {
	cuo.mutation.ResetValitiy()
	cuo.mutation.SetValitiy(i)
	return cuo
}

// AddValitiy adds i to the "valitiy" field.
func (cuo *CouponUpdateOne) AddValitiy(i int) *CouponUpdateOne {
	cuo.mutation.AddValitiy(i)
	return cuo
}

// SetActivityID sets the "activity_id" field.
func (cuo *CouponUpdateOne) SetActivityID(i int64) *CouponUpdateOne {
	cuo.mutation.ResetActivityID()
	cuo.mutation.SetActivityID(i)
	return cuo
}

// SetNillableActivityID sets the "activity_id" field if the given value is not nil.
func (cuo *CouponUpdateOne) SetNillableActivityID(i *int64) *CouponUpdateOne {
	if i != nil {
		cuo.SetActivityID(*i)
	}
	return cuo
}

// AddActivityID adds i to the "activity_id" field.
func (cuo *CouponUpdateOne) AddActivityID(i int64) *CouponUpdateOne {
	cuo.mutation.AddActivityID(i)
	return cuo
}

// ClearActivityID clears the value of the "activity_id" field.
func (cuo *CouponUpdateOne) ClearActivityID() *CouponUpdateOne {
	cuo.mutation.ClearActivityID()
	return cuo
}

// SetRemark sets the "remark" field.
func (cuo *CouponUpdateOne) SetRemark(s string) *CouponUpdateOne {
	cuo.mutation.SetRemark(s)
	return cuo
}

// SetWholeStore sets the "whole_store" field.
func (cuo *CouponUpdateOne) SetWholeStore(i int) *CouponUpdateOne {
	cuo.mutation.ResetWholeStore()
	cuo.mutation.SetWholeStore(i)
	return cuo
}

// AddWholeStore adds i to the "whole_store" field.
func (cuo *CouponUpdateOne) AddWholeStore(i int) *CouponUpdateOne {
	cuo.mutation.AddWholeStore(i)
	return cuo
}

// Mutation returns the CouponMutation object of the builder.
func (cuo *CouponUpdateOne) Mutation() *CouponMutation {
	return cuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CouponUpdateOne) Select(field string, fields ...string) *CouponUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Coupon entity.
func (cuo *CouponUpdateOne) Save(ctx context.Context) (*Coupon, error) {
	var (
		err  error
		node *Coupon
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CouponMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CouponUpdateOne) SaveX(ctx context.Context) *Coupon {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CouponUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CouponUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CouponUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdateTime(); !ok {
		v := coupon.UpdateDefaultUpdateTime()
		cuo.mutation.SetUpdateTime(v)
	}
}

func (cuo *CouponUpdateOne) sqlSave(ctx context.Context) (_node *Coupon, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coupon.Table,
			Columns: coupon.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: coupon.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "Coupon.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coupon.FieldID)
		for _, f := range fields {
			if !coupon.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != coupon.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coupon.FieldUpdateTime,
		})
	}
	if value, ok := cuo.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coupon.FieldDeleteTime,
		})
	}
	if cuo.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: coupon.FieldDeleteTime,
		})
	}
	if value, ok := cuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coupon.FieldTitle,
		})
	}
	if value, ok := cuo.mutation.StartTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coupon.FieldStartTime,
		})
	}
	if value, ok := cuo.mutation.EndTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coupon.FieldEndTime,
		})
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coupon.FieldDescription,
		})
	}
	if value, ok := cuo.mutation.FullMoney(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupon.FieldFullMoney,
		})
	}
	if value, ok := cuo.mutation.AddedFullMoney(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupon.FieldFullMoney,
		})
	}
	if value, ok := cuo.mutation.Minus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupon.FieldMinus,
		})
	}
	if value, ok := cuo.mutation.AddedMinus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupon.FieldMinus,
		})
	}
	if value, ok := cuo.mutation.Rate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupon.FieldRate,
		})
	}
	if value, ok := cuo.mutation.AddedRate(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupon.FieldRate,
		})
	}
	if value, ok := cuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coupon.FieldType,
		})
	}
	if value, ok := cuo.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coupon.FieldType,
		})
	}
	if value, ok := cuo.mutation.Valitiy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coupon.FieldValitiy,
		})
	}
	if value, ok := cuo.mutation.AddedValitiy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coupon.FieldValitiy,
		})
	}
	if value, ok := cuo.mutation.ActivityID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: coupon.FieldActivityID,
		})
	}
	if value, ok := cuo.mutation.AddedActivityID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: coupon.FieldActivityID,
		})
	}
	if cuo.mutation.ActivityIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: coupon.FieldActivityID,
		})
	}
	if value, ok := cuo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coupon.FieldRemark,
		})
	}
	if value, ok := cuo.mutation.WholeStore(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coupon.FieldWholeStore,
		})
	}
	if value, ok := cuo.mutation.AddedWholeStore(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coupon.FieldWholeStore,
		})
	}
	_node = &Coupon{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coupon.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
