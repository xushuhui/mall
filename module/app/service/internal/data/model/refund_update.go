// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/module/app/service/internal/data/model/predicate"
	"mall-go/module/app/service/internal/data/model/refund"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RefundUpdate is the builder for updating Refund entities.
type RefundUpdate struct {
	config
	hooks    []Hook
	mutation *RefundMutation
}

// Where appends a list predicates to the RefundUpdate builder.
func (ru *RefundUpdate) Where(ps ...predicate.Refund) *RefundUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUpdateTime sets the "update_time" field.
func (ru *RefundUpdate) SetUpdateTime(t time.Time) *RefundUpdate {
	ru.mutation.SetUpdateTime(t)
	return ru
}

// SetDeleteTime sets the "delete_time" field.
func (ru *RefundUpdate) SetDeleteTime(t time.Time) *RefundUpdate {
	ru.mutation.SetDeleteTime(t)
	return ru
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ru *RefundUpdate) SetNillableDeleteTime(t *time.Time) *RefundUpdate {
	if t != nil {
		ru.SetDeleteTime(*t)
	}
	return ru
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (ru *RefundUpdate) ClearDeleteTime() *RefundUpdate {
	ru.mutation.ClearDeleteTime()
	return ru
}

// SetRefundNo sets the "refund_no" field.
func (ru *RefundUpdate) SetRefundNo(s string) *RefundUpdate {
	ru.mutation.SetRefundNo(s)
	return ru
}

// SetTransactionID sets the "transaction_id" field.
func (ru *RefundUpdate) SetTransactionID(s string) *RefundUpdate {
	ru.mutation.SetTransactionID(s)
	return ru
}

// SetUserID sets the "user_id" field.
func (ru *RefundUpdate) SetUserID(i int64) *RefundUpdate {
	ru.mutation.ResetUserID()
	ru.mutation.SetUserID(i)
	return ru
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ru *RefundUpdate) SetNillableUserID(i *int64) *RefundUpdate {
	if i != nil {
		ru.SetUserID(*i)
	}
	return ru
}

// AddUserID adds i to the "user_id" field.
func (ru *RefundUpdate) AddUserID(i int64) *RefundUpdate {
	ru.mutation.AddUserID(i)
	return ru
}

// ClearUserID clears the value of the "user_id" field.
func (ru *RefundUpdate) ClearUserID() *RefundUpdate {
	ru.mutation.ClearUserID()
	return ru
}

// SetReason sets the "reason" field.
func (ru *RefundUpdate) SetReason(s string) *RefundUpdate {
	ru.mutation.SetReason(s)
	return ru
}

// SetOrderID sets the "order_id" field.
func (ru *RefundUpdate) SetOrderID(i int64) *RefundUpdate {
	ru.mutation.ResetOrderID()
	ru.mutation.SetOrderID(i)
	return ru
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (ru *RefundUpdate) SetNillableOrderID(i *int64) *RefundUpdate {
	if i != nil {
		ru.SetOrderID(*i)
	}
	return ru
}

// AddOrderID adds i to the "order_id" field.
func (ru *RefundUpdate) AddOrderID(i int64) *RefundUpdate {
	ru.mutation.AddOrderID(i)
	return ru
}

// ClearOrderID clears the value of the "order_id" field.
func (ru *RefundUpdate) ClearOrderID() *RefundUpdate {
	ru.mutation.ClearOrderID()
	return ru
}

// SetOrderSubID sets the "order_sub_id" field.
func (ru *RefundUpdate) SetOrderSubID(i int64) *RefundUpdate {
	ru.mutation.ResetOrderSubID()
	ru.mutation.SetOrderSubID(i)
	return ru
}

// SetNillableOrderSubID sets the "order_sub_id" field if the given value is not nil.
func (ru *RefundUpdate) SetNillableOrderSubID(i *int64) *RefundUpdate {
	if i != nil {
		ru.SetOrderSubID(*i)
	}
	return ru
}

// AddOrderSubID adds i to the "order_sub_id" field.
func (ru *RefundUpdate) AddOrderSubID(i int64) *RefundUpdate {
	ru.mutation.AddOrderSubID(i)
	return ru
}

// ClearOrderSubID clears the value of the "order_sub_id" field.
func (ru *RefundUpdate) ClearOrderSubID() *RefundUpdate {
	ru.mutation.ClearOrderSubID()
	return ru
}

// SetStatus sets the "status" field.
func (ru *RefundUpdate) SetStatus(i int) *RefundUpdate {
	ru.mutation.ResetStatus()
	ru.mutation.SetStatus(i)
	return ru
}

// AddStatus adds i to the "status" field.
func (ru *RefundUpdate) AddStatus(i int) *RefundUpdate {
	ru.mutation.AddStatus(i)
	return ru
}

// Mutation returns the RefundMutation object of the builder.
func (ru *RefundUpdate) Mutation() *RefundMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RefundUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ru.defaults()
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RefundMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RefundUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RefundUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RefundUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RefundUpdate) defaults() {
	if _, ok := ru.mutation.UpdateTime(); !ok {
		v := refund.UpdateDefaultUpdateTime()
		ru.mutation.SetUpdateTime(v)
	}
}

func (ru *RefundUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   refund.Table,
			Columns: refund.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: refund.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: refund.FieldUpdateTime,
		})
	}
	if value, ok := ru.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: refund.FieldDeleteTime,
		})
	}
	if ru.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: refund.FieldDeleteTime,
		})
	}
	if value, ok := ru.mutation.RefundNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: refund.FieldRefundNo,
		})
	}
	if value, ok := ru.mutation.TransactionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: refund.FieldTransactionID,
		})
	}
	if value, ok := ru.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: refund.FieldUserID,
		})
	}
	if value, ok := ru.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: refund.FieldUserID,
		})
	}
	if ru.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: refund.FieldUserID,
		})
	}
	if value, ok := ru.mutation.Reason(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: refund.FieldReason,
		})
	}
	if value, ok := ru.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: refund.FieldOrderID,
		})
	}
	if value, ok := ru.mutation.AddedOrderID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: refund.FieldOrderID,
		})
	}
	if ru.mutation.OrderIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: refund.FieldOrderID,
		})
	}
	if value, ok := ru.mutation.OrderSubID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: refund.FieldOrderSubID,
		})
	}
	if value, ok := ru.mutation.AddedOrderSubID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: refund.FieldOrderSubID,
		})
	}
	if ru.mutation.OrderSubIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: refund.FieldOrderSubID,
		})
	}
	if value, ok := ru.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: refund.FieldStatus,
		})
	}
	if value, ok := ru.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: refund.FieldStatus,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{refund.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// RefundUpdateOne is the builder for updating a single Refund entity.
type RefundUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RefundMutation
}

