// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"prodcat/ent/attributevaluestring"
	"prodcat/ent/attributevariantstring"
	"prodcat/ent/predicate"
	"prodcat/ent/product"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AttributeValueStringQuery is the builder for querying AttributeValueString entities.
type AttributeValueStringQuery struct {
	config
	ctx         *QueryContext
	order       []attributevaluestring.OrderOption
	inters      []Interceptor
	predicates  []predicate.AttributeValueString
	withVariant *AttributeVariantStringQuery
	withProduct *ProductQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AttributeValueStringQuery builder.
func (avsq *AttributeValueStringQuery) Where(ps ...predicate.AttributeValueString) *AttributeValueStringQuery {
	avsq.predicates = append(avsq.predicates, ps...)
	return avsq
}

// Limit the number of records to be returned by this query.
func (avsq *AttributeValueStringQuery) Limit(limit int) *AttributeValueStringQuery {
	avsq.ctx.Limit = &limit
	return avsq
}

// Offset to start from.
func (avsq *AttributeValueStringQuery) Offset(offset int) *AttributeValueStringQuery {
	avsq.ctx.Offset = &offset
	return avsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (avsq *AttributeValueStringQuery) Unique(unique bool) *AttributeValueStringQuery {
	avsq.ctx.Unique = &unique
	return avsq
}

// Order specifies how the records should be ordered.
func (avsq *AttributeValueStringQuery) Order(o ...attributevaluestring.OrderOption) *AttributeValueStringQuery {
	avsq.order = append(avsq.order, o...)
	return avsq
}

// QueryVariant chains the current query on the "variant" edge.
func (avsq *AttributeValueStringQuery) QueryVariant() *AttributeVariantStringQuery {
	query := (&AttributeVariantStringClient{config: avsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := avsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := avsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attributevaluestring.Table, attributevaluestring.FieldID, selector),
			sqlgraph.To(attributevariantstring.Table, attributevariantstring.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, attributevaluestring.VariantTable, attributevaluestring.VariantColumn),
		)
		fromU = sqlgraph.SetNeighbors(avsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProduct chains the current query on the "product" edge.
func (avsq *AttributeValueStringQuery) QueryProduct() *ProductQuery {
	query := (&ProductClient{config: avsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := avsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := avsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attributevaluestring.Table, attributevaluestring.FieldID, selector),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attributevaluestring.ProductTable, attributevaluestring.ProductColumn),
		)
		fromU = sqlgraph.SetNeighbors(avsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AttributeValueString entity from the query.
// Returns a *NotFoundError when no AttributeValueString was found.
func (avsq *AttributeValueStringQuery) First(ctx context.Context) (*AttributeValueString, error) {
	nodes, err := avsq.Limit(1).All(setContextOp(ctx, avsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{attributevaluestring.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (avsq *AttributeValueStringQuery) FirstX(ctx context.Context) *AttributeValueString {
	node, err := avsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AttributeValueString ID from the query.
// Returns a *NotFoundError when no AttributeValueString ID was found.
func (avsq *AttributeValueStringQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = avsq.Limit(1).IDs(setContextOp(ctx, avsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{attributevaluestring.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (avsq *AttributeValueStringQuery) FirstIDX(ctx context.Context) int {
	id, err := avsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AttributeValueString entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AttributeValueString entity is found.
// Returns a *NotFoundError when no AttributeValueString entities are found.
func (avsq *AttributeValueStringQuery) Only(ctx context.Context) (*AttributeValueString, error) {
	nodes, err := avsq.Limit(2).All(setContextOp(ctx, avsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{attributevaluestring.Label}
	default:
		return nil, &NotSingularError{attributevaluestring.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (avsq *AttributeValueStringQuery) OnlyX(ctx context.Context) *AttributeValueString {
	node, err := avsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AttributeValueString ID in the query.
// Returns a *NotSingularError when more than one AttributeValueString ID is found.
// Returns a *NotFoundError when no entities are found.
func (avsq *AttributeValueStringQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = avsq.Limit(2).IDs(setContextOp(ctx, avsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{attributevaluestring.Label}
	default:
		err = &NotSingularError{attributevaluestring.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (avsq *AttributeValueStringQuery) OnlyIDX(ctx context.Context) int {
	id, err := avsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AttributeValueStrings.
func (avsq *AttributeValueStringQuery) All(ctx context.Context) ([]*AttributeValueString, error) {
	ctx = setContextOp(ctx, avsq.ctx, "All")
	if err := avsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AttributeValueString, *AttributeValueStringQuery]()
	return withInterceptors[[]*AttributeValueString](ctx, avsq, qr, avsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (avsq *AttributeValueStringQuery) AllX(ctx context.Context) []*AttributeValueString {
	nodes, err := avsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AttributeValueString IDs.
func (avsq *AttributeValueStringQuery) IDs(ctx context.Context) (ids []int, err error) {
	if avsq.ctx.Unique == nil && avsq.path != nil {
		avsq.Unique(true)
	}
	ctx = setContextOp(ctx, avsq.ctx, "IDs")
	if err = avsq.Select(attributevaluestring.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (avsq *AttributeValueStringQuery) IDsX(ctx context.Context) []int {
	ids, err := avsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (avsq *AttributeValueStringQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, avsq.ctx, "Count")
	if err := avsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, avsq, querierCount[*AttributeValueStringQuery](), avsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (avsq *AttributeValueStringQuery) CountX(ctx context.Context) int {
	count, err := avsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (avsq *AttributeValueStringQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, avsq.ctx, "Exist")
	switch _, err := avsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (avsq *AttributeValueStringQuery) ExistX(ctx context.Context) bool {
	exist, err := avsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AttributeValueStringQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (avsq *AttributeValueStringQuery) Clone() *AttributeValueStringQuery {
	if avsq == nil {
		return nil
	}
	return &AttributeValueStringQuery{
		config:      avsq.config,
		ctx:         avsq.ctx.Clone(),
		order:       append([]attributevaluestring.OrderOption{}, avsq.order...),
		inters:      append([]Interceptor{}, avsq.inters...),
		predicates:  append([]predicate.AttributeValueString{}, avsq.predicates...),
		withVariant: avsq.withVariant.Clone(),
		withProduct: avsq.withProduct.Clone(),
		// clone intermediate query.
		sql:  avsq.sql.Clone(),
		path: avsq.path,
	}
}

// WithVariant tells the query-builder to eager-load the nodes that are connected to
// the "variant" edge. The optional arguments are used to configure the query builder of the edge.
func (avsq *AttributeValueStringQuery) WithVariant(opts ...func(*AttributeVariantStringQuery)) *AttributeValueStringQuery {
	query := (&AttributeVariantStringClient{config: avsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	avsq.withVariant = query
	return avsq
}

// WithProduct tells the query-builder to eager-load the nodes that are connected to
// the "product" edge. The optional arguments are used to configure the query builder of the edge.
func (avsq *AttributeValueStringQuery) WithProduct(opts ...func(*ProductQuery)) *AttributeValueStringQuery {
	query := (&ProductClient{config: avsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	avsq.withProduct = query
	return avsq
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
//	client.AttributeValueString.Query().
//		GroupBy(attributevaluestring.FieldVariantID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (avsq *AttributeValueStringQuery) GroupBy(field string, fields ...string) *AttributeValueStringGroupBy {
	avsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AttributeValueStringGroupBy{build: avsq}
	grbuild.flds = &avsq.ctx.Fields
	grbuild.label = attributevaluestring.Label
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
//	client.AttributeValueString.Query().
//		Select(attributevaluestring.FieldVariantID).
//		Scan(ctx, &v)
func (avsq *AttributeValueStringQuery) Select(fields ...string) *AttributeValueStringSelect {
	avsq.ctx.Fields = append(avsq.ctx.Fields, fields...)
	sbuild := &AttributeValueStringSelect{AttributeValueStringQuery: avsq}
	sbuild.label = attributevaluestring.Label
	sbuild.flds, sbuild.scan = &avsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AttributeValueStringSelect configured with the given aggregations.
func (avsq *AttributeValueStringQuery) Aggregate(fns ...AggregateFunc) *AttributeValueStringSelect {
	return avsq.Select().Aggregate(fns...)
}

func (avsq *AttributeValueStringQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range avsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, avsq); err != nil {
				return err
			}
		}
	}
	for _, f := range avsq.ctx.Fields {
		if !attributevaluestring.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if avsq.path != nil {
		prev, err := avsq.path(ctx)
		if err != nil {
			return err
		}
		avsq.sql = prev
	}
	return nil
}

func (avsq *AttributeValueStringQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AttributeValueString, error) {
	var (
		nodes       = []*AttributeValueString{}
		_spec       = avsq.querySpec()
		loadedTypes = [2]bool{
			avsq.withVariant != nil,
			avsq.withProduct != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AttributeValueString).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AttributeValueString{config: avsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, avsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := avsq.withVariant; query != nil {
		if err := avsq.loadVariant(ctx, query, nodes, nil,
			func(n *AttributeValueString, e *AttributeVariantString) { n.Edges.Variant = e }); err != nil {
			return nil, err
		}
	}
	if query := avsq.withProduct; query != nil {
		if err := avsq.loadProduct(ctx, query, nodes, nil,
			func(n *AttributeValueString, e *Product) { n.Edges.Product = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (avsq *AttributeValueStringQuery) loadVariant(ctx context.Context, query *AttributeVariantStringQuery, nodes []*AttributeValueString, init func(*AttributeValueString), assign func(*AttributeValueString, *AttributeVariantString)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*AttributeValueString)
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
	query.Where(attributevariantstring.IDIn(ids...))
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
func (avsq *AttributeValueStringQuery) loadProduct(ctx context.Context, query *ProductQuery, nodes []*AttributeValueString, init func(*AttributeValueString), assign func(*AttributeValueString, *Product)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*AttributeValueString)
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

func (avsq *AttributeValueStringQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := avsq.querySpec()
	_spec.Node.Columns = avsq.ctx.Fields
	if len(avsq.ctx.Fields) > 0 {
		_spec.Unique = avsq.ctx.Unique != nil && *avsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, avsq.driver, _spec)
}

func (avsq *AttributeValueStringQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(attributevaluestring.Table, attributevaluestring.Columns, sqlgraph.NewFieldSpec(attributevaluestring.FieldID, field.TypeInt))
	_spec.From = avsq.sql
	if unique := avsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if avsq.path != nil {
		_spec.Unique = true
	}
	if fields := avsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attributevaluestring.FieldID)
		for i := range fields {
			if fields[i] != attributevaluestring.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if avsq.withVariant != nil {
			_spec.Node.AddColumnOnce(attributevaluestring.FieldVariantID)
		}
		if avsq.withProduct != nil {
			_spec.Node.AddColumnOnce(attributevaluestring.FieldProductID)
		}
	}
	if ps := avsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := avsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := avsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := avsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (avsq *AttributeValueStringQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(avsq.driver.Dialect())
	t1 := builder.Table(attributevaluestring.Table)
	columns := avsq.ctx.Fields
	if len(columns) == 0 {
		columns = attributevaluestring.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if avsq.sql != nil {
		selector = avsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if avsq.ctx.Unique != nil && *avsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range avsq.predicates {
		p(selector)
	}
	for _, p := range avsq.order {
		p(selector)
	}
	if offset := avsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := avsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AttributeValueStringGroupBy is the group-by builder for AttributeValueString entities.
type AttributeValueStringGroupBy struct {
	selector
	build *AttributeValueStringQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (avsgb *AttributeValueStringGroupBy) Aggregate(fns ...AggregateFunc) *AttributeValueStringGroupBy {
	avsgb.fns = append(avsgb.fns, fns...)
	return avsgb
}

// Scan applies the selector query and scans the result into the given value.
func (avsgb *AttributeValueStringGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, avsgb.build.ctx, "GroupBy")
	if err := avsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttributeValueStringQuery, *AttributeValueStringGroupBy](ctx, avsgb.build, avsgb, avsgb.build.inters, v)
}

func (avsgb *AttributeValueStringGroupBy) sqlScan(ctx context.Context, root *AttributeValueStringQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(avsgb.fns))
	for _, fn := range avsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*avsgb.flds)+len(avsgb.fns))
		for _, f := range *avsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*avsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := avsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AttributeValueStringSelect is the builder for selecting fields of AttributeValueString entities.
type AttributeValueStringSelect struct {
	*AttributeValueStringQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (avss *AttributeValueStringSelect) Aggregate(fns ...AggregateFunc) *AttributeValueStringSelect {
	avss.fns = append(avss.fns, fns...)
	return avss
}

// Scan applies the selector query and scans the result into the given value.
func (avss *AttributeValueStringSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, avss.ctx, "Select")
	if err := avss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttributeValueStringQuery, *AttributeValueStringSelect](ctx, avss.AttributeValueStringQuery, avss, avss.inters, v)
}

func (avss *AttributeValueStringSelect) sqlScan(ctx context.Context, root *AttributeValueStringQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(avss.fns))
	for _, fn := range avss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*avss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := avss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
