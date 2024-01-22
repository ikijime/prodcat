// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"prodcat/ent/brand"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Brand is the model entity for the Brand schema.
type Brand struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BrandQuery when eager-loading is set.
	Edges        BrandEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BrandEdges holds the relations/edges for other nodes in the graph.
type BrandEdges struct {
	// Product holds the value of the product edge.
	Product []*Product `json:"product,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading.
func (e BrandEdges) ProductOrErr() ([]*Product, error) {
	if e.loadedTypes[0] {
		return e.Product, nil
	}
	return nil, &NotLoadedError{edge: "product"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Brand) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case brand.FieldID:
			values[i] = new(sql.NullInt64)
		case brand.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Brand fields.
func (b *Brand) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case brand.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case brand.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				b.Name = value.String
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Brand.
// This includes values selected through modifiers, order, etc.
func (b *Brand) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QueryProduct queries the "product" edge of the Brand entity.
func (b *Brand) QueryProduct() *ProductQuery {
	return NewBrandClient(b.config).QueryProduct(b)
}

// Update returns a builder for updating this Brand.
// Note that you need to call Brand.Unwrap() before calling this method if this Brand
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Brand) Update() *BrandUpdateOne {
	return NewBrandClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Brand entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Brand) Unwrap() *Brand {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Brand is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Brand) String() string {
	var builder strings.Builder
	builder.WriteString("Brand(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("name=")
	builder.WriteString(b.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Brands is a parsable slice of Brand.
type Brands []*Brand