// SetUpdateTime sets the "update_time" field.
func (ruo *RefundUpdateOne) SetUpdateTime(t time.Time) *RefundUpdateOne {
	ruo.mutation.SetUpdateTime(t)
	return ruo
}

// SetDeleteTime sets the "delete_time" field.
func (ruo *RefundUpdateOne) SetDeleteTime(t time.Time) *RefundUpdateOne {
	ruo.mutation.SetDeleteTime(t)
	return ruo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ruo *RefundUpdateOne) SetNillableDeleteTime(t *time.Time) *RefundUpdateOne {
	if t != nil {
		ruo.SetDeleteTime(*t)
	}
	return ruo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (ruo *RefundUpdateOne) ClearDeleteTime() *RefundUpdateOne {
	ruo.mutation.ClearDeleteTime()
	return ruo
}

// SetRefundNo sets the "refund_no" field.
func (ruo *RefundUpdateOne) SetRefundNo(s string) *RefundUpdateOne {
	ruo.mutation.SetRefundNo(s)
	return ruo
}

// SetTransactionID sets the "transaction_id" field.
func (ruo *RefundUpdateOne) SetTransactionID(s string) *RefundUpdateOne {
	ruo.mutation.SetTransactionID(s)
	return ruo
}

// SetUserID sets the "user_id" field.
func (ruo *RefundUpdateOne) SetUserID(i int64) *RefundUpdateOne {
	ruo.mutation.ResetUserID()
	ruo.mutation.SetUserID(i)
	return ruo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ruo *RefundUpdateOne) SetNillableUserID(i *int64) *RefundUpdateOne {
	if i != nil {
		ruo.SetUserID(*i)
	}
	return ruo
}

// AddUserID adds i to the "user_id" field.
func (ruo *RefundUpdateOne) AddUserID(i int64) *RefundUpdateOne {
	ruo.mutation.AddUserID(i)
	return ruo
}

// ClearUserID clears the value of the "user_id" field.
func (ruo *RefundUpdateOne) ClearUserID() *RefundUpdateOne {
	ruo.mutation.ClearUserID()
	return ruo
}

// SetReason sets the "reason" field.
func (ruo *RefundUpdateOne) SetReason(s string) *RefundUpdateOne {
	ruo.mutation.SetReason(s)
	return ruo
}

// SetOrderID sets the "order_id" field.
func (ruo *RefundUpdateOne) SetOrderID(i int64) *RefundUpdateOne {
	ruo.mutation.ResetOrderID()
	ruo.mutation.SetOrderID(i)
	return ruo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (ruo *RefundUpdateOne) SetNillableOrderID(i *int64) *RefundUpdateOne {
	if i != nil {
		ruo.SetOrderID(*i)
	}
	return ruo
}

// AddOrderID adds i to the "order_id" field.
func (ruo *RefundUpdateOne) AddOrderID(i int64) *RefundUpdateOne {
	ruo.mutation.AddOrderID(i)
	return ruo
}

// ClearOrderID clears the value of the "order_id" field.
func (ruo *RefundUpdateOne) ClearOrderID() *RefundUpdateOne {
	ruo.mutation.ClearOrderID()
	return ruo
}

// SetOrderSubID sets the "order_sub_id" field.
func (ruo *RefundUpdateOne) SetOrderSubID(i int64) *RefundUpdateOne {
	ruo.mutation.ResetOrderSubID()
	ruo.mutation.SetOrderSubID(i)
	return ruo
}

// SetNillableOrderSubID sets the "order_sub_id" field if the given value is not nil.
func (ruo *RefundUpdateOne) SetNillableOrderSubID(i *int64) *RefundUpdateOne {
	if i != nil {
		ruo.SetOrderSubID(*i)
	}
	return ruo
}

// AddOrderSubID adds i to the "order_sub_id" field.
func (ruo *RefundUpdateOne) AddOrderSubID(i int64) *RefundUpdateOne {
	ruo.mutation.AddOrderSubID(i)
	return ruo
}

// ClearOrderSubID clears the value of the "order_sub_id" field.
func (ruo *RefundUpdateOne) ClearOrderSubID() *RefundUpdateOne {
	ruo.mutation.ClearOrderSubID()
	return ruo
}

// SetStatus sets the "status" field.
func (ruo *RefundUpdateOne) SetStatus(i int) *RefundUpdateOne {
	ruo.mutation.ResetStatus()
	ruo.mutation.SetStatus(i)
	return ruo
}

// AddStatus adds i to the "status" field.
func (ruo *RefundUpdateOne) AddStatus(i int) *RefundUpdateOne {
	ruo.mutation.AddStatus(i)
	return ruo
}

// Mutation returns the RefundMutation object of the builder.
func (ruo *RefundUpdateOne) Mutation() *RefundMutation {
	return ruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RefundUpdateOne) Select(field string, fields ...string) *RefundUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Refund entity.
func (ruo *RefundUpdateOne) Save(ctx context.Context) (*Refund, error) {
	var (
		err  error
		node *Refund
	)
	ruo.defaults()
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RefundMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ruo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Refund)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RefundMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RefundUpdateOne) SaveX(ctx context.Context) *Refund {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RefundUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RefundUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RefundUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdateTime(); !ok {
		v := refund.UpdateDefaultUpdateTime()
		ruo.mutation.SetUpdateTime(v)
	}
}

func (ruo *RefundUpdateOne) sqlSave(ctx context.Context) (_node *Refund, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   refund.Table,
			Columns: refund.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: refund.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "Refund.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, refund.FieldID)
		for _, f := range fields {
			if !refund.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != refund.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: refund.FieldUpdateTime,
		})
	}
	if value, ok := ruo.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: refund.FieldDeleteTime,
		})
	}
	if ruo.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: refund.FieldDeleteTime,
		})
	}
	if value, ok := ruo.mutation.RefundNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: refund.FieldRefundNo,
		})
	}
	if value, ok := ruo.mutation.TransactionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: refund.FieldTransactionID,
		})
	}
	if value, ok := ruo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: refund.FieldUserID,
		})
	}
	if value, ok := ruo.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: refund.FieldUserID,
		})
	}
	if ruo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: refund.FieldUserID,
		})
	}
	if value, ok := ruo.mutation.Reason(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: refund.FieldReason,
		})
	}
	if value, ok := ruo.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: refund.FieldOrderID,
		})
	}
	if value, ok := ruo.mutation.AddedOrderID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: refund.FieldOrderID,
		})
	}
	if ruo.mutation.OrderIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: refund.FieldOrderID,
		})
	}
	if value, ok := ruo.mutation.OrderSubID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: refund.FieldOrderSubID,
		})
	}
	if value, ok := ruo.mutation.AddedOrderSubID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: refund.FieldOrderSubID,
		})
	}
	if ruo.mutation.OrderSubIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: refund.FieldOrderSubID,
		})
	}
	if value, ok := ruo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: refund.FieldStatus,
		})
	}
	if value, ok := ruo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: refund.FieldStatus,
		})
	}
	_node = &Refund{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{refund.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
