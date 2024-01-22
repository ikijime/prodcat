// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"prodcat/ent/attributevaluenum"
	"prodcat/ent/attributevariantnum"
	"prodcat/ent/product"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// AttributeValueNum is the model entity for the AttributeValueNum schema.
type AttributeValueNum struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// VariantID holds the value of the "variant_id" field.
	VariantID int `json:"variant_id,omitempty"`
	// ProductID holds the value of the "product_id" field.
	ProductID int `json:"product_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AttributeValueNumQuery when eager-loading is set.
	Edges        AttributeValueNumEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AttributeValueNumEdges holds the relations/edges for other nodes in the graph.
type AttributeValueNumEdges struct {
	// Variant holds the value of the variant edge.
	Variant *AttributeVariantNum `json:"variant,omitempty"`
	// Product holds the value of the product edge.
	Product *Product `json:"product,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// VariantOrErr returns the Variant value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AttributeValueNumEdges) VariantOrErr() (*AttributeVariantNum, error) {
	if e.loadedTypes[0] {
		if e.Variant == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: attributevariantnum.Label}
		}
		return e.Variant, nil
	}
	return nil, &NotLoadedError{edge: "variant"}
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AttributeValueNumEdges) ProductOrErr() (*Product, error) {
	if e.loadedTypes[1] {
		if e.Product == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: product.Label}
		}
		return e.Product, nil
	}
	return nil, &NotLoadedError{edge: "product"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AttributeValueNum) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case attributevaluenum.FieldID, attributevaluenum.FieldVariantID, attributevaluenum.FieldProductID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AttributeValueNum fields.
func (avn *AttributeValueNum) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case attributevaluenum.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			avn.ID = int(value.Int64)
		case attributevaluenum.FieldVariantID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field variant_id", values[i])
			} else if value.Valid {
				avn.VariantID = int(value.Int64)
			}
		case attributevaluenum.FieldProductID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value.Valid {
				avn.ProductID = int(value.Int64)
			}
		default:
			avn.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AttributeValueNum.
// This includes values selected through modifiers, order, etc.
func (avn *AttributeValueNum) Value(name string) (ent.Value, error) {
	return avn.selectValues.Get(name)
}

// QueryVariant queries the "variant" edge of the AttributeValueNum entity.
func (avn *AttributeValueNum) QueryVariant() *AttributeVariantNumQuery {
	return NewAttributeValueNumClient(avn.config).QueryVariant(avn)
}

// QueryProduct queries the "product" edge of the AttributeValueNum entity.
func (avn *AttributeValueNum) QueryProduct() *ProductQuery {
	return NewAttributeValueNumClient(avn.config).QueryProduct(avn)
}

// Update returns a builder for updating this AttributeValueNum.
// Note that you need to call AttributeValueNum.Unwrap() before calling this method if this AttributeValueNum
// was returned from a transaction, and the transaction was committed or rolled back.
func (avn *AttributeValueNum) Update() *AttributeValueNumUpdateOne {
	return NewAttributeValueNumClient(avn.config).UpdateOne(avn)
}

// Unwrap unwraps the AttributeValueNum entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (avn *AttributeValueNum) Unwrap() *AttributeValueNum {
	_tx, ok := avn.config.driver.(*txDriver)
	if !ok {
		panic("ent: AttributeValueNum is not a transactional entity")
	}
	avn.config.driver = _tx.drv
	return avn
}

// String implements the fmt.Stringer.
func (avn *AttributeValueNum) String() string {
	var builder strings.Builder
	builder.WriteString("AttributeValueNum(")
	builder.WriteString(fmt.Sprintf("id=%v, ", avn.ID))
	builder.WriteString("variant_id=")
	builder.WriteString(fmt.Sprintf("%v", avn.VariantID))
	builder.WriteString(", ")
	builder.WriteString("product_id=")
	builder.WriteString(fmt.Sprintf("%v", avn.ProductID))
	builder.WriteByte(')')
	return builder.String()
}

// AttributeValueNums is a parsable slice of AttributeValueNum.
type AttributeValueNums []*AttributeValueNum