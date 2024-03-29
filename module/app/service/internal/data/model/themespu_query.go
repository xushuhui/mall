// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"fmt"
	"mall-go/module/app/service/internal/data/model/predicate"
	"mall-go/module/app/service/internal/data/model/theme"
	"mall-go/module/app/service/internal/data/model/themespu"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ThemeSpuQuery is the builder for querying ThemeSpu entities.
type ThemeSpuQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ThemeSpu
	// eager-loading edges.
	withTheme *ThemeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ThemeSpuQuery builder.
func (tsq *ThemeSpuQuery) Where(ps ...predicate.ThemeSpu) *ThemeSpuQuery {
	tsq.predicates = append(tsq.predicates, ps...)
	return tsq
}

// Limit adds a limit step to the query.
func (tsq *ThemeSpuQuery) Limit(limit int) *ThemeSpuQuery {
	tsq.limit = &limit
	return tsq
}

// Offset adds an offset step to the query.
func (tsq *ThemeSpuQuery) Offset(offset int) *ThemeSpuQuery {
	tsq.offset = &offset
	return tsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tsq *ThemeSpuQuery) Unique(unique bool) *ThemeSpuQuery {
	tsq.unique = &unique
	return tsq
}

// Order adds an order step to the query.
func (tsq *ThemeSpuQuery) Order(o ...OrderFunc) *ThemeSpuQuery {
	tsq.order = append(tsq.order, o...)
	return tsq
}

