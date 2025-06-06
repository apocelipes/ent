// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/card"
	"entgo.io/ent/entc/integration/ent/predicate"
	"entgo.io/ent/entc/integration/ent/spec"
	"entgo.io/ent/schema/field"
)

// SpecQuery is the builder for querying Spec entities.
type SpecQuery struct {
	config
	ctx           *QueryContext
	order         []OrderFunc
	inters        []Interceptor
	predicates    []predicate.Spec
	withCard      *CardQuery
	modifiers     []func(*sql.Selector)
	withNamedCard map[string]*CardQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SpecQuery builder.
func (sq *SpecQuery) Where(ps ...predicate.Spec) *SpecQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *SpecQuery) Limit(limit int) *SpecQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *SpecQuery) Offset(offset int) *SpecQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SpecQuery) Unique(unique bool) *SpecQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *SpecQuery) Order(o ...OrderFunc) *SpecQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryCard chains the current query on the "card" edge.
func (sq *SpecQuery) QueryCard() *CardQuery {
	query := (&CardClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(spec.Table, spec.FieldID, selector),
			sqlgraph.To(card.Table, card.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, spec.CardTable, spec.CardPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Spec entity from the query.
// Returns a *NotFoundError when no Spec was found.
func (sq *SpecQuery) First(ctx context.Context) (*Spec, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{spec.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SpecQuery) FirstX(ctx context.Context) *Spec {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Spec ID from the query.
// Returns a *NotFoundError when no Spec ID was found.
func (sq *SpecQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{spec.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SpecQuery) FirstIDX(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Spec entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Spec entity is found.
// Returns a *NotFoundError when no Spec entities are found.
func (sq *SpecQuery) Only(ctx context.Context) (*Spec, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{spec.Label}
	default:
		return nil, &NotSingularError{spec.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SpecQuery) OnlyX(ctx context.Context) *Spec {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Spec ID in the query.
// Returns a *NotSingularError when more than one Spec ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SpecQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{spec.Label}
	default:
		err = &NotSingularError{spec.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SpecQuery) OnlyIDX(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Specs.
func (sq *SpecQuery) All(ctx context.Context) ([]*Spec, error) {
	ctx = setContextOp(ctx, sq.ctx, "All")
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Spec, *SpecQuery]()
	return withInterceptors[[]*Spec](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *SpecQuery) AllX(ctx context.Context) []*Spec {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Spec IDs.
func (sq *SpecQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = setContextOp(ctx, sq.ctx, "IDs")
	if err := sq.Select(spec.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SpecQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SpecQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, "Count")
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*SpecQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SpecQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SpecQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, "Exist")
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SpecQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SpecQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SpecQuery) Clone() *SpecQuery {
	if sq == nil {
		return nil
	}
	return &SpecQuery{
		config:     sq.config,
		ctx:        sq.ctx.Clone(),
		order:      append([]OrderFunc{}, sq.order...),
		inters:     append([]Interceptor{}, sq.inters...),
		predicates: append([]predicate.Spec{}, sq.predicates...),
		withCard:   sq.withCard.Clone(),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// WithCard tells the query-builder to eager-load the nodes that are connected to
// the "card" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SpecQuery) WithCard(opts ...func(*CardQuery)) *SpecQuery {
	query := (&CardClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withCard = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (sq *SpecQuery) GroupBy(field string, fields ...string) *SpecGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SpecGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = spec.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (sq *SpecQuery) Select(fields ...string) *SpecSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &SpecSelect{SpecQuery: sq}
	sbuild.label = spec.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SpecSelect configured with the given aggregations.
func (sq *SpecQuery) Aggregate(fns ...AggregateFunc) *SpecSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *SpecQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !spec.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SpecQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Spec, error) {
	var (
		nodes       = []*Spec{}
		_spec       = sq.querySpec()
		loadedTypes = [1]bool{
			sq.withCard != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Spec).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Spec{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withCard; query != nil {
		if err := sq.loadCard(ctx, query, nodes,
			func(n *Spec) { n.Edges.Card = []*Card{} },
			func(n *Spec, e *Card) { n.Edges.Card = append(n.Edges.Card, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range sq.withNamedCard {
		if err := sq.loadCard(ctx, query, nodes,
			func(n *Spec) { n.appendNamedCard(name) },
			func(n *Spec, e *Card) { n.appendNamedCard(name, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *SpecQuery) loadCard(ctx context.Context, query *CardQuery, nodes []*Spec, init func(*Spec), assign func(*Spec, *Card)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Spec)
	nids := make(map[int]map[*Spec]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(spec.CardTable)
		s.Join(joinT).On(s.C(card.FieldID), joinT.C(spec.CardPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(spec.CardPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(spec.CardPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Spec]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Card](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "card" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (sq *SpecQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SpecQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   spec.Table,
			Columns: spec.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: spec.FieldID,
			},
		},
		From:   sq.sql,
		Unique: true,
	}
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, spec.FieldID)
		for i := range fields {
			if fields[i] != spec.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SpecQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(spec.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = spec.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range sq.modifiers {
		m(selector)
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (sq *SpecQuery) ForUpdate(opts ...sql.LockOption) *SpecQuery {
	if sq.driver.Dialect() == dialect.Postgres {
		sq.Unique(false)
	}
	sq.modifiers = append(sq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return sq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (sq *SpecQuery) ForShare(opts ...sql.LockOption) *SpecQuery {
	if sq.driver.Dialect() == dialect.Postgres {
		sq.Unique(false)
	}
	sq.modifiers = append(sq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return sq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sq *SpecQuery) Modify(modifiers ...func(s *sql.Selector)) *SpecSelect {
	sq.modifiers = append(sq.modifiers, modifiers...)
	return sq.Select()
}

// WithNamedCard tells the query-builder to eager-load the nodes that are connected to the "card"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (sq *SpecQuery) WithNamedCard(name string, opts ...func(*CardQuery)) *SpecQuery {
	query := (&CardClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if sq.withNamedCard == nil {
		sq.withNamedCard = make(map[string]*CardQuery)
	}
	sq.withNamedCard[name] = query
	return sq
}

// SpecGroupBy is the group-by builder for Spec entities.
type SpecGroupBy struct {
	selector
	build *SpecQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SpecGroupBy) Aggregate(fns ...AggregateFunc) *SpecGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *SpecGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, "GroupBy")
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SpecQuery, *SpecGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *SpecGroupBy) sqlScan(ctx context.Context, root *SpecQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SpecSelect is the builder for selecting fields of Spec entities.
type SpecSelect struct {
	*SpecQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SpecSelect) Aggregate(fns ...AggregateFunc) *SpecSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SpecSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, "Select")
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SpecQuery, *SpecSelect](ctx, ss.SpecQuery, ss, ss.inters, v)
}

func (ss *SpecSelect) sqlScan(ctx context.Context, root *SpecQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ss *SpecSelect) Modify(modifiers ...func(s *sql.Selector)) *SpecSelect {
	ss.modifiers = append(ss.modifiers, modifiers...)
	return ss
}
