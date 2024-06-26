// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"prodcat/ent/attributevaluestring"
	"prodcat/ent/attributevariantstring"
	"prodcat/ent/product"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AttributeValueStringCreate is the builder for creating a AttributeValueString entity.
type AttributeValueStringCreate struct {
	config
	mutation *AttributeValueStringMutation
	hooks    []Hook
}

// SetVariantID sets the "variant_id" field.
func (avsc *AttributeValueStringCreate) SetVariantID(i int) *AttributeValueStringCreate {
	avsc.mutation.SetVariantID(i)
	return avsc
}

// SetProductID sets the "product_id" field.
func (avsc *AttributeValueStringCreate) SetProductID(i int) *AttributeValueStringCreate {
	avsc.mutation.SetProductID(i)
	return avsc
}

// SetVariant sets the "variant" edge to the AttributeVariantString entity.
func (avsc *AttributeValueStringCreate) SetVariant(a *AttributeVariantString) *AttributeValueStringCreate {
	return avsc.SetVariantID(a.ID)
}

// SetProduct sets the "product" edge to the Product entity.
func (avsc *AttributeValueStringCreate) SetProduct(p *Product) *AttributeValueStringCreate {
	return avsc.SetProductID(p.ID)
}

// Mutation returns the AttributeValueStringMutation object of the builder.
func (avsc *AttributeValueStringCreate) Mutation() *AttributeValueStringMutation {
	return avsc.mutation
}

// Save creates the AttributeValueString in the database.
func (avsc *AttributeValueStringCreate) Save(ctx context.Context) (*AttributeValueString, error) {
	return withHooks(ctx, avsc.sqlSave, avsc.mutation, avsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (avsc *AttributeValueStringCreate) SaveX(ctx context.Context) *AttributeValueString {
	v, err := avsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (avsc *AttributeValueStringCreate) Exec(ctx context.Context) error {
	_, err := avsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (avsc *AttributeValueStringCreate) ExecX(ctx context.Context) {
	if err := avsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (avsc *AttributeValueStringCreate) check() error {
	if _, ok := avsc.mutation.VariantID(); !ok {
		return &ValidationError{Name: "variant_id", err: errors.New(`ent: missing required field "AttributeValueString.variant_id"`)}
	}
	if _, ok := avsc.mutation.ProductID(); !ok {
		return &ValidationError{Name: "product_id", err: errors.New(`ent: missing required field "AttributeValueString.product_id"`)}
	}
	if _, ok := avsc.mutation.VariantID(); !ok {
		return &ValidationError{Name: "variant", err: errors.New(`ent: missing required edge "AttributeValueString.variant"`)}
	}
	if _, ok := avsc.mutation.ProductID(); !ok {
		return &ValidationError{Name: "product", err: errors.New(`ent: missing required edge "AttributeValueString.product"`)}
	}
	return nil
}

func (avsc *AttributeValueStringCreate) sqlSave(ctx context.Context) (*AttributeValueString, error) {
	if err := avsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := avsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, avsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	avsc.mutation.id = &_node.ID
	avsc.mutation.done = true
	return _node, nil
}

func (avsc *AttributeValueStringCreate) createSpec() (*AttributeValueString, *sqlgraph.CreateSpec) {
	var (
		_node = &AttributeValueString{config: avsc.config}
		_spec = sqlgraph.NewCreateSpec(attributevaluestring.Table, sqlgraph.NewFieldSpec(attributevaluestring.FieldID, field.TypeInt))
	)
	if nodes := avsc.mutation.VariantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   attributevaluestring.VariantTable,
			Columns: []string{attributevaluestring.VariantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attributevariantstring.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.VariantID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := avsc.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attributevaluestring.ProductTable,
			Columns: []string{attributevaluestring.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProductID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AttributeValueStringCreateBulk is the builder for creating many AttributeValueString entities in bulk.
type AttributeValueStringCreateBulk struct {
	config
	err      error
	builders []*AttributeValueStringCreate
}

// Save creates the AttributeValueString entities in the database.
func (avscb *AttributeValueStringCreateBulk) Save(ctx context.Context) ([]*AttributeValueString, error) {
	if avscb.err != nil {
		return nil, avscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(avscb.builders))
	nodes := make([]*AttributeValueString, len(avscb.builders))
	mutators := make([]Mutator, len(avscb.builders))
	for i := range avscb.builders {
		func(i int, root context.Context) {
			builder := avscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AttributeValueStringMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, avscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, avscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, avscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (avscb *AttributeValueStringCreateBulk) SaveX(ctx context.Context) []*AttributeValueString {
	v, err := avscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (avscb *AttributeValueStringCreateBulk) Exec(ctx context.Context) error {
	_, err := avscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (avscb *AttributeValueStringCreateBulk) ExecX(ctx context.Context) {
	if err := avscb.Exec(ctx); err != nil {
		panic(err)
	}
}
