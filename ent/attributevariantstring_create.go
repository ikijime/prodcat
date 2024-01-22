// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"prodcat/ent/attribute"
	"prodcat/ent/attributevariantstring"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AttributeVariantStringCreate is the builder for creating a AttributeVariantString entity.
type AttributeVariantStringCreate struct {
	config
	mutation *AttributeVariantStringMutation
	hooks    []Hook
}

// SetAttributeID sets the "attribute_id" field.
func (avsc *AttributeVariantStringCreate) SetAttributeID(i int) *AttributeVariantStringCreate {
	avsc.mutation.SetAttributeID(i)
	return avsc
}

// SetValue sets the "value" field.
func (avsc *AttributeVariantStringCreate) SetValue(s string) *AttributeVariantStringCreate {
	avsc.mutation.SetValue(s)
	return avsc
}

// SetAttribute sets the "attribute" edge to the Attribute entity.
func (avsc *AttributeVariantStringCreate) SetAttribute(a *Attribute) *AttributeVariantStringCreate {
	return avsc.SetAttributeID(a.ID)
}

// Mutation returns the AttributeVariantStringMutation object of the builder.
func (avsc *AttributeVariantStringCreate) Mutation() *AttributeVariantStringMutation {
	return avsc.mutation
}

// Save creates the AttributeVariantString in the database.
func (avsc *AttributeVariantStringCreate) Save(ctx context.Context) (*AttributeVariantString, error) {
	return withHooks(ctx, avsc.sqlSave, avsc.mutation, avsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (avsc *AttributeVariantStringCreate) SaveX(ctx context.Context) *AttributeVariantString {
	v, err := avsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (avsc *AttributeVariantStringCreate) Exec(ctx context.Context) error {
	_, err := avsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (avsc *AttributeVariantStringCreate) ExecX(ctx context.Context) {
	if err := avsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (avsc *AttributeVariantStringCreate) check() error {
	if _, ok := avsc.mutation.AttributeID(); !ok {
		return &ValidationError{Name: "attribute_id", err: errors.New(`ent: missing required field "AttributeVariantString.attribute_id"`)}
	}
	if _, ok := avsc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "AttributeVariantString.value"`)}
	}
	if v, ok := avsc.mutation.Value(); ok {
		if err := attributevariantstring.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "AttributeVariantString.value": %w`, err)}
		}
	}
	if _, ok := avsc.mutation.AttributeID(); !ok {
		return &ValidationError{Name: "attribute", err: errors.New(`ent: missing required edge "AttributeVariantString.attribute"`)}
	}
	return nil
}

func (avsc *AttributeVariantStringCreate) sqlSave(ctx context.Context) (*AttributeVariantString, error) {
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

func (avsc *AttributeVariantStringCreate) createSpec() (*AttributeVariantString, *sqlgraph.CreateSpec) {
	var (
		_node = &AttributeVariantString{config: avsc.config}
		_spec = sqlgraph.NewCreateSpec(attributevariantstring.Table, sqlgraph.NewFieldSpec(attributevariantstring.FieldID, field.TypeInt))
	)
	if value, ok := avsc.mutation.Value(); ok {
		_spec.SetField(attributevariantstring.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	if nodes := avsc.mutation.AttributeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attributevariantstring.AttributeTable,
			Columns: []string{attributevariantstring.AttributeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attribute.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AttributeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AttributeVariantStringCreateBulk is the builder for creating many AttributeVariantString entities in bulk.
type AttributeVariantStringCreateBulk struct {
	config
	err      error
	builders []*AttributeVariantStringCreate
}

// Save creates the AttributeVariantString entities in the database.
func (avscb *AttributeVariantStringCreateBulk) Save(ctx context.Context) ([]*AttributeVariantString, error) {
	if avscb.err != nil {
		return nil, avscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(avscb.builders))
	nodes := make([]*AttributeVariantString, len(avscb.builders))
	mutators := make([]Mutator, len(avscb.builders))
	for i := range avscb.builders {
		func(i int, root context.Context) {
			builder := avscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AttributeVariantStringMutation)
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
func (avscb *AttributeVariantStringCreateBulk) SaveX(ctx context.Context) []*AttributeVariantString {
	v, err := avscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (avscb *AttributeVariantStringCreateBulk) Exec(ctx context.Context) error {
	_, err := avscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (avscb *AttributeVariantStringCreateBulk) ExecX(ctx context.Context) {
	if err := avscb.Exec(ctx); err != nil {
		panic(err)
	}
}
