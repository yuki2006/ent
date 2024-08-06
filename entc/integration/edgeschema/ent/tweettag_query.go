// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/entc/integration/edgeschema/ent/tag"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweet"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweettag"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TweetTagQuery is the builder for querying TweetTag entities.
type TweetTagQuery struct {
	config
	ctx        *QueryContext
	order      []tweettag.OrderOption
	inters     []Interceptor
	predicates []predicate.TweetTag
	withTag    *TagQuery
	withTweet  *TweetQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TweetTagQuery builder.
func (ttq *TweetTagQuery) Where(ps ...predicate.TweetTag) *TweetTagQuery {
	ttq.predicates = append(ttq.predicates, ps...)
	return ttq
}

// Limit the number of records to be returned by this query.
func (ttq *TweetTagQuery) Limit(limit int) *TweetTagQuery {
	ttq.ctx.Limit = &limit
	return ttq
}

// Offset to start from.
func (ttq *TweetTagQuery) Offset(offset int) *TweetTagQuery {
	ttq.ctx.Offset = &offset
	return ttq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ttq *TweetTagQuery) Unique(unique bool) *TweetTagQuery {
	ttq.ctx.Unique = &unique
	return ttq
}

// Order specifies how the records should be ordered.
func (ttq *TweetTagQuery) Order(o ...tweettag.OrderOption) *TweetTagQuery {
	ttq.order = append(ttq.order, o...)
	return ttq
}

// QueryTag chains the current query on the "tag" edge.
func (ttq *TweetTagQuery) QueryTag() *TagQuery {
	query := (&TagClient{config: ttq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ttq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tweettag.Table, tweettag.FieldID, selector),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, tweettag.TagTable, tweettag.TagColumn),
		)
		fromU = sqlgraph.SetNeighbors(ttq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTweet chains the current query on the "tweet" edge.
func (ttq *TweetTagQuery) QueryTweet() *TweetQuery {
	query := (&TweetClient{config: ttq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ttq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tweettag.Table, tweettag.FieldID, selector),
			sqlgraph.To(tweet.Table, tweet.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, tweettag.TweetTable, tweettag.TweetColumn),
		)
		fromU = sqlgraph.SetNeighbors(ttq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TweetTag entity from the query.
// Returns a *NotFoundError when no TweetTag was found.
func (ttq *TweetTagQuery) First(ctx context.Context) (*TweetTag, error) {
	nodes, err := ttq.Limit(1).All(setContextOp(ctx, ttq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tweettag.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ttq *TweetTagQuery) FirstX(ctx context.Context) *TweetTag {
	node, err := ttq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TweetTag ID from the query.
// Returns a *NotFoundError when no TweetTag ID was found.
func (ttq *TweetTagQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ttq.Limit(1).IDs(setContextOp(ctx, ttq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tweettag.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ttq *TweetTagQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ttq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TweetTag entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TweetTag entity is found.
// Returns a *NotFoundError when no TweetTag entities are found.
func (ttq *TweetTagQuery) Only(ctx context.Context) (*TweetTag, error) {
	nodes, err := ttq.Limit(2).All(setContextOp(ctx, ttq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tweettag.Label}
	default:
		return nil, &NotSingularError{tweettag.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ttq *TweetTagQuery) OnlyX(ctx context.Context) *TweetTag {
	node, err := ttq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TweetTag ID in the query.
// Returns a *NotSingularError when more than one TweetTag ID is found.
// Returns a *NotFoundError when no entities are found.
func (ttq *TweetTagQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ttq.Limit(2).IDs(setContextOp(ctx, ttq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tweettag.Label}
	default:
		err = &NotSingularError{tweettag.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ttq *TweetTagQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ttq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TweetTags.
func (ttq *TweetTagQuery) All(ctx context.Context) ([]*TweetTag, error) {
	ctx = setContextOp(ctx, ttq.ctx, ent.OpQueryAll)
	if err := ttq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TweetTag, *TweetTagQuery]()
	return withInterceptors[[]*TweetTag](ctx, ttq, qr, ttq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ttq *TweetTagQuery) AllX(ctx context.Context) []*TweetTag {
	nodes, err := ttq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TweetTag IDs.
func (ttq *TweetTagQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if ttq.ctx.Unique == nil && ttq.path != nil {
		ttq.Unique(true)
	}
	ctx = setContextOp(ctx, ttq.ctx, ent.OpQueryIDs)
	if err = ttq.Select(tweettag.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ttq *TweetTagQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ttq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ttq *TweetTagQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ttq.ctx, ent.OpQueryCount)
	if err := ttq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ttq, querierCount[*TweetTagQuery](), ttq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ttq *TweetTagQuery) CountX(ctx context.Context) int {
	count, err := ttq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ttq *TweetTagQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ttq.ctx, ent.OpQueryExist)
	switch _, err := ttq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ttq *TweetTagQuery) ExistX(ctx context.Context) bool {
	exist, err := ttq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TweetTagQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ttq *TweetTagQuery) Clone() *TweetTagQuery {
	if ttq == nil {
		return nil
	}
	return &TweetTagQuery{
		config:     ttq.config,
		ctx:        ttq.ctx.Clone(),
		order:      append([]tweettag.OrderOption{}, ttq.order...),
		inters:     append([]Interceptor{}, ttq.inters...),
		predicates: append([]predicate.TweetTag{}, ttq.predicates...),
		withTag:    ttq.withTag.Clone(),
		withTweet:  ttq.withTweet.Clone(),
		// clone intermediate query.
		sql:  ttq.sql.Clone(),
		path: ttq.path,
	}
}

// WithTag tells the query-builder to eager-load the nodes that are connected to
// the "tag" edge. The optional arguments are used to configure the query builder of the edge.
func (ttq *TweetTagQuery) WithTag(opts ...func(*TagQuery)) *TweetTagQuery {
	query := (&TagClient{config: ttq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ttq.withTag = query
	return ttq
}

// WithTweet tells the query-builder to eager-load the nodes that are connected to
// the "tweet" edge. The optional arguments are used to configure the query builder of the edge.
func (ttq *TweetTagQuery) WithTweet(opts ...func(*TweetQuery)) *TweetTagQuery {
	query := (&TweetClient{config: ttq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ttq.withTweet = query
	return ttq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AddedAt time.Time `json:"added_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TweetTag.Query().
//		GroupBy(tweettag.FieldAddedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ttq *TweetTagQuery) GroupBy(field string, fields ...string) *TweetTagGroupBy {
	ttq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TweetTagGroupBy{build: ttq}
	grbuild.flds = &ttq.ctx.Fields
	grbuild.label = tweettag.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AddedAt time.Time `json:"added_at,omitempty"`
//	}
//
//	client.TweetTag.Query().
//		Select(tweettag.FieldAddedAt).
//		Scan(ctx, &v)
func (ttq *TweetTagQuery) Select(fields ...string) *TweetTagSelect {
	ttq.ctx.Fields = append(ttq.ctx.Fields, fields...)
	sbuild := &TweetTagSelect{TweetTagQuery: ttq}
	sbuild.label = tweettag.Label
	sbuild.flds, sbuild.scan = &ttq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TweetTagSelect configured with the given aggregations.
func (ttq *TweetTagQuery) Aggregate(fns ...AggregateFunc) *TweetTagSelect {
	return ttq.Select().Aggregate(fns...)
}

func (ttq *TweetTagQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ttq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ttq); err != nil {
				return err
			}
		}
	}
	for _, f := range ttq.ctx.Fields {
		if !tweettag.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ttq.path != nil {
		prev, err := ttq.path(ctx)
		if err != nil {
			return err
		}
		ttq.sql = prev
	}
	return nil
}

func (ttq *TweetTagQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TweetTag, error) {
	var (
		nodes       = []*TweetTag{}
		_spec       = ttq.querySpec()
		loadedTypes = [2]bool{
			ttq.withTag != nil,
			ttq.withTweet != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TweetTag).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TweetTag{config: ttq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ttq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ttq.withTag; query != nil {
		if err := ttq.loadTag(ctx, query, nodes, nil,
			func(n *TweetTag, e *Tag) { n.Edges.Tag = e }); err != nil {
			return nil, err
		}
	}
	if query := ttq.withTweet; query != nil {
		if err := ttq.loadTweet(ctx, query, nodes, nil,
			func(n *TweetTag, e *Tweet) { n.Edges.Tweet = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ttq *TweetTagQuery) loadTag(ctx context.Context, query *TagQuery, nodes []*TweetTag, init func(*TweetTag), assign func(*TweetTag, *Tag)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*TweetTag)
	for i := range nodes {
		fk := nodes[i].TagID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(tag.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "tag_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ttq *TweetTagQuery) loadTweet(ctx context.Context, query *TweetQuery, nodes []*TweetTag, init func(*TweetTag), assign func(*TweetTag, *Tweet)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*TweetTag)
	for i := range nodes {
		fk := nodes[i].TweetID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(tweet.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "tweet_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ttq *TweetTagQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ttq.querySpec()
	_spec.Node.Columns = ttq.ctx.Fields
	if len(ttq.ctx.Fields) > 0 {
		_spec.Unique = ttq.ctx.Unique != nil && *ttq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ttq.driver, _spec)
}

func (ttq *TweetTagQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(tweettag.Table, tweettag.Columns, sqlgraph.NewFieldSpec(tweettag.FieldID, field.TypeUUID))
	_spec.From = ttq.sql
	if unique := ttq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ttq.path != nil {
		_spec.Unique = true
	}
	if fields := ttq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tweettag.FieldID)
		for i := range fields {
			if fields[i] != tweettag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if ttq.withTag != nil {
			_spec.Node.AddColumnOnce(tweettag.FieldTagID)
		}
		if ttq.withTweet != nil {
			_spec.Node.AddColumnOnce(tweettag.FieldTweetID)
		}
	}
	if ps := ttq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ttq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ttq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ttq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ttq *TweetTagQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ttq.driver.Dialect())
	t1 := builder.Table(tweettag.Table)
	columns := ttq.ctx.Fields
	if len(columns) == 0 {
		columns = tweettag.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ttq.sql != nil {
		selector = ttq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ttq.ctx.Unique != nil && *ttq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ttq.predicates {
		p(selector)
	}
	for _, p := range ttq.order {
		p(selector)
	}
	if offset := ttq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ttq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TweetTagGroupBy is the group-by builder for TweetTag entities.
type TweetTagGroupBy struct {
	selector
	build *TweetTagQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ttgb *TweetTagGroupBy) Aggregate(fns ...AggregateFunc) *TweetTagGroupBy {
	ttgb.fns = append(ttgb.fns, fns...)
	return ttgb
}

// Scan applies the selector query and scans the result into the given value.
func (ttgb *TweetTagGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ttgb.build.ctx, ent.OpQueryGroupBy)
	if err := ttgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TweetTagQuery, *TweetTagGroupBy](ctx, ttgb.build, ttgb, ttgb.build.inters, v)
}

func (ttgb *TweetTagGroupBy) sqlScan(ctx context.Context, root *TweetTagQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ttgb.fns))
	for _, fn := range ttgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ttgb.flds)+len(ttgb.fns))
		for _, f := range *ttgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ttgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ttgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TweetTagSelect is the builder for selecting fields of TweetTag entities.
type TweetTagSelect struct {
	*TweetTagQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tts *TweetTagSelect) Aggregate(fns ...AggregateFunc) *TweetTagSelect {
	tts.fns = append(tts.fns, fns...)
	return tts
}

// Scan applies the selector query and scans the result into the given value.
func (tts *TweetTagSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tts.ctx, ent.OpQuerySelect)
	if err := tts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TweetTagQuery, *TweetTagSelect](ctx, tts.TweetTagQuery, tts, tts.inters, v)
}

func (tts *TweetTagSelect) sqlScan(ctx context.Context, root *TweetTagQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tts.fns))
	for _, fn := range tts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
