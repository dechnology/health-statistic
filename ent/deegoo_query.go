// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/eesoymilk/health-statistic-api/ent/deegoo"
	"github.com/eesoymilk/health-statistic-api/ent/predicate"
	"github.com/eesoymilk/health-statistic-api/ent/user"
	"github.com/google/uuid"
)

// DeegooQuery is the builder for querying Deegoo entities.
type DeegooQuery struct {
	config
	ctx        *QueryContext
	order      []deegoo.OrderOption
	inters     []Interceptor
	predicates []predicate.Deegoo
	withUser   *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DeegooQuery builder.
func (dq *DeegooQuery) Where(ps ...predicate.Deegoo) *DeegooQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DeegooQuery) Limit(limit int) *DeegooQuery {
	dq.ctx.Limit = &limit
	return dq
}

// Offset to start from.
func (dq *DeegooQuery) Offset(offset int) *DeegooQuery {
	dq.ctx.Offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DeegooQuery) Unique(unique bool) *DeegooQuery {
	dq.ctx.Unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DeegooQuery) Order(o ...deegoo.OrderOption) *DeegooQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryUser chains the current query on the "user" edge.
func (dq *DeegooQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(deegoo.Table, deegoo.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, deegoo.UserTable, deegoo.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Deegoo entity from the query.
// Returns a *NotFoundError when no Deegoo was found.
func (dq *DeegooQuery) First(ctx context.Context) (*Deegoo, error) {
	nodes, err := dq.Limit(1).All(setContextOp(ctx, dq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{deegoo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DeegooQuery) FirstX(ctx context.Context) *Deegoo {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Deegoo ID from the query.
// Returns a *NotFoundError when no Deegoo ID was found.
func (dq *DeegooQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(1).IDs(setContextOp(ctx, dq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{deegoo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DeegooQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Deegoo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Deegoo entity is found.
// Returns a *NotFoundError when no Deegoo entities are found.
func (dq *DeegooQuery) Only(ctx context.Context) (*Deegoo, error) {
	nodes, err := dq.Limit(2).All(setContextOp(ctx, dq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{deegoo.Label}
	default:
		return nil, &NotSingularError{deegoo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DeegooQuery) OnlyX(ctx context.Context) *Deegoo {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Deegoo ID in the query.
// Returns a *NotSingularError when more than one Deegoo ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DeegooQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(2).IDs(setContextOp(ctx, dq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{deegoo.Label}
	default:
		err = &NotSingularError{deegoo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DeegooQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Deegoos.
func (dq *DeegooQuery) All(ctx context.Context) ([]*Deegoo, error) {
	ctx = setContextOp(ctx, dq.ctx, "All")
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Deegoo, *DeegooQuery]()
	return withInterceptors[[]*Deegoo](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DeegooQuery) AllX(ctx context.Context) []*Deegoo {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Deegoo IDs.
func (dq *DeegooQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if dq.ctx.Unique == nil && dq.path != nil {
		dq.Unique(true)
	}
	ctx = setContextOp(ctx, dq.ctx, "IDs")
	if err = dq.Select(deegoo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DeegooQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DeegooQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dq.ctx, "Count")
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DeegooQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DeegooQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DeegooQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dq.ctx, "Exist")
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DeegooQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DeegooQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DeegooQuery) Clone() *DeegooQuery {
	if dq == nil {
		return nil
	}
	return &DeegooQuery{
		config:     dq.config,
		ctx:        dq.ctx.Clone(),
		order:      append([]deegoo.OrderOption{}, dq.order...),
		inters:     append([]Interceptor{}, dq.inters...),
		predicates: append([]predicate.Deegoo{}, dq.predicates...),
		withUser:   dq.withUser.Clone(),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DeegooQuery) WithUser(opts ...func(*UserQuery)) *DeegooQuery {
	query := (&UserClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withUser = query
	return dq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Perception int8 `json:"perception"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Deegoo.Query().
//		GroupBy(deegoo.FieldPerception).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DeegooQuery) GroupBy(field string, fields ...string) *DeegooGroupBy {
	dq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DeegooGroupBy{build: dq}
	grbuild.flds = &dq.ctx.Fields
	grbuild.label = deegoo.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Perception int8 `json:"perception"`
//	}
//
//	client.Deegoo.Query().
//		Select(deegoo.FieldPerception).
//		Scan(ctx, &v)
func (dq *DeegooQuery) Select(fields ...string) *DeegooSelect {
	dq.ctx.Fields = append(dq.ctx.Fields, fields...)
	sbuild := &DeegooSelect{DeegooQuery: dq}
	sbuild.label = deegoo.Label
	sbuild.flds, sbuild.scan = &dq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DeegooSelect configured with the given aggregations.
func (dq *DeegooQuery) Aggregate(fns ...AggregateFunc) *DeegooSelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DeegooQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dq); err != nil {
				return err
			}
		}
	}
	for _, f := range dq.ctx.Fields {
		if !deegoo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DeegooQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Deegoo, error) {
	var (
		nodes       = []*Deegoo{}
		withFKs     = dq.withFKs
		_spec       = dq.querySpec()
		loadedTypes = [1]bool{
			dq.withUser != nil,
		}
	)
	if dq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, deegoo.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Deegoo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Deegoo{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dq.withUser; query != nil {
		if err := dq.loadUser(ctx, query, nodes, nil,
			func(n *Deegoo, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DeegooQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Deegoo, init func(*Deegoo), assign func(*Deegoo, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Deegoo)
	for i := range nodes {
		if nodes[i].user_deegoo == nil {
			continue
		}
		fk := *nodes[i].user_deegoo
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_deegoo" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (dq *DeegooQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Columns = dq.ctx.Fields
	if len(dq.ctx.Fields) > 0 {
		_spec.Unique = dq.ctx.Unique != nil && *dq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DeegooQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(deegoo.Table, deegoo.Columns, sqlgraph.NewFieldSpec(deegoo.FieldID, field.TypeUUID))
	_spec.From = dq.sql
	if unique := dq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dq.path != nil {
		_spec.Unique = true
	}
	if fields := dq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deegoo.FieldID)
		for i := range fields {
			if fields[i] != deegoo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DeegooQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(deegoo.Table)
	columns := dq.ctx.Fields
	if len(columns) == 0 {
		columns = deegoo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.ctx.Unique != nil && *dq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DeegooGroupBy is the group-by builder for Deegoo entities.
type DeegooGroupBy struct {
	selector
	build *DeegooQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DeegooGroupBy) Aggregate(fns ...AggregateFunc) *DeegooGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DeegooGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dgb.build.ctx, "GroupBy")
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DeegooQuery, *DeegooGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DeegooGroupBy) sqlScan(ctx context.Context, root *DeegooQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dgb.flds)+len(dgb.fns))
		for _, f := range *dgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DeegooSelect is the builder for selecting fields of Deegoo entities.
type DeegooSelect struct {
	*DeegooQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DeegooSelect) Aggregate(fns ...AggregateFunc) *DeegooSelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DeegooSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ds.ctx, "Select")
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DeegooQuery, *DeegooSelect](ctx, ds.DeegooQuery, ds, ds.inters, v)
}

func (ds *DeegooSelect) sqlScan(ctx context.Context, root *DeegooQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
