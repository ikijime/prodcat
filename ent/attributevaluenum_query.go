// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"prodcat/ent/attributevaluenum"
	"prodcat/ent/attributevariantnum"
	"prodcat/ent/predicate"
	"prodcat/ent/product"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AttributeValueNumQuery is the builder for querying AttributeValueNum entities.
type AttributeValueNumQuery struct {
	config
	ctx         *QueryContext
	order       []attributevaluenum.OrderOption
	inters      []Interceptor
	predicates  []predicate.AttributeValueNum
	withVariant *AttributeVariantNumQuery
	withProduct *ProductQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AttributeValueNumQuery builder.
func (avnq *AttributeValueNumQuery) Where(ps ...predicate.AttributeValueNum) *AttributeValueNumQuery {
	avnq.predicates = append(avnq.predicates, ps...)
	return avnq
}

// Limit the number of records to be returned by this query.
func (avnq *AttributeValueNumQuery) Limit(limit int) *AttributeValueNumQuery {
	avnq.ctx.Limit = &limit
	return avnq
}

// Offset to start from.
func (avnq *AttributeValueNumQuery) Offset(offset int) *AttributeValueNumQuery {
	avnq.ctx.Offset = &offset
	return avnq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (avnq *AttributeValueNumQuery) Unique(unique bool) *AttributeValueNumQuery {
	avnq.ctx.Unique = &unique
	return avnq
}

// Order specifies how the records should be ordered.
func (avnq *AttributeValueNumQuery) Order(o ...attributevaluenum.OrderOption) *AttributeValueNumQuery {
	avnq.order = append(avnq.order, o...)
	return avnq
}

// QueryVariant chains the current query on the "variant" edge.
func (avnq *AttributeValueNumQuery) QueryVariant() *AttributeVariantNumQuery {
	query := (&AttributeVariantNumClient{config: avnq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := avnq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := avnq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attributevaluenum.Table, attributevaluenum.FieldID, selector),
			sqlgraph.To(attributevariantnum.Table, attributevariantnum.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, attributevaluenum.VariantTable, attributevaluenum.VariantColumn),
		)
		fromU = sqlgraph.SetNeighbors(avnq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProduct chains the current query on the "product" edge.
func (avnq *AttributeValueNumQuery) QueryProduct() *ProductQuery {
	query := (&ProductClient{config: avnq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := avnq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := avnq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attributevaluenum.Table, attributevaluenum.FieldID, selector),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attributevaluenum.ProductTable, attributevaluenum.ProductColumn),
		)
		fromU = sqlgraph.SetNeighbors(avnq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AttributeValueNum entity from the query.
// Returns a *NotFoundError when no AttributeValueNum was found.
func (avnq *AttributeValueNumQuery) First(ctx context.Context) (*AttributeValueNum, error) {
	nodes, err := avnq.Limit(1).All(setContextOp(ctx, avnq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{attributevaluenum.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (avnq *AttributeValueNumQuery) FirstX(ctx context.Context) *AttributeValueNum {
	node, err := avnq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AttributeValueNum ID from the query.
// Returns a *NotFoundError when no AttributeValueNum ID was found.
func (avnq *AttributeValueNumQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = avnq.Limit(1).IDs(setContextOp(ctx, avnq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{attributevaluenum.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (avnq *AttributeValueNumQuery) FirstIDX(ctx context.Context) int {
	id, err := avnq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AttributeValueNum entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AttributeValueNum entity is found.
// Returns a *NotFoundError when no AttributeValueNum entities are found.
func (avnq *AttributeValueNumQuery) Only(ctx context.Context) (*AttributeValueNum, error) {
	nodes, err := avnq.Limit(2).All(setContextOp(ctx, avnq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{attributevaluenum.Label}
	default:
		return nil, &NotSingularError{attributevaluenum.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (avnq *AttributeValueNumQuery) OnlyX(ctx context.Context) *AttributeValueNum {
	node, err := avnq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AttributeValueNum ID in the query.
// Returns a *NotSingularError when more than one AttributeValueNum ID is found.
// Returns a *NotFoundError when no entities are found.
func (avnq *AttributeValueNumQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = avnq.Limit(2).IDs(setContextOp(ctx, avnq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{attributevaluenum.Label}
	default:
		err = &NotSingularError{attributevaluenum.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (avnq *AttributeValueNumQuery) OnlyIDX(ctx context.Context) int {
	id, err := avnq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AttributeValueNums.
func (avnq *AttributeValueNumQuery) All(ctx context.Context) ([]*AttributeValueNum, error) {
	ctx = setContextOp(ctx, avnq.ctx, "All")
	if err := avnq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AttributeValueNum, *AttributeValueNumQuery]()
	return withInterceptors[[]*AttributeValueNum](ctx, avnq, qr, avnq.inters)
}

// AllX is like All, but panics if an error occurs.
func (avnq *AttributeValueNumQuery) AllX(ctx context.Context) []*AttributeValueNum {
	nodes, err := avnq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AttributeValueNum IDs.
func (avnq *AttributeValueNumQuery) IDs(ctx context.Context) (ids []int, err error) {
	if avnq.ctx.Unique == nil && avnq.path != nil {
		avnq.Unique(true)
	}
	ctx = setContextOp(ctx, avnq.ctx, "IDs")
	if err = avnq.Select(attributevaluenum.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (avnq *AttributeValueNumQuery) IDsX(ctx context.Context) []int {
	ids, err := avnq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (avnq *AttributeValueNumQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, avnq.ctx, "Count")
	if err := avnq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, avnq, querierCount[*AttributeValueNumQuery](), avnq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (avnq *AttributeValueNumQuery) CountX(ctx context.Context) int {
	count, err := avnq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (avnq *AttributeValueNumQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, avnq.ctx, "Exist")
	switch _, err := avnq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (avnq *AttributeValueNumQuery) ExistX(ctx context.Context) bool {
	exist, err := avnq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AttributeValueNumQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (avnq *AttributeValueNumQuery) Clone() *AttributeValueNumQuery {
	if avnq == nil {
		return nil
	}
	return &AttributeValueNumQuery{
		config:      avnq.config,
		ctx:         avnq.ctx.Clone(),
		order:       append([]attributevaluenum.OrderOption{}, avnq.order...),
		inters:      append([]Interceptor{}, avnq.inters...),
		predicates:  append([]predicate.AttributeValueNum{}, avnq.predicates...),
		withVariant: avnq.withVariant.Clone(),
		withProduct: avnq.withProduct.Clone(),
		// clone intermediate query.
		sql:  avnq.sql.Clone(),
		path: avnq.path,
	}
}

// WithVariant tells the query-builder to eager-load the nodes that are connected to
// the "variant" edge. The optional arguments are used to configure the query builder of the edge.
func (avnq *AttributeValueNumQuery) WithVariant(opts ...func(*AttributeVariantNumQuery)) *AttributeValueNumQuery {
	query := (&AttributeVariantNumClient{config: avnq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	avnq.withVariant = query
	return avnq
}

// WithProduct tells the query-builder to eager-load the nodes that are connected to
// the "product" edge. The optional arguments are used to configure the query builder of the edge.
func (avnq *AttributeValueNumQuery) WithProduct(opts ...func(*ProductQuery)) *AttributeValueNumQuery {
	query := (&ProductClient{config: avnq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	avnq.withProduct = query
	return avnq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		VariantID int `json:"variant_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AttributeValueNum.Query().
//		GroupBy(attributevaluenum.FieldVariantID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (avnq *AttributeValueNumQuery) GroupBy(field string, fields ...string) *AttributeValueNumGroupBy {
	avnq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AttributeValueNumGroupBy{build: avnq}
	grbuild.flds = &avnq.ctx.Fields
	grbuild.label = attributevaluenum.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		VariantID int `json:"variant_id,omitempty"`
//	}
//
//	client.AttributeValueNum.Query().
//		Select(attributevaluenum.FieldVariantID).
//		Scan(ctx, &v)
func (avnq *AttributeValueNumQuery) Select(fields ...string) *AttributeValueNumSelect {
	avnq.ctx.Fields = append(avnq.ctx.Fields, fields...)
	sbuild := &AttributeValueNumSelect{AttributeValueNumQuery: avnq}
	sbuild.label = attributevaluenum.Label
	sbuild.flds, sbuild.scan = &avnq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AttributeValueNumSelect configured with the given aggregations.
func (avnq *AttributeValueNumQuery) Aggregate(fns ...AggregateFunc) *AttributeValueNumSelect {
	return avnq.Select().Aggregate(fns...)
}

func (avnq *AttributeValueNumQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range avnq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, avnq); err != nil {
				return err
			}
		}
	}
	for _, f := range avnq.ctx.Fields {
		if !attributevaluenum.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if avnq.path != nil {
		prev, err := avnq.path(ctx)
		if err != nil {
			return err
		}
		avnq.sql = prev
	}
	return nil
}

func (avnq *AttributeValueNumQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AttributeValueNum, error) {
	var (
		nodes       = []*AttributeValueNum{}
		_spec       = avnq.querySpec()
		loadedTypes = [2]bool{
			avnq.withVariant != nil,
			avnq.withProduct != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AttributeValueNum).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AttributeValueNum{config: avnq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, avnq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := avnq.withVariant; query != nil {
		if err := avnq.loadVariant(ctx, query, nodes, nil,
			func(n *AttributeValueNum, e *AttributeVariantNum) { n.Edges.Variant = e }); err != nil {
			return nil, err
		}
	}
	if query := avnq.withProduct; query != nil {
		if err := avnq.loadProduct(ctx, query, nodes, nil,
			func(n *AttributeValueNum, e *Product) { n.Edges.Product = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (avnq *AttributeValueNumQuery) loadVariant(ctx context.Context, query *AttributeVariantNumQuery, nodes []*AttributeValueNum, init func(*AttributeValueNum), assign func(*AttributeValueNum, *AttributeVariantNum)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*AttributeValueNum)
	for i := range nodes {
		fk := nodes[i].VariantID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(attributevariantnum.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "variant_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (avnq *AttributeValueNumQuery) loadProduct(ctx context.Context, query *ProductQuery, nodes []*AttributeValueNum, init func(*AttributeValueNum), assign func(*AttributeValueNum, *Product)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*AttributeValueNum)
	for i := range nodes {
		fk := nodes[i].ProductID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(product.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "product_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (avnq *AttributeValueNumQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := avnq.querySpec()
	_spec.Node.Columns = avnq.ctx.Fields
	if len(avnq.ctx.Fields) > 0 {
		_spec.Unique = avnq.ctx.Unique != nil && *avnq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, avnq.driver, _spec)
}

func (avnq *AttributeValueNumQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(attributevaluenum.Table, attributevaluenum.Columns, sqlgraph.NewFieldSpec(attributevaluenum.FieldID, field.TypeInt))
	_spec.From = avnq.sql
	if unique := avnq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if avnq.path != nil {
		_spec.Unique = true
	}
	if fields := avnq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attributevaluenum.FieldID)
		for i := range fields {
			if fields[i] != attributevaluenum.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if avnq.withVariant != nil {
			_spec.Node.AddColumnOnce(attributevaluenum.FieldVariantID)
		}
		if avnq.withProduct != nil {
			_spec.Node.AddColumnOnce(attributevaluenum.FieldProductID)
		}
	}
	if ps := avnq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := avnq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := avnq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := avnq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (avnq *AttributeValueNumQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(avnq.driver.Dialect())
	t1 := builder.Table(attributevaluenum.Table)
	columns := avnq.ctx.Fields
	if len(columns) == 0 {
		columns = attributevaluenum.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if avnq.sql != nil {
		selector = avnq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if avnq.ctx.Unique != nil && *avnq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range avnq.predicates {
		p(selector)
	}
	for _, p := range avnq.order {
		p(selector)
	}
	if offset := avnq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := avnq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AttributeValueNumGroupBy is the group-by builder for AttributeValueNum entities.
type AttributeValueNumGroupBy struct {
	selector
	build *AttributeValueNumQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (avngb *AttributeValueNumGroupBy) Aggregate(fns ...AggregateFunc) *AttributeValueNumGroupBy {
	avngb.fns = append(avngb.fns, fns...)
	return avngb
}

// Scan applies the selector query and scans the result into the given value.
func (avngb *AttributeValueNumGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, avngb.build.ctx, "GroupBy")
	if err := avngb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttributeValueNumQuery, *AttributeValueNumGroupBy](ctx, avngb.build, avngb, avngb.build.inters, v)
}

func (avngb *AttributeValueNumGroupBy) sqlScan(ctx context.Context, root *AttributeValueNumQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(avngb.fns))
	for _, fn := range avngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*avngb.flds)+len(avngb.fns))
		for _, f := range *avngb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*avngb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := avngb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AttributeValueNumSelect is the builder for selecting fields of AttributeValueNum entities.
type AttributeValueNumSelect struct {
	*AttributeValueNumQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (avns *AttributeValueNumSelect) Aggregate(fns ...AggregateFunc) *AttributeValueNumSelect {
	avns.fns = append(avns.fns, fns...)
	return avns
}

// Scan applies the selector query and scans the result into the given value.
func (avns *AttributeValueNumSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, avns.ctx, "Select")
	if err := avns.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttributeValueNumQuery, *AttributeValueNumSelect](ctx, avns.AttributeValueNumQuery, avns, avns.inters, v)
}

func (avns *AttributeValueNumSelect) sqlScan(ctx context.Context, root *AttributeValueNumQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(avns.fns))
	for _, fn := range avns.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*avns.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := avns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}