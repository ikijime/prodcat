// Code generated by ent, DO NOT EDIT.

package attribute

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the attribute type in the database.
	Label = "attribute"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgeAttributeValuesBool holds the string denoting the attributevaluesbool edge name in mutations.
	EdgeAttributeValuesBool = "attributeValuesBool"
	// EdgeAttributeVariantsString holds the string denoting the attributevariantsstring edge name in mutations.
	EdgeAttributeVariantsString = "attributeVariantsString"
	// EdgeAttributeVariantsNum holds the string denoting the attributevariantsnum edge name in mutations.
	EdgeAttributeVariantsNum = "attributeVariantsNum"
	// Table holds the table name of the attribute in the database.
	Table = "attributes"
	// AttributeValuesBoolTable is the table that holds the attributeValuesBool relation/edge.
	AttributeValuesBoolTable = "attribute_value_bools"
	// AttributeValuesBoolInverseTable is the table name for the AttributeValueBool entity.
	// It exists in this package in order to avoid circular dependency with the "attributevaluebool" package.
	AttributeValuesBoolInverseTable = "attribute_value_bools"
	// AttributeValuesBoolColumn is the table column denoting the attributeValuesBool relation/edge.
	AttributeValuesBoolColumn = "attribute_id"
	// AttributeVariantsStringTable is the table that holds the attributeVariantsString relation/edge.
	AttributeVariantsStringTable = "attribute_variant_strings"
	// AttributeVariantsStringInverseTable is the table name for the AttributeVariantString entity.
	// It exists in this package in order to avoid circular dependency with the "attributevariantstring" package.
	AttributeVariantsStringInverseTable = "attribute_variant_strings"
	// AttributeVariantsStringColumn is the table column denoting the attributeVariantsString relation/edge.
	AttributeVariantsStringColumn = "attribute_id"
	// AttributeVariantsNumTable is the table that holds the attributeVariantsNum relation/edge.
	AttributeVariantsNumTable = "attribute_variant_nums"
	// AttributeVariantsNumInverseTable is the table name for the AttributeVariantNum entity.
	// It exists in this package in order to avoid circular dependency with the "attributevariantnum" package.
	AttributeVariantsNumInverseTable = "attribute_variant_nums"
	// AttributeVariantsNumColumn is the table column denoting the attributeVariantsNum relation/edge.
	AttributeVariantsNumColumn = "attribute_id"
)

// Columns holds all SQL columns for attribute fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldType,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// OrderOption defines the ordering options for the Attribute queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByAttributeValuesBoolCount orders the results by attributeValuesBool count.
func ByAttributeValuesBoolCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAttributeValuesBoolStep(), opts...)
	}
}

// ByAttributeValuesBool orders the results by attributeValuesBool terms.
func ByAttributeValuesBool(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAttributeValuesBoolStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAttributeVariantsStringCount orders the results by attributeVariantsString count.
func ByAttributeVariantsStringCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAttributeVariantsStringStep(), opts...)
	}
}

// ByAttributeVariantsString orders the results by attributeVariantsString terms.
func ByAttributeVariantsString(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAttributeVariantsStringStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAttributeVariantsNumCount orders the results by attributeVariantsNum count.
func ByAttributeVariantsNumCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAttributeVariantsNumStep(), opts...)
	}
}

// ByAttributeVariantsNum orders the results by attributeVariantsNum terms.
func ByAttributeVariantsNum(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAttributeVariantsNumStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAttributeValuesBoolStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AttributeValuesBoolInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AttributeValuesBoolTable, AttributeValuesBoolColumn),
	)
}
func newAttributeVariantsStringStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AttributeVariantsStringInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AttributeVariantsStringTable, AttributeVariantsStringColumn),
	)
}
func newAttributeVariantsNumStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AttributeVariantsNumInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AttributeVariantsNumTable, AttributeVariantsNumColumn),
	)
}
