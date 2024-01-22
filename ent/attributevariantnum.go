// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"prodcat/ent/attribute"
	"prodcat/ent/attributevariantnum"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// AttributeVariantNum is the model entity for the AttributeVariantNum schema.
type AttributeVariantNum struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// AttributeID holds the value of the "attribute_id" field.
	AttributeID int `json:"attribute_id,omitempty"`
	// Value holds the value of the "value" field.
	Value int `json:"value,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AttributeVariantNumQuery when eager-loading is set.
	Edges        AttributeVariantNumEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AttributeVariantNumEdges holds the relations/edges for other nodes in the graph.
type AttributeVariantNumEdges struct {
	// Attribute holds the value of the attribute edge.
	Attribute *Attribute `json:"attribute,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AttributeOrErr returns the Attribute value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AttributeVariantNumEdges) AttributeOrErr() (*Attribute, error) {
	if e.loadedTypes[0] {
		if e.Attribute == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: attribute.Label}
		}
		return e.Attribute, nil
	}
	return nil, &NotLoadedError{edge: "attribute"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AttributeVariantNum) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case attributevariantnum.FieldID, attributevariantnum.FieldAttributeID, attributevariantnum.FieldValue:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AttributeVariantNum fields.
func (avn *AttributeVariantNum) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case attributevariantnum.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			avn.ID = int(value.Int64)
		case attributevariantnum.FieldAttributeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field attribute_id", values[i])
			} else if value.Valid {
				avn.AttributeID = int(value.Int64)
			}
		case attributevariantnum.FieldValue:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				avn.Value = int(value.Int64)
			}
		default:
			avn.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the AttributeVariantNum.
// This includes values selected through modifiers, order, etc.
func (avn *AttributeVariantNum) GetValue(name string) (ent.Value, error) {
	return avn.selectValues.Get(name)
}

// QueryAttribute queries the "attribute" edge of the AttributeVariantNum entity.
func (avn *AttributeVariantNum) QueryAttribute() *AttributeQuery {
	return NewAttributeVariantNumClient(avn.config).QueryAttribute(avn)
}

// Update returns a builder for updating this AttributeVariantNum.
// Note that you need to call AttributeVariantNum.Unwrap() before calling this method if this AttributeVariantNum
// was returned from a transaction, and the transaction was committed or rolled back.
func (avn *AttributeVariantNum) Update() *AttributeVariantNumUpdateOne {
	return NewAttributeVariantNumClient(avn.config).UpdateOne(avn)
}

// Unwrap unwraps the AttributeVariantNum entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (avn *AttributeVariantNum) Unwrap() *AttributeVariantNum {
	_tx, ok := avn.config.driver.(*txDriver)
	if !ok {
		panic("ent: AttributeVariantNum is not a transactional entity")
	}
	avn.config.driver = _tx.drv
	return avn
}

// String implements the fmt.Stringer.
func (avn *AttributeVariantNum) String() string {
	var builder strings.Builder
	builder.WriteString("AttributeVariantNum(")
	builder.WriteString(fmt.Sprintf("id=%v, ", avn.ID))
	builder.WriteString("attribute_id=")
	builder.WriteString(fmt.Sprintf("%v", avn.AttributeID))
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(fmt.Sprintf("%v", avn.Value))
	builder.WriteByte(')')
	return builder.String()
}

// AttributeVariantNums is a parsable slice of AttributeVariantNum.
type AttributeVariantNums []*AttributeVariantNum
