// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"prodcat/ent/predicate"
	"prodcat/ent/usersettings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserSettingsQuery is the builder for querying UserSettings entities.
type UserSettingsQuery struct {
	config
	ctx        *QueryContext
	order      []usersettings.OrderOption
	inters     []Interceptor
	predicates []predicate.UserSettings
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserSettingsQuery builder.
func (usq *UserSettingsQuery) Where(ps ...predicate.UserSettings) *UserSettingsQuery {
	usq.predicates = append(usq.predicates, ps...)
	return usq
}

// Limit the number of records to be returned by this query.
func (usq *UserSettingsQuery) Limit(limit int) *UserSettingsQuery {
	usq.ctx.Limit = &limit
	return usq
}

// Offset to start from.
func (usq *UserSettingsQuery) Offset(offset int) *UserSettingsQuery {
	usq.ctx.Offset = &offset
	return usq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (usq *UserSettingsQuery) Unique(unique bool) *UserSettingsQuery {
	usq.ctx.Unique = &unique
	return usq
}

// Order specifies how the records should be ordered.
func (usq *UserSettingsQuery) Order(o ...usersettings.OrderOption) *UserSettingsQuery {
	usq.order = append(usq.order, o...)
	return usq
}

// First returns the first UserSettings entity from the query.
// Returns a *NotFoundError when no UserSettings was found.
func (usq *UserSettingsQuery) First(ctx context.Context) (*UserSettings, error) {
	nodes, err := usq.Limit(1).All(setContextOp(ctx, usq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usersettings.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (usq *UserSettingsQuery) FirstX(ctx context.Context) *UserSettings {
	node, err := usq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserSettings ID from the query.
// Returns a *NotFoundError when no UserSettings ID was found.
func (usq *UserSettingsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = usq.Limit(1).IDs(setContextOp(ctx, usq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usersettings.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (usq *UserSettingsQuery) FirstIDX(ctx context.Context) int {
	id, err := usq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserSettings entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserSettings entity is found.
// Returns a *NotFoundError when no UserSettings entities are found.
func (usq *UserSettingsQuery) Only(ctx context.Context) (*UserSettings, error) {
	nodes, err := usq.Limit(2).All(setContextOp(ctx, usq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usersettings.Label}
	default:
		return nil, &NotSingularError{usersettings.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (usq *UserSettingsQuery) OnlyX(ctx context.Context) *UserSettings {
	node, err := usq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserSettings ID in the query.
// Returns a *NotSingularError when more than one UserSettings ID is found.
// Returns a *NotFoundError when no entities are found.
func (usq *UserSettingsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = usq.Limit(2).IDs(setContextOp(ctx, usq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usersettings.Label}
	default:
		err = &NotSingularError{usersettings.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (usq *UserSettingsQuery) OnlyIDX(ctx context.Context) int {
	id, err := usq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserSettingsSlice.
func (usq *UserSettingsQuery) All(ctx context.Context) ([]*UserSettings, error) {
	ctx = setContextOp(ctx, usq.ctx, "All")
	if err := usq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserSettings, *UserSettingsQuery]()
	return withInterceptors[[]*UserSettings](ctx, usq, qr, usq.inters)
}

// AllX is like All, but panics if an error occurs.
func (usq *UserSettingsQuery) AllX(ctx context.Context) []*UserSettings {
	nodes, err := usq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserSettings IDs.
func (usq *UserSettingsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if usq.ctx.Unique == nil && usq.path != nil {
		usq.Unique(true)
	}
	ctx = setContextOp(ctx, usq.ctx, "IDs")
	if err = usq.Select(usersettings.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (usq *UserSettingsQuery) IDsX(ctx context.Context) []int {
	ids, err := usq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (usq *UserSettingsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, usq.ctx, "Count")
	if err := usq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, usq, querierCount[*UserSettingsQuery](), usq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (usq *UserSettingsQuery) CountX(ctx context.Context) int {
	count, err := usq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (usq *UserSettingsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, usq.ctx, "Exist")
	switch _, err := usq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (usq *UserSettingsQuery) ExistX(ctx context.Context) bool {
	exist, err := usq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserSettingsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (usq *UserSettingsQuery) Clone() *UserSettingsQuery {
	if usq == nil {
		return nil
	}
	return &UserSettingsQuery{
		config:     usq.config,
		ctx:        usq.ctx.Clone(),
		order:      append([]usersettings.OrderOption{}, usq.order...),
		inters:     append([]Interceptor{}, usq.inters...),
		predicates: append([]predicate.UserSettings{}, usq.predicates...),
		// clone intermediate query.
		sql:  usq.sql.Clone(),
		path: usq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Frontend schema.DefaultSetting `json:"frontend,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserSettings.Query().
//		GroupBy(usersettings.FieldFrontend).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (usq *UserSettingsQuery) GroupBy(field string, fields ...string) *UserSettingsGroupBy {
	usq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserSettingsGroupBy{build: usq}
	grbuild.flds = &usq.ctx.Fields
	grbuild.label = usersettings.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Frontend schema.DefaultSetting `json:"frontend,omitempty"`
//	}
//
//	client.UserSettings.Query().
//		Select(usersettings.FieldFrontend).
//		Scan(ctx, &v)
func (usq *UserSettingsQuery) Select(fields ...string) *UserSettingsSelect {
	usq.ctx.Fields = append(usq.ctx.Fields, fields...)
	sbuild := &UserSettingsSelect{UserSettingsQuery: usq}
	sbuild.label = usersettings.Label
	sbuild.flds, sbuild.scan = &usq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserSettingsSelect configured with the given aggregations.
func (usq *UserSettingsQuery) Aggregate(fns ...AggregateFunc) *UserSettingsSelect {
	return usq.Select().Aggregate(fns...)
}

func (usq *UserSettingsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range usq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, usq); err != nil {
				return err
			}
		}
	}
	for _, f := range usq.ctx.Fields {
		if !usersettings.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if usq.path != nil {
		prev, err := usq.path(ctx)
		if err != nil {
			return err
		}
		usq.sql = prev
	}
	return nil
}

func (usq *UserSettingsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserSettings, error) {
	var (
		nodes = []*UserSettings{}
		_spec = usq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserSettings).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserSettings{config: usq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, usq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (usq *UserSettingsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := usq.querySpec()
	_spec.Node.Columns = usq.ctx.Fields
	if len(usq.ctx.Fields) > 0 {
		_spec.Unique = usq.ctx.Unique != nil && *usq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, usq.driver, _spec)
}

func (usq *UserSettingsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(usersettings.Table, usersettings.Columns, sqlgraph.NewFieldSpec(usersettings.FieldID, field.TypeInt))
	_spec.From = usq.sql
	if unique := usq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if usq.path != nil {
		_spec.Unique = true
	}
	if fields := usq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usersettings.FieldID)
		for i := range fields {
			if fields[i] != usersettings.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := usq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := usq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := usq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := usq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (usq *UserSettingsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(usq.driver.Dialect())
	t1 := builder.Table(usersettings.Table)
	columns := usq.ctx.Fields
	if len(columns) == 0 {
		columns = usersettings.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if usq.sql != nil {
		selector = usq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if usq.ctx.Unique != nil && *usq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range usq.predicates {
		p(selector)
	}
	for _, p := range usq.order {
		p(selector)
	}
	if offset := usq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := usq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserSettingsGroupBy is the group-by builder for UserSettings entities.
type UserSettingsGroupBy struct {
	selector
	build *UserSettingsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (usgb *UserSettingsGroupBy) Aggregate(fns ...AggregateFunc) *UserSettingsGroupBy {
	usgb.fns = append(usgb.fns, fns...)
	return usgb
}

// Scan applies the selector query and scans the result into the given value.
func (usgb *UserSettingsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, usgb.build.ctx, "GroupBy")
	if err := usgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserSettingsQuery, *UserSettingsGroupBy](ctx, usgb.build, usgb, usgb.build.inters, v)
}

func (usgb *UserSettingsGroupBy) sqlScan(ctx context.Context, root *UserSettingsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(usgb.fns))
	for _, fn := range usgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*usgb.flds)+len(usgb.fns))
		for _, f := range *usgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*usgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := usgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserSettingsSelect is the builder for selecting fields of UserSettings entities.
type UserSettingsSelect struct {
	*UserSettingsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (uss *UserSettingsSelect) Aggregate(fns ...AggregateFunc) *UserSettingsSelect {
	uss.fns = append(uss.fns, fns...)
	return uss
}

// Scan applies the selector query and scans the result into the given value.
func (uss *UserSettingsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uss.ctx, "Select")
	if err := uss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserSettingsQuery, *UserSettingsSelect](ctx, uss.UserSettingsQuery, uss, uss.inters, v)
}

func (uss *UserSettingsSelect) sqlScan(ctx context.Context, root *UserSettingsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(uss.fns))
	for _, fn := range uss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*uss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}