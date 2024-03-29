// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/module/order/service/internal/data/model/order"
	"mall-go/module/order/service/internal/data/model/ordersub"
	"mall-go/module/order/service/internal/data/model/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderSubUpdate is the builder for updating OrderSub entities.
type OrderSubUpdate struct {
	config
	hooks    []Hook
	mutation *OrderSubMutation
}

// Where appends a list predicates to the OrderSubUpdate builder.
func (osu *OrderSubUpdate) Where(ps ...predicate.OrderSub) *OrderSubUpdate {
	osu.mutation.Where(ps...)
	return osu
}

// SetUpdateTime sets the "update_time" field.
func (osu *OrderSubUpdate) SetUpdateTime(t time.Time) *OrderSubUpdate {
	osu.mutation.SetUpdateTime(t)
	return osu
}

// SetDeleteTime sets the "delete_time" field.
func (osu *OrderSubUpdate) SetDeleteTime(t time.Time) *OrderSubUpdate {
	osu.mutation.SetDeleteTime(t)
	return osu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (osu *OrderSubUpdate) SetNillableDeleteTime(t *time.Time) *OrderSubUpdate {
	if t != nil {
		osu.SetDeleteTime(*t)
	}
	return osu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (osu *OrderSubUpdate) ClearDeleteTime() *OrderSubUpdate {
	osu.mutation.ClearDeleteTime()
	return osu
}

// SetOrderNo sets the "order_no" field.
func (osu *OrderSubUpdate) SetOrderNo(s string) *OrderSubUpdate {
	osu.mutation.SetOrderNo(s)
	return osu
}

// SetUserID sets the "user_id" field.
func (osu *OrderSubUpdate) SetUserID(i int64) *OrderSubUpdate {
	osu.mutation.ResetUserID()
	osu.mutation.SetUserID(i)
	return osu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (osu *OrderSubUpdate) SetNillableUserID(i *int64) *OrderSubUpdate {
	if i != nil {
		osu.SetUserID(*i)
	}
	return osu
}

// AddUserID adds i to the "user_id" field.
func (osu *OrderSubUpdate) AddUserID(i int64) *OrderSubUpdate {
	osu.mutation.AddUserID(i)
	return osu
}

// ClearUserID clears the value of the "user_id" field.
func (osu *OrderSubUpdate) ClearUserID() *OrderSubUpdate {
	osu.mutation.ClearUserID()
	return osu
}

// SetPrice sets the "price" field.
func (osu *OrderSubUpdate) SetPrice(f float64) *OrderSubUpdate {
	osu.mutation.ResetPrice()
	osu.mutation.SetPrice(f)
	return osu
}

// AddPrice adds f to the "price" field.
func (osu *OrderSubUpdate) AddPrice(f float64) *OrderSubUpdate {
	osu.mutation.AddPrice(f)
	return osu
}

// SetCount sets the "count" field.
func (osu *OrderSubUpdate) SetCount(i int) *OrderSubUpdate {
	osu.mutation.ResetCount()
	osu.mutation.SetCount(i)
	return osu
}

// AddCount adds i to the "count" field.
func (osu *OrderSubUpdate) AddCount(i int) *OrderSubUpdate {
	osu.mutation.AddCount(i)
	return osu
}

// SetFinalPrice sets the "final_price" field.
func (osu *OrderSubUpdate) SetFinalPrice(f float64) *OrderSubUpdate {
	osu.mutation.ResetFinalPrice()
	osu.mutation.SetFinalPrice(f)
	return osu
}

// AddFinalPrice adds f to the "final_price" field.
func (osu *OrderSubUpdate) AddFinalPrice(f float64) *OrderSubUpdate {
	osu.mutation.AddFinalPrice(f)
	return osu
}

// SetStatus sets the "status" field.
func (osu *OrderSubUpdate) SetStatus(i int) *OrderSubUpdate {
	osu.mutation.ResetStatus()
	osu.mutation.SetStatus(i)
	return osu
}

// AddStatus adds i to the "status" field.
func (osu *OrderSubUpdate) AddStatus(i int) *OrderSubUpdate {
	osu.mutation.AddStatus(i)
	return osu
}

// SetOrderID sets the "order_id" field.
func (osu *OrderSubUpdate) SetOrderID(i int64) *OrderSubUpdate {
	osu.mutation.SetOrderID(i)
	return osu
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (osu *OrderSubUpdate) SetNillableOrderID(i *int64) *OrderSubUpdate {
	if i != nil {
		osu.SetOrderID(*i)
	}
	return osu
}

// ClearOrderID clears the value of the "order_id" field.
func (osu *OrderSubUpdate) ClearOrderID() *OrderSubUpdate {
	osu.mutation.ClearOrderID()
	return osu
}

// SetOrder sets the "order" edge to the Order entity.
func (osu *OrderSubUpdate) SetOrder(o *Order) *OrderSubUpdate {
	return osu.SetOrderID(o.ID)
}

// Mutation returns the OrderSubMutation object of the builder.
func (osu *OrderSubUpdate) Mutation() *OrderSubMutation {
	return osu.mutation
}

// ClearOrder clears the "order" edge to the Order entity.
func (osu *OrderSubUpdate) ClearOrder() *OrderSubUpdate {
	osu.mutation.ClearOrder()
	return osu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (osu *OrderSubUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	osu.defaults()
	if len(osu.hooks) == 0 {
		affected, err = osu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderSubMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			osu.mutation = mutation
			affected, err = osu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(osu.hooks) - 1; i >= 0; i-- {
			if osu.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = osu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, osu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (osu *OrderSubUpdate) SaveX(ctx context.Context) int {
	affected, err := osu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (osu *OrderSubUpdate) Exec(ctx context.Context) error {
	_, err := osu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osu *OrderSubUpdate) ExecX(ctx context.Context) {
	if err := osu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (osu *OrderSubUpdate) defaults() {
	if _, ok := osu.mutation.UpdateTime(); !ok {
		v := ordersub.UpdateDefaultUpdateTime()
		osu.mutation.SetUpdateTime(v)
	}
}

func (osu *OrderSubUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ordersub.Table,
			Columns: ordersub.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: ordersub.FieldID,
			},
		},
	}
	if ps := osu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := osu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: ordersub.FieldUpdateTime,
		})
	}
	if value, ok := osu.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: ordersub.FieldDeleteTime,
		})
	}
	if osu.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: ordersub.FieldDeleteTime,
		})
	}
	if value, ok := osu.mutation.OrderNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ordersub.FieldOrderNo,
		})
	}
	if value, ok := osu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: ordersub.FieldUserID,
		})
	}
	if value, ok := osu.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: ordersub.FieldUserID,
		})
	}
	if osu.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: ordersub.FieldUserID,
		})
	}
	if value, ok := osu.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: ordersub.FieldPrice,
		})
	}
	if value, ok := osu.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: ordersub.FieldPrice,
		})
	}
	if value, ok := osu.mutation.Count(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ordersub.FieldCount,
		})
	}
	if value, ok := osu.mutation.AddedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ordersub.FieldCount,
		})
	}
	if value, ok := osu.mutation.FinalPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: ordersub.FieldFinalPrice,
		})
	}
	if value, ok := osu.mutation.AddedFinalPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: ordersub.FieldFinalPrice,
		})
	}
	if value, ok := osu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ordersub.FieldStatus,
		})
	}
	if value, ok := osu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ordersub.FieldStatus,
		})
	}
	if osu.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ordersub.OrderTable,
			Columns: []string{ordersub.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: order.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osu.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ordersub.OrderTable,
			Columns: []string{ordersub.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, osu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ordersub.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// OrderSubUpdateOne is the builder for updating a single OrderSub entity.
type OrderSubUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderSubMutation
}

// SetUpdateTime sets the "update_time" field.
func (osuo *OrderSubUpdateOne) SetUpdateTime(t time.Time) *OrderSubUpdateOne {
	osuo.mutation.SetUpdateTime(t)
	return osuo
}

// SetDeleteTime sets the "delete_time" field.
func (osuo *OrderSubUpdateOne) SetDeleteTime(t time.Time) *OrderSubUpdateOne {
	osuo.mutation.SetDeleteTime(t)
	return osuo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (osuo *OrderSubUpdateOne) SetNillableDeleteTime(t *time.Time) *OrderSubUpdateOne {
	if t != nil {
		osuo.SetDeleteTime(*t)
	}
	return osuo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (osuo *OrderSubUpdateOne) ClearDeleteTime() *OrderSubUpdateOne {
	osuo.mutation.ClearDeleteTime()
	return osuo
}

// SetOrderNo sets the "order_no" field.
func (osuo *OrderSubUpdateOne) SetOrderNo(s string) *OrderSubUpdateOne {
	osuo.mutation.SetOrderNo(s)
	return osuo
}

// SetUserID sets the "user_id" field.
func (osuo *OrderSubUpdateOne) SetUserID(i int64) *OrderSubUpdateOne {
	osuo.mutation.ResetUserID()
	osuo.mutation.SetUserID(i)
	return osuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (osuo *OrderSubUpdateOne) SetNillableUserID(i *int64) *OrderSubUpdateOne {
	if i != nil {
		osuo.SetUserID(*i)
	}
	return osuo
}

// AddUserID adds i to the "user_id" field.
func (osuo *OrderSubUpdateOne) AddUserID(i int64) *OrderSubUpdateOne {
	osuo.mutation.AddUserID(i)
	return osuo
}

// ClearUserID clears the value of the "user_id" field.
func (osuo *OrderSubUpdateOne) ClearUserID() *OrderSubUpdateOne {
	osuo.mutation.ClearUserID()
	return osuo
}

// SetPrice sets the "price" field.
func (osuo *OrderSubUpdateOne) SetPrice(f float64) *OrderSubUpdateOne {
	osuo.mutation.ResetPrice()
	osuo.mutation.SetPrice(f)
	return osuo
}

// AddPrice adds f to the "price" field.
func (osuo *OrderSubUpdateOne) AddPrice(f float64) *OrderSubUpdateOne {
	osuo.mutation.AddPrice(f)
	return osuo
}

// SetCount sets the "count" field.
func (osuo *OrderSubUpdateOne) SetCount(i int) *OrderSubUpdateOne {
	osuo.mutation.ResetCount()
	osuo.mutation.SetCount(i)
	return osuo
}

// AddCount adds i to the "count" field.
func (osuo *OrderSubUpdateOne) AddCount(i int) *OrderSubUpdateOne {
	osuo.mutation.AddCount(i)
	return osuo
}

// SetFinalPrice sets the "final_price" field.
func (osuo *OrderSubUpdateOne) SetFinalPrice(f float64) *OrderSubUpdateOne {
	osuo.mutation.ResetFinalPrice()
	osuo.mutation.SetFinalPrice(f)
	return osuo
}

// AddFinalPrice adds f to the "final_price" field.
func (osuo *OrderSubUpdateOne) AddFinalPrice(f float64) *OrderSubUpdateOne {
	osuo.mutation.AddFinalPrice(f)
	return osuo
}

// SetStatus sets the "status" field.
func (osuo *OrderSubUpdateOne) SetStatus(i int) *OrderSubUpdateOne {
	osuo.mutation.ResetStatus()
	osuo.mutation.SetStatus(i)
	return osuo
}

// AddStatus adds i to the "status" field.
func (osuo *OrderSubUpdateOne) AddStatus(i int) *OrderSubUpdateOne {
	osuo.mutation.AddStatus(i)
	return osuo
}

// SetOrderID sets the "order_id" field.
func (osuo *OrderSubUpdateOne) SetOrderID(i int64) *OrderSubUpdateOne {
	osuo.mutation.SetOrderID(i)
	return osuo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (osuo *OrderSubUpdateOne) SetNillableOrderID(i *int64) *OrderSubUpdateOne {
	if i != nil {
		osuo.SetOrderID(*i)
	}
	return osuo
}

// ClearOrderID clears the value of the "order_id" field.
func (osuo *OrderSubUpdateOne) ClearOrderID() *OrderSubUpdateOne {
	osuo.mutation.ClearOrderID()
	return osuo
}

// SetOrder sets the "order" edge to the Order entity.
func (osuo *OrderSubUpdateOne) SetOrder(o *Order) *OrderSubUpdateOne {
	return osuo.SetOrderID(o.ID)
}

// Mutation returns the OrderSubMutation object of the builder.
func (osuo *OrderSubUpdateOne) Mutation() *OrderSubMutation {
	return osuo.mutation
}

// ClearOrder clears the "order" edge to the Order entity.
func (osuo *OrderSubUpdateOne) ClearOrder() *OrderSubUpdateOne {
	osuo.mutation.ClearOrder()
	return osuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (osuo *OrderSubUpdateOne) Select(field string, fields ...string) *OrderSubUpdateOne {
	osuo.fields = append([]string{field}, fields...)
	return osuo
}

// Save executes the query and returns the updated OrderSub entity.
func (osuo *OrderSubUpdateOne) Save(ctx context.Context) (*OrderSub, error) {
	var (
		err  error
		node *OrderSub
	)
	osuo.defaults()
	if len(osuo.hooks) == 0 {
		node, err = osuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderSubMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			osuo.mutation = mutation
			node, err = osuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(osuo.hooks) - 1; i >= 0; i-- {
			if osuo.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = osuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, osuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*OrderSub)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OrderSubMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (osuo *OrderSubUpdateOne) SaveX(ctx context.Context) *OrderSub {
	node, err := osuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (osuo *OrderSubUpdateOne) Exec(ctx context.Context) error {
	_, err := osuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osuo *OrderSubUpdateOne) ExecX(ctx context.Context) {
	if err := osuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (osuo *OrderSubUpdateOne) defaults() {
	if _, ok := osuo.mutation.UpdateTime(); !ok {
		v := ordersub.UpdateDefaultUpdateTime()
		osuo.mutation.SetUpdateTime(v)
	}
}

func (osuo *OrderSubUpdateOne) sqlSave(ctx context.Context) (_node *OrderSub, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ordersub.Table,
			Columns: ordersub.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: ordersub.FieldID,
			},
		},
	}
	id, ok := osuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "OrderSub.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := osuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ordersub.FieldID)
		for _, f := range fields {
			if !ordersub.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != ordersub.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := osuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := osuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: ordersub.FieldUpdateTime,
		})
	}
	if value, ok := osuo.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: ordersub.FieldDeleteTime,
		})
	}
	if osuo.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: ordersub.FieldDeleteTime,
		})
	}
	if value, ok := osuo.mutation.OrderNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ordersub.FieldOrderNo,
		})
	}
	if value, ok := osuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: ordersub.FieldUserID,
		})
	}
	if value, ok := osuo.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: ordersub.FieldUserID,
		})
	}
	if osuo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: ordersub.FieldUserID,
		})
	}
	if value, ok := osuo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: ordersub.FieldPrice,
		})
	}
	if value, ok := osuo.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: ordersub.FieldPrice,
		})
	}
	if value, ok := osuo.mutation.Count(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ordersub.FieldCount,
		})
	}
	if value, ok := osuo.mutation.AddedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ordersub.FieldCount,
		})
	}
	if value, ok := osuo.mutation.FinalPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: ordersub.FieldFinalPrice,
		})
	}
	if value, ok := osuo.mutation.AddedFinalPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: ordersub.FieldFinalPrice,
		})
	}
	if value, ok := osuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ordersub.FieldStatus,
		})
	}
	if value, ok := osuo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ordersub.FieldStatus,
		})
	}
	if osuo.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ordersub.OrderTable,
			Columns: []string{ordersub.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: order.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osuo.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ordersub.OrderTable,
			Columns: []string{ordersub.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OrderSub{config: osuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, osuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ordersub.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
