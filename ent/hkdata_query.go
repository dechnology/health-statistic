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
	"github.com/eesoymilk/health-statistic-api/ent/hkdata"
	"github.com/eesoymilk/health-statistic-api/ent/predicate"
	"github.com/google/uuid"
)

// HKDataQuery is the builder for querying HKData entities.
type HKDataQuery struct {
	config
	ctx           *QueryContext
	order         []hkdata.OrderOption
	inters        []Interceptor
	predicates    []predicate.HKData
	withHealthkit *HealthKitQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the HKDataQuery builder.
func (hdq *HKDataQuery) Where(ps ...predicate.HKData) *HKDataQuery {
	hdq.predicates = append(hdq.predicates, ps...)
	return hdq
}

// Limit the number of records to be returned by this query.
func (hdq *HKDataQuery) Limit(limit int) *HKDataQuery {
	hdq.ctx.Limit = &limit
	return hdq
}

// Offset to start from.
func (hdq *HKDataQuery) Offset(offset int) *HKDataQuery {
	hdq.ctx.Offset = &offset
	return hdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (hdq *HKDataQuery) Unique(unique bool) *HKDataQuery {
	hdq.ctx.Unique = &unique
	return hdq
}

// Order specifies how the records should be ordered.
func (hdq *HKDataQuery) Order(o ...hkdata.OrderOption) *HKDataQuery {
	hdq.order = append(hdq.order, o...)
	return hdq
}

// QueryHealthkit chains the current query on the "healthkit" edge.
func (hdq *HKDataQuery) QueryHealthkit() *HealthKitQuery {
	query := (&HealthKitClient{config: hdq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(hkdata.Table, hkdata.FieldID, selector),
			sqlgraph.To(healthkit.Table, healthkit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, hkdata.HealthkitTable, hkdata.HealthkitColumn),
		)
		fromU = sqlgraph.SetNeighbors(hdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first HKData entity from the query.
// Returns a *NotFoundError when no HKData was found.
func (hdq *HKDataQuery) First(ctx context.Context) (*HKData, error) {
	nodes, err := hdq.Limit(1).All(setContextOp(ctx, hdq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{hkdata.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (hdq *HKDataQuery) FirstX(ctx context.Context) *HKData {
	node, err := hdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first HKData ID from the query.
// Returns a *NotFoundError when no HKData ID was found.
func (hdq *HKDataQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = hdq.Limit(1).IDs(setContextOp(ctx, hdq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{hkdata.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (hdq *HKDataQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := hdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single HKData entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one HKData entity is found.
// Returns a *NotFoundError when no HKData entities are found.
func (hdq *HKDataQuery) Only(ctx context.Context) (*HKData, error) {
	nodes, err := hdq.Limit(2).All(setContextOp(ctx, hdq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{hkdata.Label}
	default:
		return nil, &NotSingularError{hkdata.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (hdq *HKDataQuery) OnlyX(ctx context.Context) *HKData {
	node, err := hdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only HKData ID in the query.
// Returns a *NotSingularError when more than one HKData ID is found.
// Returns a *NotFoundError when no entities are found.
func (hdq *HKDataQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = hdq.Limit(2).IDs(setContextOp(ctx, hdq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{hkdata.Label}
	default:
		err = &NotSingularError{hkdata.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (hdq *HKDataQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := hdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of HKDataSlice.
func (hdq *HKDataQuery) All(ctx context.Context) ([]*HKData, error) {
	ctx = setContextOp(ctx, hdq.ctx, "All")
	if err := hdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*HKData, *HKDataQuery]()
	return withInterceptors[[]*HKData](ctx, hdq, qr, hdq.inters)
}

// AllX is like All, but panics if an error occurs.
func (hdq *HKDataQuery) AllX(ctx context.Context) []*HKData {
	nodes, err := hdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of HKData IDs.
func (hdq *HKDataQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if hdq.ctx.Unique == nil && hdq.path != nil {
		hdq.Unique(true)
	}
	ctx = setContextOp(ctx, hdq.ctx, "IDs")
	if err = hdq.Select(hkdata.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (hdq *HKDataQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := hdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (hdq *HKDataQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, hdq.ctx, "Count")
	if err := hdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, hdq, querierCount[*HKDataQuery](), hdq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (hdq *HKDataQuery) CountX(ctx context.Context) int {
	count, err := hdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (hdq *HKDataQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, hdq.ctx, "Exist")
	switch _, err := hdq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (hdq *HKDataQuery) ExistX(ctx context.Context) bool {
	exist, err := hdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the HKDataQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (hdq *HKDataQuery) Clone() *HKDataQuery {
	if hdq == nil {
		return nil
	}
	return &HKDataQuery{
		config:        hdq.config,
		ctx:           hdq.ctx.Clone(),
		order:         append([]hkdata.OrderOption{}, hdq.order...),
		inters:        append([]Interceptor{}, hdq.inters...),
		predicates:    append([]predicate.HKData{}, hdq.predicates...),
		withHealthkit: hdq.withHealthkit.Clone(),
		// clone intermediate query.
		sql:  hdq.sql.Clone(),
		path: hdq.path,
	}
}

// WithHealthkit tells the query-builder to eager-load the nodes that are connected to
// the "healthkit" edge. The optional arguments are used to configure the query builder of the edge.
func (hdq *HKDataQuery) WithHealthkit(opts ...func(*HealthKitQuery)) *HKDataQuery {
	query := (&HealthKitClient{config: hdq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	hdq.withHealthkit = query
	return hdq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DataID string `json:"data_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.HKData.Query().
//		GroupBy(hkdata.FieldDataID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (hdq *HKDataQuery) GroupBy(field string, fields ...string) *HKDataGroupBy {
	hdq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &HKDataGroupBy{build: hdq}
	grbuild.flds = &hdq.ctx.Fields
	grbuild.label = hkdata.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DataID string `json:"data_id,omitempty"`
//	}
//
//	client.HKData.Query().
//		Select(hkdata.FieldDataID).
//		Scan(ctx, &v)
func (hdq *HKDataQuery) Select(fields ...string) *HKDataSelect {
	hdq.ctx.Fields = append(hdq.ctx.Fields, fields...)
	sbuild := &HKDataSelect{HKDataQuery: hdq}
	sbuild.label = hkdata.Label
	sbuild.flds, sbuild.scan = &hdq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a HKDataSelect configured with the given aggregations.
func (hdq *HKDataQuery) Aggregate(fns ...AggregateFunc) *HKDataSelect {
	return hdq.Select().Aggregate(fns...)
}

func (hdq *HKDataQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range hdq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, hdq); err != nil {
				return err
			}
		}
	}
	for _, f := range hdq.ctx.Fields {
		if !hkdata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if hdq.path != nil {
		prev, err := hdq.path(ctx)
		if err != nil {
			return err
		}
		hdq.sql = prev
	}
	return nil
}

func (hdq *HKDataQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*HKData, error) {
	var (
		nodes       = []*HKData{}
		withFKs     = hdq.withFKs
		_spec       = hdq.querySpec()
		loadedTypes = [1]bool{
			hdq.withHealthkit != nil,
		}
	)
	if hdq.withHealthkit != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, hkdata.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*HKData).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &HKData{config: hdq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, hdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := hdq.withHealthkit; query != nil {
		if err := hdq.loadHealthkit(ctx, query, nodes, nil,
			func(n *HKData, e *HealthKit) { n.Edges.Healthkit = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (hdq *HKDataQuery) loadHealthkit(ctx context.Context, query *HealthKitQuery, nodes []*HKData, init func(*HKData), assign func(*HKData, *HealthKit)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*HKData)
	for i := range nodes {
		if nodes[i].health_kit_data == nil {
			continue
		}
		fk := *nodes[i].health_kit_data
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(healthkit.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "health_kit_data" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (hdq *HKDataQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := hdq.querySpec()
	_spec.Node.Columns = hdq.ctx.Fields
	if len(hdq.ctx.Fields) > 0 {
		_spec.Unique = hdq.ctx.Unique != nil && *hdq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, hdq.driver, _spec)
}

func (hdq *HKDataQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(hkdata.Table, hkdata.Columns, sqlgraph.NewFieldSpec(hkdata.FieldID, field.TypeUUID))
	_spec.From = hdq.sql
	if unique := hdq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if hdq.path != nil {
		_spec.Unique = true
	}
	if fields := hdq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hkdata.FieldID)
		for i := range fields {
			if fields[i] != hkdata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := hdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := hdq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := hdq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := hdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (hdq *HKDataQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(hdq.driver.Dialect())
	t1 := builder.Table(hkdata.Table)
	columns := hdq.ctx.Fields
	if len(columns) == 0 {
		columns = hkdata.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if hdq.sql != nil {
		selector = hdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if hdq.ctx.Unique != nil && *hdq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range hdq.predicates {
		p(selector)
	}
	for _, p := range hdq.order {
		p(selector)
	}
	if offset := hdq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := hdq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// HKDataGroupBy is the group-by builder for HKData entities.
type HKDataGroupBy struct {
	selector
	build *HKDataQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (hdgb *HKDataGroupBy) Aggregate(fns ...AggregateFunc) *HKDataGroupBy {
	hdgb.fns = append(hdgb.fns, fns...)
	return hdgb
}

// Scan applies the selector query and scans the result into the given value.
func (hdgb *HKDataGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hdgb.build.ctx, "GroupBy")
	if err := hdgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HKDataQuery, *HKDataGroupBy](ctx, hdgb.build, hdgb, hdgb.build.inters, v)
}

func (hdgb *HKDataGroupBy) sqlScan(ctx context.Context, root *HKDataQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(hdgb.fns))
	for _, fn := range hdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*hdgb.flds)+len(hdgb.fns))
		for _, f := range *hdgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*hdgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hdgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// HKDataSelect is the builder for selecting fields of HKData entities.
type HKDataSelect struct {
	*HKDataQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (hds *HKDataSelect) Aggregate(fns ...AggregateFunc) *HKDataSelect {
	hds.fns = append(hds.fns, fns...)
	return hds
}

// Scan applies the selector query and scans the result into the given value.
func (hds *HKDataSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hds.ctx, "Select")
	if err := hds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HKDataQuery, *HKDataSelect](ctx, hds.HKDataQuery, hds, hds.inters, v)
}

func (hds *HKDataSelect) sqlScan(ctx context.Context, root *HKDataQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(hds.fns))
	for _, fn := range hds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*hds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
