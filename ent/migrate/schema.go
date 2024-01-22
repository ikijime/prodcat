// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AttributesColumns holds the columns for the "attributes" table.
	AttributesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
	}
	// AttributesTable holds the schema information for the "attributes" table.
	AttributesTable = &schema.Table{
		Name:       "attributes",
		Columns:    AttributesColumns,
		PrimaryKey: []*schema.Column{AttributesColumns[0]},
	}
	// AttributeValueBoolsColumns holds the columns for the "attribute_value_bools" table.
	AttributeValueBoolsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "value", Type: field.TypeBool},
		{Name: "attribute_id", Type: field.TypeInt},
		{Name: "product_id", Type: field.TypeInt},
	}
	// AttributeValueBoolsTable holds the schema information for the "attribute_value_bools" table.
	AttributeValueBoolsTable = &schema.Table{
		Name:       "attribute_value_bools",
		Columns:    AttributeValueBoolsColumns,
		PrimaryKey: []*schema.Column{AttributeValueBoolsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attribute_value_bools_attributes_attributeValuesBool",
				Columns:    []*schema.Column{AttributeValueBoolsColumns[2]},
				RefColumns: []*schema.Column{AttributesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "attribute_value_bools_products_attributeValuesBool",
				Columns:    []*schema.Column{AttributeValueBoolsColumns[3]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// AttributeValueNumsColumns holds the columns for the "attribute_value_nums" table.
	AttributeValueNumsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "variant_id", Type: field.TypeInt},
		{Name: "product_id", Type: field.TypeInt},
	}
	// AttributeValueNumsTable holds the schema information for the "attribute_value_nums" table.
	AttributeValueNumsTable = &schema.Table{
		Name:       "attribute_value_nums",
		Columns:    AttributeValueNumsColumns,
		PrimaryKey: []*schema.Column{AttributeValueNumsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attribute_value_nums_attribute_variant_nums_variant",
				Columns:    []*schema.Column{AttributeValueNumsColumns[1]},
				RefColumns: []*schema.Column{AttributeVariantNumsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "attribute_value_nums_products_attributeValuesNum",
				Columns:    []*schema.Column{AttributeValueNumsColumns[2]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// AttributeValueStringsColumns holds the columns for the "attribute_value_strings" table.
	AttributeValueStringsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "variant_id", Type: field.TypeInt},
		{Name: "product_id", Type: field.TypeInt},
	}
	// AttributeValueStringsTable holds the schema information for the "attribute_value_strings" table.
	AttributeValueStringsTable = &schema.Table{
		Name:       "attribute_value_strings",
		Columns:    AttributeValueStringsColumns,
		PrimaryKey: []*schema.Column{AttributeValueStringsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attribute_value_strings_attribute_variant_strings_variant",
				Columns:    []*schema.Column{AttributeValueStringsColumns[1]},
				RefColumns: []*schema.Column{AttributeVariantStringsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "attribute_value_strings_products_attributeValuesString",
				Columns:    []*schema.Column{AttributeValueStringsColumns[2]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// AttributeVariantNumsColumns holds the columns for the "attribute_variant_nums" table.
	AttributeVariantNumsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "value", Type: field.TypeInt},
		{Name: "attribute_id", Type: field.TypeInt},
	}
	// AttributeVariantNumsTable holds the schema information for the "attribute_variant_nums" table.
	AttributeVariantNumsTable = &schema.Table{
		Name:       "attribute_variant_nums",
		Columns:    AttributeVariantNumsColumns,
		PrimaryKey: []*schema.Column{AttributeVariantNumsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attribute_variant_nums_attributes_attributeVariantsNum",
				Columns:    []*schema.Column{AttributeVariantNumsColumns[2]},
				RefColumns: []*schema.Column{AttributesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// AttributeVariantStringsColumns holds the columns for the "attribute_variant_strings" table.
	AttributeVariantStringsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "value", Type: field.TypeString, Unique: true},
		{Name: "attribute_id", Type: field.TypeInt},
	}
	// AttributeVariantStringsTable holds the schema information for the "attribute_variant_strings" table.
	AttributeVariantStringsTable = &schema.Table{
		Name:       "attribute_variant_strings",
		Columns:    AttributeVariantStringsColumns,
		PrimaryKey: []*schema.Column{AttributeVariantStringsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attribute_variant_strings_attributes_attributeVariantsString",
				Columns:    []*schema.Column{AttributeVariantStringsColumns[2]},
				RefColumns: []*schema.Column{AttributesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// BrandsColumns holds the columns for the "brands" table.
	BrandsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// BrandsTable holds the schema information for the "brands" table.
	BrandsTable = &schema.Table{
		Name:       "brands",
		Columns:    BrandsColumns,
		PrimaryKey: []*schema.Column{BrandsColumns[0]},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "code", Type: field.TypeInt, Unique: true},
		{Name: "barcode", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "brand_product", Type: field.TypeInt, Nullable: true},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "products_brands_product",
				Columns:    []*schema.Column{ProductsColumns[7]},
				RefColumns: []*schema.Column{BrandsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "login", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "first_name", Type: field.TypeString, Nullable: true},
		{Name: "last_name", Type: field.TypeString, Nullable: true},
		{Name: "role", Type: field.TypeString},
		{Name: "phonenumber", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "email", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "token", Type: field.TypeString, Default: ""},
		{Name: "refresh_token", Type: field.TypeString, Default: ""},
		{Name: "user_settings", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_user_settings_settings",
				Columns:    []*schema.Column{UsersColumns[13]},
				RefColumns: []*schema.Column{UserSettingsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserSettingsColumns holds the columns for the "user_settings" table.
	UserSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "frontend", Type: field.TypeJSON},
	}
	// UserSettingsTable holds the schema information for the "user_settings" table.
	UserSettingsTable = &schema.Table{
		Name:       "user_settings",
		Columns:    UserSettingsColumns,
		PrimaryKey: []*schema.Column{UserSettingsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AttributesTable,
		AttributeValueBoolsTable,
		AttributeValueNumsTable,
		AttributeValueStringsTable,
		AttributeVariantNumsTable,
		AttributeVariantStringsTable,
		BrandsTable,
		ProductsTable,
		UsersTable,
		UserSettingsTable,
	}
)

func init() {
	AttributeValueBoolsTable.ForeignKeys[0].RefTable = AttributesTable
	AttributeValueBoolsTable.ForeignKeys[1].RefTable = ProductsTable
	AttributeValueNumsTable.ForeignKeys[0].RefTable = AttributeVariantNumsTable
	AttributeValueNumsTable.ForeignKeys[1].RefTable = ProductsTable
	AttributeValueStringsTable.ForeignKeys[0].RefTable = AttributeVariantStringsTable
	AttributeValueStringsTable.ForeignKeys[1].RefTable = ProductsTable
	AttributeVariantNumsTable.ForeignKeys[0].RefTable = AttributesTable
	AttributeVariantStringsTable.ForeignKeys[0].RefTable = AttributesTable
	ProductsTable.ForeignKeys[0].RefTable = BrandsTable
	UsersTable.ForeignKeys[0].RefTable = UserSettingsTable
}