// QueryTheme chains the current query on the "theme" edge.
func (tsq *ThemeSpuQuery) QueryTheme() *ThemeQuery {
	query := &ThemeQuery{config: tsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(themespu.Table, themespu.FieldID, selector),
			sqlgraph.To(theme.Table, theme.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, themespu.ThemeTable, themespu.ThemeColumn),
		)
		fromU = sqlgraph.SetNeighbors(tsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ThemeSpu entity from the query.
// Returns a *NotFoundError when no ThemeSpu was found.
func (tsq *ThemeSpuQuery) First(ctx context.Context) (*ThemeSpu, error) {
	nodes, err := tsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{themespu.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tsq *ThemeSpuQuery) FirstX(ctx context.Context) *ThemeSpu {
	node, err := tsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ThemeSpu ID from the query.
// Returns a *NotFoundError when no ThemeSpu ID was found.
func (tsq *ThemeSpuQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = tsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{themespu.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tsq *ThemeSpuQuery) FirstIDX(ctx context.Context) int64 {
	id, err := tsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ThemeSpu entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ThemeSpu entity is found.
// Returns a *NotFoundError when no ThemeSpu entities are found.
func (tsq *ThemeSpuQuery) Only(ctx context.Context) (*ThemeSpu, error) {
	nodes, err := tsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{themespu.Label}
	default:
		return nil, &NotSingularError{themespu.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tsq *ThemeSpuQuery) OnlyX(ctx context.Context) *ThemeSpu {
	node, err := tsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ThemeSpu ID in the query.
// Returns a *NotSingularError when more than one ThemeSpu ID is found.
// Returns a *NotFoundError when no entities are found.
func (tsq *ThemeSpuQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = tsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{themespu.Label}
	default:
		err = &NotSingularError{themespu.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tsq *ThemeSpuQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := tsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ThemeSpus.
func (tsq *ThemeSpuQuery) All(ctx context.Context) ([]*ThemeSpu, error) {
	if err := tsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tsq *ThemeSpuQuery) AllX(ctx context.Context) []*ThemeSpu {
	nodes, err := tsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ThemeSpu IDs.
func (tsq *ThemeSpuQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := tsq.Select(themespu.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tsq *ThemeSpuQuery) IDsX(ctx context.Context) []int64 {
	ids, err := tsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tsq *ThemeSpuQuery) Count(ctx context.Context) (int, error) {
	if err := tsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tsq *ThemeSpuQuery) CountX(ctx context.Context) int {
	count, err := tsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tsq *ThemeSpuQuery) Exist(ctx context.Context) (bool, error) {
	if err := tsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tsq *ThemeSpuQuery) ExistX(ctx context.Context) bool {
	exist, err := tsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ThemeSpuQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tsq *ThemeSpuQuery) Clone() *ThemeSpuQuery {
	if tsq == nil {
		return nil
	}
	return &ThemeSpuQuery{
		config:     tsq.config,
		limit:      tsq.limit,
		offset:     tsq.offset,
		order:      append([]OrderFunc{}, tsq.order...),
		predicates: append([]predicate.ThemeSpu{}, tsq.predicates...),
		withTheme:  tsq.withTheme.Clone(),
		// clone intermediate query.
		sql:    tsq.sql.Clone(),
		path:   tsq.path,
		unique: tsq.unique,
	}
}

// WithTheme tells the query-builder to eager-load the nodes that are connected to
// the "theme" edge. The optional arguments are used to configure the query builder of the edge.
func (tsq *ThemeSpuQuery) WithTheme(opts ...func(*ThemeQuery)) *ThemeSpuQuery {
	query := &ThemeQuery{config: tsq.config}
	for _, opt := range opts {
		opt(query)
	}
	tsq.withTheme = query
	return tsq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ThemeID int64 `json:"theme_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ThemeSpu.Query().
//		GroupBy(themespu.FieldThemeID).
//		Aggregate(model.Count()).
//		Scan(ctx, &v)
//
func (tsq *ThemeSpuQuery) GroupBy(field string, fields ...string) *ThemeSpuGroupBy {
	grbuild := &ThemeSpuGroupBy{config: tsq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tsq.sqlQuery(ctx), nil
	}
	grbuild.label = themespu.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ThemeID int64 `json:"theme_id,omitempty"`
//	}
//
//	client.ThemeSpu.Query().
//		Select(themespu.FieldThemeID).
//		Scan(ctx, &v)
//
func (tsq *ThemeSpuQuery) Select(fields ...string) *ThemeSpuSelect {
	tsq.fields = append(tsq.fields, fields...)
	selbuild := &ThemeSpuSelect{ThemeSpuQuery: tsq}
	selbuild.label = themespu.Label
	selbuild.flds, selbuild.scan = &tsq.fields, selbuild.Scan
	return selbuild
}

func (tsq *ThemeSpuQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tsq.fields {
		if !themespu.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
		}
	}
	if tsq.path != nil {
		prev, err := tsq.path(ctx)
		if err != nil {
			return err
		}
		tsq.sql = prev
	}
	return nil
}

func (tsq *ThemeSpuQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ThemeSpu, error) {
	var (
		nodes       = []*ThemeSpu{}
		_spec       = tsq.querySpec()
		loadedTypes = [1]bool{
			tsq.withTheme != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*ThemeSpu).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &ThemeSpu{config: tsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := tsq.withTheme; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*ThemeSpu)
		for i := range nodes {
			fk := nodes[i].ThemeID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(theme.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "theme_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Theme = n
			}
		}
	}

	return nodes, nil
}

func (tsq *ThemeSpuQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tsq.querySpec()
	_spec.Node.Columns = tsq.fields
	if len(tsq.fields) > 0 {
		_spec.Unique = tsq.unique != nil && *tsq.unique
	}
	return sqlgraph.CountNodes(ctx, tsq.driver, _spec)
}

func (tsq *ThemeSpuQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("model: check existence: %w", err)
	}
	return n > 0, nil
}

func (tsq *ThemeSpuQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   themespu.Table,
			Columns: themespu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: themespu.FieldID,
			},
		},
		From:   tsq.sql,
		Unique: true,
	}
	if unique := tsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, themespu.FieldID)
		for i := range fields {
			if fields[i] != themespu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tsq *ThemeSpuQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tsq.driver.Dialect())
	t1 := builder.Table(themespu.Table)
	columns := tsq.fields
	if len(columns) == 0 {
		columns = themespu.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tsq.sql != nil {
		selector = tsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tsq.unique != nil && *tsq.unique {
		selector.Distinct()
	}
	for _, p := range tsq.predicates {
		p(selector)
	}
	for _, p := range tsq.order {
		p(selector)
	}
	if offset := tsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ThemeSpuGroupBy is the group-by builder for ThemeSpu entities.
type ThemeSpuGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tsgb *ThemeSpuGroupBy) Aggregate(fns ...AggregateFunc) *ThemeSpuGroupBy {
	tsgb.fns = append(tsgb.fns, fns...)
	return tsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tsgb *ThemeSpuGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tsgb.path(ctx)
	if err != nil {
		return err
	}
	tsgb.sql = query
	return tsgb.sqlScan(ctx, v)
}

func (tsgb *ThemeSpuGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tsgb.fields {
		if !themespu.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tsgb *ThemeSpuGroupBy) sqlQuery() *sql.Selector {
	selector := tsgb.sql.Select()
	aggregation := make([]string, 0, len(tsgb.fns))
	for _, fn := range tsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tsgb.fields)+len(tsgb.fns))
		for _, f := range tsgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tsgb.fields...)...)
}

// ThemeSpuSelect is the builder for selecting fields of ThemeSpu entities.
type ThemeSpuSelect struct {
	*ThemeSpuQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tss *ThemeSpuSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tss.prepareQuery(ctx); err != nil {
		return err
	}
	tss.sql = tss.ThemeSpuQuery.sqlQuery(ctx)
	return tss.sqlScan(ctx, v)
}

func (tss *ThemeSpuSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tss.sql.Query()
	if err := tss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
