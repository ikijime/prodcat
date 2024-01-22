// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"prodcat/ent/schema"
	"prodcat/ent/usersettings"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UserSettings is the model entity for the UserSettings schema.
type UserSettings struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Frontend holds the value of the "frontend" field.
	Frontend     schema.DefaultSetting `json:"frontend,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserSettings) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case usersettings.FieldFrontend:
			values[i] = new([]byte)
		case usersettings.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserSettings fields.
func (us *UserSettings) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usersettings.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			us.ID = int(value.Int64)
		case usersettings.FieldFrontend:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field frontend", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &us.Frontend); err != nil {
					return fmt.Errorf("unmarshal field frontend: %w", err)
				}
			}
		default:
			us.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserSettings.
// This includes values selected through modifiers, order, etc.
func (us *UserSettings) Value(name string) (ent.Value, error) {
	return us.selectValues.Get(name)
}

// Update returns a builder for updating this UserSettings.
// Note that you need to call UserSettings.Unwrap() before calling this method if this UserSettings
// was returned from a transaction, and the transaction was committed or rolled back.
func (us *UserSettings) Update() *UserSettingsUpdateOne {
	return NewUserSettingsClient(us.config).UpdateOne(us)
}

// Unwrap unwraps the UserSettings entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (us *UserSettings) Unwrap() *UserSettings {
	_tx, ok := us.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserSettings is not a transactional entity")
	}
	us.config.driver = _tx.drv
	return us
}

// String implements the fmt.Stringer.
func (us *UserSettings) String() string {
	var builder strings.Builder
	builder.WriteString("UserSettings(")
	builder.WriteString(fmt.Sprintf("id=%v, ", us.ID))
	builder.WriteString("frontend=")
	builder.WriteString(fmt.Sprintf("%v", us.Frontend))
	builder.WriteByte(')')
	return builder.String()
}

// UserSettingsSlice is a parsable slice of UserSettings.
type UserSettingsSlice []*UserSettings