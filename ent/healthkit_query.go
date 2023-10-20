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
	"github.com/eesoymilk/health-statistic-api/ent/healthkit"
	"github.com/eesoymilk/health-statistic-api/ent/predicate"
	"github.com/eesoymilk/health-statistic-api/ent/user"
	"github.com/google/uuid"
)

// HealthKitQuery is the builder for querying HealthKit entities.
type HealthKitQuery struct {
	config
	ctx        *QueryContext
	order      []healthkit.OrderOption
	inters     []Interceptor
	predicates []predicate.HealthKit
	withUser   *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the HealthKitQuery builder.
func (hkq *HealthKitQuery) Where(ps ...predicate.HealthKit) *HealthKitQuery {
	hkq.predicates = append(hkq.predicates, ps...)
	return hkq
}

// Limit the number of records to be returned by this query.
func (hkq *HealthKitQuery) Limit(limit int) *HealthKitQuery {
	hkq.ctx.Limit = &limit
	return hkq
}

// Offset to start from.
func (hkq *HealthKitQuery) Offset(offset int) *HealthKitQuery {
	hkq.ctx.Offset = &offset
	return hkq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (hkq *HealthKitQuery) Unique(unique bool) *HealthKitQuery {
	hkq.ctx.Unique = &unique
	return hkq
}

// Order specifies how the records should be ordered.
func (hkq *HealthKitQuery) Order(o ...healthkit.OrderOption) *HealthKitQuery {
	hkq.order = append(hkq.order, o...)
	return hkq
}

// QueryUser chains the current query on the "user" edge.
func (hkq *HealthKitQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: hkq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hkq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hkq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(healthkit.Table, healthkit.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, healthkit.UserTable, healthkit.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(hkq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first HealthKit entity from the query.
// Returns a *NotFoundError when no HealthKit was found.
func (hkq *HealthKitQuery) First(ctx context.Context) (*HealthKit, error) {
	nodes, err := hkq.Limit(1).All(setContextOp(ctx, hkq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{healthkit.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (hkq *HealthKitQuery) FirstX(ctx context.Context) *HealthKit {
	node, err := hkq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first HealthKit ID from the query.
// Returns a *NotFoundError when no HealthKit ID was found.
func (hkq *HealthKitQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = hkq.Limit(1).IDs(setContextOp(ctx, hkq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{healthkit.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (hkq *HealthKitQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := hkq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single HealthKit entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one HealthKit entity is found.
// Returns a *NotFoundError when no HealthKit entities are found.
func (hkq *HealthKitQuery) Only(ctx context.Context) (*HealthKit, error) {
	nodes, err := hkq.Limit(2).All(setContextOp(ctx, hkq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{healthkit.Label}
	default:
		return nil, &NotSingularError{healthkit.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (hkq *HealthKitQuery) OnlyX(ctx context.Context) *HealthKit {
	node, err := hkq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only HealthKit ID in the query.
// Returns a *NotSingularError when more than one HealthKit ID is found.
// Returns a *NotFoundError when no entities are found.
func (hkq *HealthKitQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = hkq.Limit(2).IDs(setContextOp(ctx, hkq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{healthkit.Label}
	default:
		err = &NotSingularError{healthkit.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (hkq *HealthKitQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := hkq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of HealthKits.
func (hkq *HealthKitQuery) All(ctx context.Context) ([]*HealthKit, error) {
	ctx = setContextOp(ctx, hkq.ctx, "All")
	if err := hkq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*HealthKit, *HealthKitQuery]()
	return withInterceptors[[]*HealthKit](ctx, hkq, qr, hkq.inters)
}

// AllX is like All, but panics if an error occurs.
func (hkq *HealthKitQuery) AllX(ctx context.Context) []*HealthKit {
	nodes, err := hkq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of HealthKit IDs.
func (hkq *HealthKitQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if hkq.ctx.Unique == nil && hkq.path != nil {
		hkq.Unique(true)
	}
	ctx = setContextOp(ctx, hkq.ctx, "IDs")
	if err = hkq.Select(healthkit.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (hkq *HealthKitQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := hkq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (hkq *HealthKitQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, hkq.ctx, "Count")
	if err := hkq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, hkq, querierCount[*HealthKitQuery](), hkq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (hkq *HealthKitQuery) CountX(ctx context.Context) int {
	count, err := hkq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (hkq *HealthKitQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, hkq.ctx, "Exist")
	switch _, err := hkq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (hkq *HealthKitQuery) ExistX(ctx context.Context) bool {
	exist, err := hkq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the HealthKitQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (hkq *HealthKitQuery) Clone() *HealthKitQuery {
	if hkq == nil {
		return nil
	}
	return &HealthKitQuery{
		config:     hkq.config,
		ctx:        hkq.ctx.Clone(),
		order:      append([]healthkit.OrderOption{}, hkq.order...),
		inters:     append([]Interceptor{}, hkq.inters...),
		predicates: append([]predicate.HealthKit{}, hkq.predicates...),
		withUser:   hkq.withUser.Clone(),
		// clone intermediate query.
		sql:  hkq.sql.Clone(),
		path: hkq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (hkq *HealthKitQuery) WithUser(opts ...func(*UserQuery)) *HealthKitQuery {
	query := (&UserClient{config: hkq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	hkq.withUser = query
	return hkq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		StartDate time.Time `json:"start_date,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.HealthKit.Query().
//		GroupBy(healthkit.FieldStartDate).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (hkq *HealthKitQuery) GroupBy(field string, fields ...string) *HealthKitGroupBy {
	hkq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &HealthKitGroupBy{build: hkq}
	grbuild.flds = &hkq.ctx.Fields
	grbuild.label = healthkit.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		StartDate time.Time `json:"start_date,omitempty"`
//	}
//
//	client.HealthKit.Query().
//		Select(healthkit.FieldStartDate).
//		Scan(ctx, &v)
func (hkq *HealthKitQuery) Select(fields ...string) *HealthKitSelect {
	hkq.ctx.Fields = append(hkq.ctx.Fields, fields...)
	sbuild := &HealthKitSelect{HealthKitQuery: hkq}
	sbuild.label = healthkit.Label
	sbuild.flds, sbuild.scan = &hkq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a HealthKitSelect configured with the given aggregations.
func (hkq *HealthKitQuery) Aggregate(fns ...AggregateFunc) *HealthKitSelect {
	return hkq.Select().Aggregate(fns...)
}

func (hkq *HealthKitQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range hkq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, hkq); err != nil {
				return err
			}
		}
	}
	for _, f := range hkq.ctx.Fields {
		if !healthkit.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if hkq.path != nil {
		prev, err := hkq.path(ctx)
		if err != nil {
			return err
		}
		hkq.sql = prev
	}
	return nil
}

func (hkq *HealthKitQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*HealthKit, error) {
	var (
		nodes       = []*HealthKit{}
		withFKs     = hkq.withFKs
		_spec       = hkq.querySpec()
		loadedTypes = [1]bool{
			hkq.withUser != nil,
		}
	)
	if hkq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, healthkit.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*HealthKit).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &HealthKit{config: hkq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, hkq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := hkq.withUser; query != nil {
		if err := hkq.loadUser(ctx, query, nodes, nil,
			func(n *HealthKit, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (hkq *HealthKitQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*HealthKit, init func(*HealthKit), assign func(*HealthKit, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*HealthKit)
	for i := range nodes {
		if nodes[i].user_healthkit == nil {
			continue
		}
		fk := *nodes[i].user_healthkit
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
			return fmt.Errorf(`unexpected foreign-key "user_healthkit" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (hkq *HealthKitQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := hkq.querySpec()
	_spec.Node.Columns = hkq.ctx.Fields
	if len(hkq.ctx.Fields) > 0 {
		_spec.Unique = hkq.ctx.Unique != nil && *hkq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, hkq.driver, _spec)
}

func (hkq *HealthKitQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(healthkit.Table, healthkit.Columns, sqlgraph.NewFieldSpec(healthkit.FieldID, field.TypeUUID))
	_spec.From = hkq.sql
	if unique := hkq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if hkq.path != nil {
		_spec.Unique = true
	}
	if fields := hkq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, healthkit.FieldID)
		for i := range fields {
			if fields[i] != healthkit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := hkq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := hkq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := hkq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := hkq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (hkq *HealthKitQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(hkq.driver.Dialect())
	t1 := builder.Table(healthkit.Table)
	columns := hkq.ctx.Fields
	if len(columns) == 0 {
		columns = healthkit.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if hkq.sql != nil {
		selector = hkq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if hkq.ctx.Unique != nil && *hkq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range hkq.predicates {
		p(selector)
	}
	for _, p := range hkq.order {
		p(selector)
	}
	if offset := hkq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := hkq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// HealthKitGroupBy is the group-by builder for HealthKit entities.
type HealthKitGroupBy struct {
	selector
	build *HealthKitQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (hkgb *HealthKitGroupBy) Aggregate(fns ...AggregateFunc) *HealthKitGroupBy {
	hkgb.fns = append(hkgb.fns, fns...)
	return hkgb
}

// Scan applies the selector query and scans the result into the given value.
func (hkgb *HealthKitGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hkgb.build.ctx, "GroupBy")
	if err := hkgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HealthKitQuery, *HealthKitGroupBy](ctx, hkgb.build, hkgb, hkgb.build.inters, v)
}

func (hkgb *HealthKitGroupBy) sqlScan(ctx context.Context, root *HealthKitQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(hkgb.fns))
	for _, fn := range hkgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*hkgb.flds)+len(hkgb.fns))
		for _, f := range *hkgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*hkgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hkgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// HealthKitSelect is the builder for selecting fields of HealthKit entities.
type HealthKitSelect struct {
	*HealthKitQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (hks *HealthKitSelect) Aggregate(fns ...AggregateFunc) *HealthKitSelect {
	hks.fns = append(hks.fns, fns...)
	return hks
}

// Scan applies the selector query and scans the result into the given value.
func (hks *HealthKitSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hks.ctx, "Select")
	if err := hks.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HealthKitQuery, *HealthKitSelect](ctx, hks.HealthKitQuery, hks, hks.inters, v)
}

func (hks *HealthKitSelect) sqlScan(ctx context.Context, root *HealthKitQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(hks.fns))
	for _, fn := range hks.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*hks.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hks.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
