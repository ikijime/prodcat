// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"prodcat/ent/brand"
	"prodcat/ent/predicate"
	"prodcat/ent/product"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BrandUpdate is the builder for updating Brand entities.
type BrandUpdate struct {
	config
	hooks    []Hook
	mutation *BrandMutation
}

// Where appends a list predicates to the BrandUpdate builder.
func (bu *BrandUpdate) Where(ps ...predicate.Brand) *BrandUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetName sets the "name" field.
func (bu *BrandUpdate) SetName(s string) *BrandUpdate {
	bu.mutation.SetName(s)
	return bu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (bu *BrandUpdate) SetNillableName(s *string) *BrandUpdate {
	if s != nil {
		bu.SetName(*s)
	}
	return bu
}

// AddProductIDs adds the "product" edge to the Product entity by IDs.
func (bu *BrandUpdate) AddProductIDs(ids ...int) *BrandUpdate {
	bu.mutation.AddProductIDs(ids...)
	return bu
}

// AddProduct adds the "product" edges to the Product entity.
func (bu *BrandUpdate) AddProduct(p ...*Product) *BrandUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bu.AddProductIDs(ids...)
}

// Mutation returns the BrandMutation object of the builder.
func (bu *BrandUpdate) Mutation() *BrandMutation {
	return bu.mutation
}

// ClearProduct clears all "product" edges to the Product entity.
func (bu *BrandUpdate) ClearProduct() *BrandUpdate {
	bu.mutation.ClearProduct()
	return bu
}

// RemoveProductIDs removes the "product" edge to Product entities by IDs.
func (bu *BrandUpdate) RemoveProductIDs(ids ...int) *BrandUpdate {
	bu.mutation.RemoveProductIDs(ids...)
	return bu
}

// RemoveProduct removes "product" edges to Product entities.
func (bu *BrandUpdate) RemoveProduct(p ...*Product) *BrandUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bu.RemoveProductIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BrandUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BrandUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BrandUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BrandUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BrandUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(brand.Table, brand.Columns, sqlgraph.NewFieldSpec(brand.FieldID, field.TypeInt))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Name(); ok {
		_spec.SetField(brand.FieldName, field.TypeString, value)
	}
	if bu.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   brand.ProductTable,
			Columns: []string{brand.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedProductIDs(); len(nodes) > 0 && !bu.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   brand.ProductTable,
			Columns: []string{brand.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   brand.ProductTable,
			Columns: []string{brand.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{brand.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BrandUpdateOne is the builder for updating a single Brand entity.
type BrandUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BrandMutation
}

// SetName sets the "name" field.
func (buo *BrandUpdateOne) SetName(s string) *BrandUpdateOne {
	buo.mutation.SetName(s)
	return buo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (buo *BrandUpdateOne) SetNillableName(s *string) *BrandUpdateOne {
	if s != nil {
		buo.SetName(*s)
	}
	return buo
}

// AddProductIDs adds the "product" edge to the Product entity by IDs.
func (buo *BrandUpdateOne) AddProductIDs(ids ...int) *BrandUpdateOne {
	buo.mutation.AddProductIDs(ids...)
	return buo
}

// AddProduct adds the "product" edges to the Product entity.
func (buo *BrandUpdateOne) AddProduct(p ...*Product) *BrandUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return buo.AddProductIDs(ids...)
}

// Mutation returns the BrandMutation object of the builder.
func (buo *BrandUpdateOne) Mutation() *BrandMutation {
	return buo.mutation
}

// ClearProduct clears all "product" edges to the Product entity.
func (buo *BrandUpdateOne) ClearProduct() *BrandUpdateOne {
	buo.mutation.ClearProduct()
	return buo
}

// RemoveProductIDs removes the "product" edge to Product entities by IDs.
func (buo *BrandUpdateOne) RemoveProductIDs(ids ...int) *BrandUpdateOne {
	buo.mutation.RemoveProductIDs(ids...)
	return buo
}

// RemoveProduct removes "product" edges to Product entities.
func (buo *BrandUpdateOne) RemoveProduct(p ...*Product) *BrandUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return buo.RemoveProductIDs(ids...)
}

// Where appends a list predicates to the BrandUpdate builder.
func (buo *BrandUpdateOne) Where(ps ...predicate.Brand) *BrandUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BrandUpdateOne) Select(field string, fields ...string) *BrandUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Brand entity.
func (buo *BrandUpdateOne) Save(ctx context.Context) (*Brand, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BrandUpdateOne) SaveX(ctx context.Context) *Brand {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BrandUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BrandUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BrandUpdateOne) sqlSave(ctx context.Context) (_node *Brand, err error) {
	_spec := sqlgraph.NewUpdateSpec(brand.Table, brand.Columns, sqlgraph.NewFieldSpec(brand.FieldID, field.TypeInt))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Brand.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, brand.FieldID)
		for _, f := range fields {
			if !brand.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != brand.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Name(); ok {
		_spec.SetField(brand.FieldName, field.TypeString, value)
	}
	if buo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   brand.ProductTable,
			Columns: []string{brand.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedProductIDs(); len(nodes) > 0 && !buo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   brand.ProductTable,
			Columns: []string{brand.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   brand.ProductTable,
			Columns: []string{brand.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Brand{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{brand.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
