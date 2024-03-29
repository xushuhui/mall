// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/module/order/service/internal/data/model/order"
	"mall-go/module/order/service/internal/data/model/ordersub"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderSubCreate is the builder for creating a OrderSub entity.
type OrderSubCreate struct {
	config
	mutation *OrderSubMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (osc *OrderSubCreate) SetCreateTime(t time.Time) *OrderSubCreate {
	osc.mutation.SetCreateTime(t)
	return osc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (osc *OrderSubCreate) SetNillableCreateTime(t *time.Time) *OrderSubCreate {
	if t != nil {
		osc.SetCreateTime(*t)
	}
	return osc
}

// SetUpdateTime sets the "update_time" field.
func (osc *OrderSubCreate) SetUpdateTime(t time.Time) *OrderSubCreate {
	osc.mutation.SetUpdateTime(t)
	return osc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (osc *OrderSubCreate) SetNillableUpdateTime(t *time.Time) *OrderSubCreate {
	if t != nil {
		osc.SetUpdateTime(*t)
	}
	return osc
}

// SetDeleteTime sets the "delete_time" field.
func (osc *OrderSubCreate) SetDeleteTime(t time.Time) *OrderSubCreate {
	osc.mutation.SetDeleteTime(t)
	return osc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (osc *OrderSubCreate) SetNillableDeleteTime(t *time.Time) *OrderSubCreate {
	if t != nil {
		osc.SetDeleteTime(*t)
	}
	return osc
}

// SetOrderNo sets the "order_no" field.
func (osc *OrderSubCreate) SetOrderNo(s string) *OrderSubCreate {
	osc.mutation.SetOrderNo(s)
	return osc
}

// SetUserID sets the "user_id" field.
func (osc *OrderSubCreate) SetUserID(i int64) *OrderSubCreate {
	osc.mutation.SetUserID(i)
	return osc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (osc *OrderSubCreate) SetNillableUserID(i *int64) *OrderSubCreate {
	if i != nil {
		osc.SetUserID(*i)
	}
	return osc
}

// SetPrice sets the "price" field.
func (osc *OrderSubCreate) SetPrice(f float64) *OrderSubCreate {
	osc.mutation.SetPrice(f)
	return osc
}

// SetCount sets the "count" field.
func (osc *OrderSubCreate) SetCount(i int) *OrderSubCreate {
	osc.mutation.SetCount(i)
	return osc
}

// SetFinalPrice sets the "final_price" field.
func (osc *OrderSubCreate) SetFinalPrice(f float64) *OrderSubCreate {
	osc.mutation.SetFinalPrice(f)
	return osc
}

// SetStatus sets the "status" field.
func (osc *OrderSubCreate) SetStatus(i int) *OrderSubCreate {
	osc.mutation.SetStatus(i)
	return osc
}

// SetOrderID sets the "order_id" field.
func (osc *OrderSubCreate) SetOrderID(i int64) *OrderSubCreate {
	osc.mutation.SetOrderID(i)
	return osc
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (osc *OrderSubCreate) SetNillableOrderID(i *int64) *OrderSubCreate {
	if i != nil {
		osc.SetOrderID(*i)
	}
	return osc
}

// SetOrder sets the "order" edge to the Order entity.
func (osc *OrderSubCreate) SetOrder(o *Order) *OrderSubCreate {
	return osc.SetOrderID(o.ID)
}

// Mutation returns the OrderSubMutation object of the builder.
func (osc *OrderSubCreate) Mutation() *OrderSubMutation {
	return osc.mutation
}

// Save creates the OrderSub in the database.
func (osc *OrderSubCreate) Save(ctx context.Context) (*OrderSub, error) {
	var (
		err  error
		node *OrderSub
	)
	osc.defaults()
	if len(osc.hooks) == 0 {
		if err = osc.check(); err != nil {
			return nil, err
		}
		node, err = osc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderSubMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = osc.check(); err != nil {
				return nil, err
			}
			osc.mutation = mutation
			if node, err = osc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(osc.hooks) - 1; i >= 0; i-- {
			if osc.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = osc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, osc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (osc *OrderSubCreate) SaveX(ctx context.Context) *OrderSub {
	v, err := osc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (osc *OrderSubCreate) Exec(ctx context.Context) error {
	_, err := osc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osc *OrderSubCreate) ExecX(ctx context.Context) {
	if err := osc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (osc *OrderSubCreate) defaults() {
	if _, ok := osc.mutation.CreateTime(); !ok {
		v := ordersub.DefaultCreateTime()
		osc.mutation.SetCreateTime(v)
	}
	if _, ok := osc.mutation.UpdateTime(); !ok {
		v := ordersub.DefaultUpdateTime()
		osc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (osc *OrderSubCreate) check() error {
	if _, ok := osc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "OrderSub.create_time"`)}
	}
	if _, ok := osc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`model: missing required field "OrderSub.update_time"`)}
	}
	if _, ok := osc.mutation.OrderNo(); !ok {
		return &ValidationError{Name: "order_no", err: errors.New(`model: missing required field "OrderSub.order_no"`)}
	}
	if _, ok := osc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`model: missing required field "OrderSub.price"`)}
	}
	if _, ok := osc.mutation.Count(); !ok {
		return &ValidationError{Name: "count", err: errors.New(`model: missing required field "OrderSub.count"`)}
	}
	if _, ok := osc.mutation.FinalPrice(); !ok {
		return &ValidationError{Name: "final_price", err: errors.New(`model: missing required field "OrderSub.final_price"`)}
	}
	if _, ok := osc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`model: missing required field "OrderSub.status"`)}
	}
	return nil
}

func (osc *OrderSubCreate) sqlSave(ctx context.Context) (*OrderSub, error) {
	_node, _spec := osc.createSpec()
	if err := sqlgraph.CreateNode(ctx, osc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (osc *OrderSubCreate) createSpec() (*OrderSub, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderSub{config: osc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: ordersub.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: ordersub.FieldID,
			},
		}
	)
	if value, ok := osc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: ordersub.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := osc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: ordersub.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := osc.mutation.DeleteTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: ordersub.FieldDeleteTime,
		})
		_node.DeleteTime = value
	}
	if value, ok := osc.mutation.OrderNo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ordersub.FieldOrderNo,
		})
		_node.OrderNo = value
	}
	if value, ok := osc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: ordersub.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := osc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: ordersub.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := osc.mutation.Count(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ordersub.FieldCount,
		})
		_node.Count = value
	}
	if value, ok := osc.mutation.FinalPrice(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: ordersub.FieldFinalPrice,
		})
		_node.FinalPrice = value
	}
	if value, ok := osc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ordersub.FieldStatus,
		})
		_node.Status = value
	}
	if nodes := osc.mutation.OrderIDs(); len(nodes) > 0 {
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
		_node.OrderID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderSubCreateBulk is the builder for creating many OrderSub entities in bulk.
type OrderSubCreateBulk struct {
	config
	builders []*OrderSubCreate
}

// Save creates the OrderSub entities in the database.
func (oscb *OrderSubCreateBulk) Save(ctx context.Context) ([]*OrderSub, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oscb.builders))
	nodes := make([]*OrderSub, len(oscb.builders))
	mutators := make([]Mutator, len(oscb.builders))
	for i := range oscb.builders {
		func(i int, root context.Context) {
			builder := oscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderSubMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, oscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, oscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oscb *OrderSubCreateBulk) SaveX(ctx context.Context) []*OrderSub {
	v, err := oscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oscb *OrderSubCreateBulk) Exec(ctx context.Context) error {
	_, err := oscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oscb *OrderSubCreateBulk) ExecX(ctx context.Context) {
	if err := oscb.Exec(ctx); err != nil {
		panic(err)
	}
}
