// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"fmt"
	"mall-go/internal/data/model/category"
	"mall-go/internal/data/model/coupon"
	"mall-go/internal/data/model/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CategoryUpdate is the builder for updating Category entities.
type CategoryUpdate struct {
	config
	hooks    []Hook
	mutation *CategoryMutation
}

// Where appends a list predicates to the CategoryUpdate builder.
func (cu *CategoryUpdate) Where(ps ...predicate.Category) *CategoryUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdateTime sets the "update_time" field.
func (cu *CategoryUpdate) SetUpdateTime(t time.Time) *CategoryUpdate {
	cu.mutation.SetUpdateTime(t)
	return cu
}

// SetDeleteTime sets the "delete_time" field.
func (cu *CategoryUpdate) SetDeleteTime(t time.Time) *CategoryUpdate {
	cu.mutation.SetDeleteTime(t)
	return cu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableDeleteTime(t *time.Time) *CategoryUpdate {
	if t != nil {
		cu.SetDeleteTime(*t)
	}
	return cu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (cu *CategoryUpdate) ClearDeleteTime() *CategoryUpdate {
	cu.mutation.ClearDeleteTime()
	return cu
}

// SetName sets the "name" field.
func (cu *CategoryUpdate) SetName(s string) *CategoryUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetDescription sets the "description" field.
func (cu *CategoryUpdate) SetDescription(s string) *CategoryUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetIsRoot sets the "is_root" field.
func (cu *CategoryUpdate) SetIsRoot(i int) *CategoryUpdate {
	cu.mutation.ResetIsRoot()
	cu.mutation.SetIsRoot(i)
	return cu
}

// AddIsRoot adds i to the "is_root" field.
func (cu *CategoryUpdate) AddIsRoot(i int) *CategoryUpdate {
	cu.mutation.AddIsRoot(i)
	return cu
}

// SetParentID sets the "parent_id" field.
func (cu *CategoryUpdate) SetParentID(i int64) *CategoryUpdate {
	cu.mutation.SetParentID(i)
	return cu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableParentID(i *int64) *CategoryUpdate {
	if i != nil {
		cu.SetParentID(*i)
	}
	return cu
}

// ClearParentID clears the value of the "parent_id" field.
func (cu *CategoryUpdate) ClearParentID() *CategoryUpdate {
	cu.mutation.ClearParentID()
	return cu
}

// SetImg sets the "img" field.
func (cu *CategoryUpdate) SetImg(s string) *CategoryUpdate {
	cu.mutation.SetImg(s)
	return cu
}

// SetIndex sets the "index" field.
func (cu *CategoryUpdate) SetIndex(i int) *CategoryUpdate {
	cu.mutation.ResetIndex()
	cu.mutation.SetIndex(i)
	return cu
}

// AddIndex adds i to the "index" field.
func (cu *CategoryUpdate) AddIndex(i int) *CategoryUpdate {
	cu.mutation.AddIndex(i)
	return cu
}

// SetOnline sets the "online" field.
func (cu *CategoryUpdate) SetOnline(i int) *CategoryUpdate {
	cu.mutation.ResetOnline()
	cu.mutation.SetOnline(i)
	return cu
}

// AddOnline adds i to the "online" field.
func (cu *CategoryUpdate) AddOnline(i int) *CategoryUpdate {
	cu.mutation.AddOnline(i)
	return cu
}

// SetLevel sets the "level" field.
func (cu *CategoryUpdate) SetLevel(i int) *CategoryUpdate {
	cu.mutation.ResetLevel()
	cu.mutation.SetLevel(i)
	return cu
}

// AddLevel adds i to the "level" field.
func (cu *CategoryUpdate) AddLevel(i int) *CategoryUpdate {
	cu.mutation.AddLevel(i)
	return cu
}

// AddCouponIDs adds the "coupon" edge to the Coupon entity by IDs.
func (cu *CategoryUpdate) AddCouponIDs(ids ...int64) *CategoryUpdate {
	cu.mutation.AddCouponIDs(ids...)
	return cu
}

// AddCoupon adds the "coupon" edges to the Coupon entity.
func (cu *CategoryUpdate) AddCoupon(c ...*Coupon) *CategoryUpdate {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddCouponIDs(ids...)
}

// SetParent sets the "parent" edge to the Category entity.
func (cu *CategoryUpdate) SetParent(c *Category) *CategoryUpdate {
	return cu.SetParentID(c.ID)
}

// AddChildIDs adds the "children" edge to the Category entity by IDs.
func (cu *CategoryUpdate) AddChildIDs(ids ...int64) *CategoryUpdate {
	cu.mutation.AddChildIDs(ids...)
	return cu
}

// AddChildren adds the "children" edges to the Category entity.
func (cu *CategoryUpdate) AddChildren(c ...*Category) *CategoryUpdate {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddChildIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cu *CategoryUpdate) Mutation() *CategoryMutation {
	return cu.mutation
}

// ClearCoupon clears all "coupon" edges to the Coupon entity.
func (cu *CategoryUpdate) ClearCoupon() *CategoryUpdate {
	cu.mutation.ClearCoupon()
	return cu
}

// RemoveCouponIDs removes the "coupon" edge to Coupon entities by IDs.
func (cu *CategoryUpdate) RemoveCouponIDs(ids ...int64) *CategoryUpdate {
	cu.mutation.RemoveCouponIDs(ids...)
	return cu
}

// RemoveCoupon removes "coupon" edges to Coupon entities.
func (cu *CategoryUpdate) RemoveCoupon(c ...*Coupon) *CategoryUpdate {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveCouponIDs(ids...)
}

// ClearParent clears the "parent" edge to the Category entity.
func (cu *CategoryUpdate) ClearParent() *CategoryUpdate {
	cu.mutation.ClearParent()
	return cu
}

// ClearChildren clears all "children" edges to the Category entity.
func (cu *CategoryUpdate) ClearChildren() *CategoryUpdate {
	cu.mutation.ClearChildren()
	return cu
}

// RemoveChildIDs removes the "children" edge to Category entities by IDs.
func (cu *CategoryUpdate) RemoveChildIDs(ids ...int64) *CategoryUpdate {
	cu.mutation.RemoveChildIDs(ids...)
	return cu
}

// RemoveChildren removes "children" edges to Category entities.
func (cu *CategoryUpdate) RemoveChildren(c ...*Category) *CategoryUpdate {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveChildIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CategoryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CategoryMutation)
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
func (cu *CategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CategoryUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CategoryUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CategoryUpdate) defaults() {
	if _, ok := cu.mutation.UpdateTime(); !ok {
		v := category.UpdateDefaultUpdateTime()
		cu.mutation.SetUpdateTime(v)
	}
}

func (cu *CategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   category.Table,
			Columns: category.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: category.FieldID,
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
			Column: category.FieldUpdateTime,
		})
	}
	if value, ok := cu.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: category.FieldDeleteTime,
		})
	}
	if cu.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: category.FieldDeleteTime,
		})
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldName,
		})
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldDescription,
		})
	}
	if value, ok := cu.mutation.IsRoot(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: category.FieldIsRoot,
		})
	}
	if value, ok := cu.mutation.AddedIsRoot(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: category.FieldIsRoot,
		})
	}
	if value, ok := cu.mutation.Img(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldImg,
		})
	}
	if value, ok := cu.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: category.FieldIndex,
		})
	}
	if value, ok := cu.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: category.FieldIndex,
		})
	}
	if value, ok := cu.mutation.Online(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: category.FieldOnline,
		})
	}
	if value, ok := cu.mutation.AddedOnline(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: category.FieldOnline,
		})
	}
	if value, ok := cu.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: category.FieldLevel,
		})
	}
	if value, ok := cu.mutation.AddedLevel(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: category.FieldLevel,
		})
	}
	if cu.mutation.CouponCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.CouponTable,
			Columns: category.CouponPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: coupon.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedCouponIDs(); len(nodes) > 0 && !cu.mutation.CouponCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.CouponTable,
			Columns: category.CouponPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: coupon.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CouponIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.CouponTable,
			Columns: category.CouponPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: coupon.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   category.ParentTable,
			Columns: []string{category.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   category.ParentTable,
			Columns: []string{category.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.ChildrenTable,
			Columns: []string{category.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !cu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.ChildrenTable,
			Columns: []string{category.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.ChildrenTable,
			Columns: []string{category.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CategoryUpdateOne is the builder for updating a single Category entity.
type CategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CategoryMutation
}

// SetUpdateTime sets the "update_time" field.
func (cuo *CategoryUpdateOne) SetUpdateTime(t time.Time) *CategoryUpdateOne {
	cuo.mutation.SetUpdateTime(t)
	return cuo
}

// SetDeleteTime sets the "delete_time" field.
func (cuo *CategoryUpdateOne) SetDeleteTime(t time.Time) *CategoryUpdateOne {
	cuo.mutation.SetDeleteTime(t)
	return cuo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableDeleteTime(t *time.Time) *CategoryUpdateOne {
	if t != nil {
		cuo.SetDeleteTime(*t)
	}
	return cuo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (cuo *CategoryUpdateOne) ClearDeleteTime() *CategoryUpdateOne {
	cuo.mutation.ClearDeleteTime()
	return cuo
}

// SetName sets the "name" field.
func (cuo *CategoryUpdateOne) SetName(s string) *CategoryUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *CategoryUpdateOne) SetDescription(s string) *CategoryUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetIsRoot sets the "is_root" field.
func (cuo *CategoryUpdateOne) SetIsRoot(i int) *CategoryUpdateOne {
	cuo.mutation.ResetIsRoot()
	cuo.mutation.SetIsRoot(i)
	return cuo
}

// AddIsRoot adds i to the "is_root" field.
func (cuo *CategoryUpdateOne) AddIsRoot(i int) *CategoryUpdateOne {
	cuo.mutation.AddIsRoot(i)
	return cuo
}

// SetParentID sets the "parent_id" field.
func (cuo *CategoryUpdateOne) SetParentID(i int64) *CategoryUpdateOne {
	cuo.mutation.SetParentID(i)
	return cuo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableParentID(i *int64) *CategoryUpdateOne {
	if i != nil {
		cuo.SetParentID(*i)
	}
	return cuo
}

// ClearParentID clears the value of the "parent_id" field.
func (cuo *CategoryUpdateOne) ClearParentID() *CategoryUpdateOne {
	cuo.mutation.ClearParentID()
	return cuo
}

// SetImg sets the "img" field.
func (cuo *CategoryUpdateOne) SetImg(s string) *CategoryUpdateOne {
	cuo.mutation.SetImg(s)
	return cuo
}

// SetIndex sets the "index" field.
func (cuo *CategoryUpdateOne) SetIndex(i int) *CategoryUpdateOne {
	cuo.mutation.ResetIndex()
	cuo.mutation.SetIndex(i)
	return cuo
}

// AddIndex adds i to the "index" field.
func (cuo *CategoryUpdateOne) AddIndex(i int) *CategoryUpdateOne {
	cuo.mutation.AddIndex(i)
	return cuo
}

// SetOnline sets the "online" field.
func (cuo *CategoryUpdateOne) SetOnline(i int) *CategoryUpdateOne {
	cuo.mutation.ResetOnline()
	cuo.mutation.SetOnline(i)
	return cuo
}

// AddOnline adds i to the "online" field.
func (cuo *CategoryUpdateOne) AddOnline(i int) *CategoryUpdateOne {
	cuo.mutation.AddOnline(i)
	return cuo
}

// SetLevel sets the "level" field.
func (cuo *CategoryUpdateOne) SetLevel(i int) *CategoryUpdateOne {
	cuo.mutation.ResetLevel()
	cuo.mutation.SetLevel(i)
	return cuo
}

// AddLevel adds i to the "level" field.
func (cuo *CategoryUpdateOne) AddLevel(i int) *CategoryUpdateOne {
	cuo.mutation.AddLevel(i)
	return cuo
}

// AddCouponIDs adds the "coupon" edge to the Coupon entity by IDs.
func (cuo *CategoryUpdateOne) AddCouponIDs(ids ...int64) *CategoryUpdateOne {
	cuo.mutation.AddCouponIDs(ids...)
	return cuo
}

// AddCoupon adds the "coupon" edges to the Coupon entity.
func (cuo *CategoryUpdateOne) AddCoupon(c ...*Coupon) *CategoryUpdateOne {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddCouponIDs(ids...)
}

// SetParent sets the "parent" edge to the Category entity.
func (cuo *CategoryUpdateOne) SetParent(c *Category) *CategoryUpdateOne {
	return cuo.SetParentID(c.ID)
}

// AddChildIDs adds the "children" edge to the Category entity by IDs.
func (cuo *CategoryUpdateOne) AddChildIDs(ids ...int64) *CategoryUpdateOne {
	cuo.mutation.AddChildIDs(ids...)
	return cuo
}

// AddChildren adds the "children" edges to the Category entity.
func (cuo *CategoryUpdateOne) AddChildren(c ...*Category) *CategoryUpdateOne {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddChildIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cuo *CategoryUpdateOne) Mutation() *CategoryMutation {
	return cuo.mutation
}

// ClearCoupon clears all "coupon" edges to the Coupon entity.
func (cuo *CategoryUpdateOne) ClearCoupon() *CategoryUpdateOne {
	cuo.mutation.ClearCoupon()
	return cuo
}

// RemoveCouponIDs removes the "coupon" edge to Coupon entities by IDs.
func (cuo *CategoryUpdateOne) RemoveCouponIDs(ids ...int64) *CategoryUpdateOne {
	cuo.mutation.RemoveCouponIDs(ids...)
	return cuo
}

// RemoveCoupon removes "coupon" edges to Coupon entities.
func (cuo *CategoryUpdateOne) RemoveCoupon(c ...*Coupon) *CategoryUpdateOne {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveCouponIDs(ids...)
}

// ClearParent clears the "parent" edge to the Category entity.
func (cuo *CategoryUpdateOne) ClearParent() *CategoryUpdateOne {
	cuo.mutation.ClearParent()
	return cuo
}

// ClearChildren clears all "children" edges to the Category entity.
func (cuo *CategoryUpdateOne) ClearChildren() *CategoryUpdateOne {
	cuo.mutation.ClearChildren()
	return cuo
}

// RemoveChildIDs removes the "children" edge to Category entities by IDs.
func (cuo *CategoryUpdateOne) RemoveChildIDs(ids ...int64) *CategoryUpdateOne {
	cuo.mutation.RemoveChildIDs(ids...)
	return cuo
}

// RemoveChildren removes "children" edges to Category entities.
func (cuo *CategoryUpdateOne) RemoveChildren(c ...*Category) *CategoryUpdateOne {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveChildIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CategoryUpdateOne) Select(field string, fields ...string) *CategoryUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Category entity.
func (cuo *CategoryUpdateOne) Save(ctx context.Context) (*Category, error) {
	var (
		err  error
		node *Category
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CategoryMutation)
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
func (cuo *CategoryUpdateOne) SaveX(ctx context.Context) *Category {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CategoryUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CategoryUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdateTime(); !ok {
		v := category.UpdateDefaultUpdateTime()
		cuo.mutation.SetUpdateTime(v)
	}
}

func (cuo *CategoryUpdateOne) sqlSave(ctx context.Context) (_node *Category, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   category.Table,
			Columns: category.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: category.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Category.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, category.FieldID)
		for _, f := range fields {
			if !category.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != category.FieldID {
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
			Column: category.FieldUpdateTime,
		})
	}
	if value, ok := cuo.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: category.FieldDeleteTime,
		})
	}
	if cuo.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: category.FieldDeleteTime,
		})
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldName,
		})
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldDescription,
		})
	}
	if value, ok := cuo.mutation.IsRoot(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: category.FieldIsRoot,
		})
	}
	if value, ok := cuo.mutation.AddedIsRoot(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: category.FieldIsRoot,
		})
	}
	if value, ok := cuo.mutation.Img(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldImg,
		})
	}
	if value, ok := cuo.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: category.FieldIndex,
		})
	}
	if value, ok := cuo.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: category.FieldIndex,
		})
	}
	if value, ok := cuo.mutation.Online(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: category.FieldOnline,
		})
	}
	if value, ok := cuo.mutation.AddedOnline(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: category.FieldOnline,
		})
	}
	if value, ok := cuo.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: category.FieldLevel,
		})
	}
	if value, ok := cuo.mutation.AddedLevel(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: category.FieldLevel,
		})
	}
	if cuo.mutation.CouponCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.CouponTable,
			Columns: category.CouponPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: coupon.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedCouponIDs(); len(nodes) > 0 && !cuo.mutation.CouponCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.CouponTable,
			Columns: category.CouponPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: coupon.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CouponIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.CouponTable,
			Columns: category.CouponPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: coupon.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   category.ParentTable,
			Columns: []string{category.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   category.ParentTable,
			Columns: []string{category.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.ChildrenTable,
			Columns: []string{category.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !cuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.ChildrenTable,
			Columns: []string{category.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.ChildrenTable,
			Columns: []string{category.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Category{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
