// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/module/coupon/service/internal/data/model/coupon"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CouponCreate is the builder for creating a Coupon entity.
type CouponCreate struct {
	config
	mutation *CouponMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (cc *CouponCreate) SetCreateTime(t time.Time) *CouponCreate {
	cc.mutation.SetCreateTime(t)
	return cc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (cc *CouponCreate) SetNillableCreateTime(t *time.Time) *CouponCreate {
	if t != nil {
		cc.SetCreateTime(*t)
	}
	return cc
}

// SetUpdateTime sets the "update_time" field.
func (cc *CouponCreate) SetUpdateTime(t time.Time) *CouponCreate {
	cc.mutation.SetUpdateTime(t)
	return cc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (cc *CouponCreate) SetNillableUpdateTime(t *time.Time) *CouponCreate {
	if t != nil {
		cc.SetUpdateTime(*t)
	}
	return cc
}

// SetDeleteTime sets the "delete_time" field.
func (cc *CouponCreate) SetDeleteTime(t time.Time) *CouponCreate {
	cc.mutation.SetDeleteTime(t)
	return cc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (cc *CouponCreate) SetNillableDeleteTime(t *time.Time) *CouponCreate {
	if t != nil {
		cc.SetDeleteTime(*t)
	}
	return cc
}

// SetTitle sets the "title" field.
func (cc *CouponCreate) SetTitle(s string) *CouponCreate {
	cc.mutation.SetTitle(s)
	return cc
}

// SetStartTime sets the "start_time" field.
func (cc *CouponCreate) SetStartTime(t time.Time) *CouponCreate {
	cc.mutation.SetStartTime(t)
	return cc
}

// SetEndTime sets the "end_time" field.
func (cc *CouponCreate) SetEndTime(t time.Time) *CouponCreate {
	cc.mutation.SetEndTime(t)
	return cc
}

// SetDescription sets the "description" field.
func (cc *CouponCreate) SetDescription(s string) *CouponCreate {
	cc.mutation.SetDescription(s)
	return cc
}

// SetFullMoney sets the "full_money" field.
func (cc *CouponCreate) SetFullMoney(f float64) *CouponCreate {
	cc.mutation.SetFullMoney(f)
	return cc
}

// SetMinus sets the "minus" field.
func (cc *CouponCreate) SetMinus(f float64) *CouponCreate {
	cc.mutation.SetMinus(f)
	return cc
}

// SetRate sets the "rate" field.
func (cc *CouponCreate) SetRate(f float64) *CouponCreate {
	cc.mutation.SetRate(f)
	return cc
}

// SetType sets the "type" field.
func (cc *CouponCreate) SetType(i int) *CouponCreate {
	cc.mutation.SetType(i)
	return cc
}

// SetValitiy sets the "valitiy" field.
func (cc *CouponCreate) SetValitiy(i int) *CouponCreate {
	cc.mutation.SetValitiy(i)
	return cc
}

// SetActivityID sets the "activity_id" field.
func (cc *CouponCreate) SetActivityID(i int64) *CouponCreate {
	cc.mutation.SetActivityID(i)
	return cc
}

// SetNillableActivityID sets the "activity_id" field if the given value is not nil.
func (cc *CouponCreate) SetNillableActivityID(i *int64) *CouponCreate {
	if i != nil {
		cc.SetActivityID(*i)
	}
	return cc
}

// SetRemark sets the "remark" field.
func (cc *CouponCreate) SetRemark(s string) *CouponCreate {
	cc.mutation.SetRemark(s)
	return cc
}

// SetWholeStore sets the "whole_store" field.
func (cc *CouponCreate) SetWholeStore(i int) *CouponCreate {
	cc.mutation.SetWholeStore(i)
	return cc
}

// Mutation returns the CouponMutation object of the builder.
func (cc *CouponCreate) Mutation() *CouponMutation {
	return cc.mutation
}

// Save creates the Coupon in the database.
func (cc *CouponCreate) Save(ctx context.Context) (*Coupon, error) {
	var (
		err  error
		node *Coupon
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CouponMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Coupon)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CouponMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CouponCreate) SaveX(ctx context.Context) *Coupon {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CouponCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CouponCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CouponCreate) defaults() {
	if _, ok := cc.mutation.CreateTime(); !ok {
		v := coupon.DefaultCreateTime()
		cc.mutation.SetCreateTime(v)
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		v := coupon.DefaultUpdateTime()
		cc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CouponCreate) check() error {
	if _, ok := cc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "Coupon.create_time"`)}
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`model: missing required field "Coupon.update_time"`)}
	}
	if _, ok := cc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`model: missing required field "Coupon.title"`)}
	}
	if _, ok := cc.mutation.StartTime(); !ok {
		return &ValidationError{Name: "start_time", err: errors.New(`model: missing required field "Coupon.start_time"`)}
	}
	if _, ok := cc.mutation.EndTime(); !ok {
		return &ValidationError{Name: "end_time", err: errors.New(`model: missing required field "Coupon.end_time"`)}
	}
	if _, ok := cc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`model: missing required field "Coupon.description"`)}
	}
	if _, ok := cc.mutation.FullMoney(); !ok {
		return &ValidationError{Name: "full_money", err: errors.New(`model: missing required field "Coupon.full_money"`)}
	}
	if _, ok := cc.mutation.Minus(); !ok {
		return &ValidationError{Name: "minus", err: errors.New(`model: missing required field "Coupon.minus"`)}
	}
	if _, ok := cc.mutation.Rate(); !ok {
		return &ValidationError{Name: "rate", err: errors.New(`model: missing required field "Coupon.rate"`)}
	}
	if _, ok := cc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`model: missing required field "Coupon.type"`)}
	}
	if _, ok := cc.mutation.Valitiy(); !ok {
		return &ValidationError{Name: "valitiy", err: errors.New(`model: missing required field "Coupon.valitiy"`)}
	}
	if _, ok := cc.mutation.Remark(); !ok {
		return &ValidationError{Name: "remark", err: errors.New(`model: missing required field "Coupon.remark"`)}
	}
	if _, ok := cc.mutation.WholeStore(); !ok {
		return &ValidationError{Name: "whole_store", err: errors.New(`model: missing required field "Coupon.whole_store"`)}
	}
	return nil
}

func (cc *CouponCreate) sqlSave(ctx context.Context) (*Coupon, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (cc *CouponCreate) createSpec() (*Coupon, *sqlgraph.CreateSpec) {
	var (
		_node = &Coupon{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: coupon.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: coupon.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coupon.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := cc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coupon.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := cc.mutation.DeleteTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coupon.FieldDeleteTime,
		})
		_node.DeleteTime = value
	}
	if value, ok := cc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coupon.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := cc.mutation.StartTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coupon.FieldStartTime,
		})
		_node.StartTime = value
	}
	if value, ok := cc.mutation.EndTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coupon.FieldEndTime,
		})
		_node.EndTime = value
	}
	if value, ok := cc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coupon.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := cc.mutation.FullMoney(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupon.FieldFullMoney,
		})
		_node.FullMoney = value
	}
	if value, ok := cc.mutation.Minus(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupon.FieldMinus,
		})
		_node.Minus = value
	}
	if value, ok := cc.mutation.Rate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupon.FieldRate,
		})
		_node.Rate = value
	}
	if value, ok := cc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coupon.FieldType,
		})
		_node.Type = value
	}
	if value, ok := cc.mutation.Valitiy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coupon.FieldValitiy,
		})
		_node.Valitiy = value
	}
	if value, ok := cc.mutation.ActivityID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: coupon.FieldActivityID,
		})
		_node.ActivityID = value
	}
	if value, ok := cc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coupon.FieldRemark,
		})
		_node.Remark = value
	}
	if value, ok := cc.mutation.WholeStore(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coupon.FieldWholeStore,
		})
		_node.WholeStore = value
	}
	return _node, _spec
}

// CouponCreateBulk is the builder for creating many Coupon entities in bulk.
type CouponCreateBulk struct {
	config
	builders []*CouponCreate
}

// Save creates the Coupon entities in the database.
func (ccb *CouponCreateBulk) Save(ctx context.Context) ([]*Coupon, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Coupon, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CouponMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CouponCreateBulk) SaveX(ctx context.Context) []*Coupon {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CouponCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CouponCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
