// Code generated by ent, DO NOT EDIT.

package model

import (
	"fmt"
	"mall-go/module/app/service/internal/data/model/theme"
	"mall-go/module/app/service/internal/data/model/themespu"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// ThemeSpu is the model entity for the ThemeSpu schema.
type ThemeSpu struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// ThemeID holds the value of the "theme_id" field.
	ThemeID int64 `json:"theme_id,omitempty"`
	// SpuID holds the value of the "spu_id" field.
	SpuID int64 `json:"spu_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ThemeSpuQuery when eager-loading is set.
	Edges ThemeSpuEdges `json:"edges"`
}

// ThemeSpuEdges holds the relations/edges for other nodes in the graph.
type ThemeSpuEdges struct {
	// Theme holds the value of the theme edge.
	Theme *Theme `json:"theme,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ThemeOrErr returns the Theme value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ThemeSpuEdges) ThemeOrErr() (*Theme, error) {
	if e.loadedTypes[0] {
		if e.Theme == nil {
			// The edge theme was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: theme.Label}
		}
		return e.Theme, nil
	}
	return nil, &NotLoadedError{edge: "theme"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ThemeSpu) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case themespu.FieldID, themespu.FieldThemeID, themespu.FieldSpuID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ThemeSpu", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ThemeSpu fields.
func (ts *ThemeSpu) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case themespu.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ts.ID = int64(value.Int64)
		case themespu.FieldThemeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field theme_id", values[i])
			} else if value.Valid {
				ts.ThemeID = value.Int64
			}
		case themespu.FieldSpuID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field spu_id", values[i])
			} else if value.Valid {
				ts.SpuID = value.Int64
			}
		}
	}
	return nil
}

// QueryTheme queries the "theme" edge of the ThemeSpu entity.
func (ts *ThemeSpu) QueryTheme() *ThemeQuery {
	return (&ThemeSpuClient{config: ts.config}).QueryTheme(ts)
}

// Update returns a builder for updating this ThemeSpu.
// Note that you need to call ThemeSpu.Unwrap() before calling this method if this ThemeSpu
// was returned from a transaction, and the transaction was committed or rolled back.
func (ts *ThemeSpu) Update() *ThemeSpuUpdateOne {
	return (&ThemeSpuClient{config: ts.config}).UpdateOne(ts)
}

// Unwrap unwraps the ThemeSpu entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ts *ThemeSpu) Unwrap() *ThemeSpu {
	_tx, ok := ts.config.driver.(*txDriver)
	if !ok {
		panic("model: ThemeSpu is not a transactional entity")
	}
	ts.config.driver = _tx.drv
	return ts
}

// String implements the fmt.Stringer.
func (ts *ThemeSpu) String() string {
	var builder strings.Builder
	builder.WriteString("ThemeSpu(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ts.ID))
	builder.WriteString("theme_id=")
	builder.WriteString(fmt.Sprintf("%v", ts.ThemeID))
	builder.WriteString(", ")
	builder.WriteString("spu_id=")
	builder.WriteString(fmt.Sprintf("%v", ts.SpuID))
	builder.WriteByte(')')
	return builder.String()
}

// ThemeSpus is a parsable slice of ThemeSpu.
type ThemeSpus []*ThemeSpu

func (ts ThemeSpus) config(cfg config) {
	for _i := range ts {
		ts[_i].config = cfg
	}
}
