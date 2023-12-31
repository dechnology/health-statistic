// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/eesoymilk/health-statistic-api/ent/deegoo"
	"github.com/eesoymilk/health-statistic-api/ent/user"
	"github.com/google/uuid"
)

// Deegoo is the model entity for the Deegoo schema.
type Deegoo struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Perception holds the value of the "perception" field.
	Perception int8 `json:"perception"`
	// Focus holds the value of the "focus" field.
	Focus int8 `json:"focus"`
	// Execution holds the value of the "execution" field.
	Execution int8 `json:"execution"`
	// Memory holds the value of the "memory" field.
	Memory int8 `json:"memory"`
	// Language holds the value of the "language" field.
	Language int8 `json:"language"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeegooQuery when eager-loading is set.
	Edges        DeegooEdges `json:"-"`
	user_deegoo  *string
	selectValues sql.SelectValues
}

// DeegooEdges holds the relations/edges for other nodes in the graph.
type DeegooEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeegooEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Deegoo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case deegoo.FieldPerception, deegoo.FieldFocus, deegoo.FieldExecution, deegoo.FieldMemory, deegoo.FieldLanguage:
			values[i] = new(sql.NullInt64)
		case deegoo.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case deegoo.FieldID:
			values[i] = new(uuid.UUID)
		case deegoo.ForeignKeys[0]: // user_deegoo
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Deegoo fields.
func (d *Deegoo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case deegoo.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				d.ID = *value
			}
		case deegoo.FieldPerception:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field perception", values[i])
			} else if value.Valid {
				d.Perception = int8(value.Int64)
			}
		case deegoo.FieldFocus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field focus", values[i])
			} else if value.Valid {
				d.Focus = int8(value.Int64)
			}
		case deegoo.FieldExecution:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field execution", values[i])
			} else if value.Valid {
				d.Execution = int8(value.Int64)
			}
		case deegoo.FieldMemory:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field memory", values[i])
			} else if value.Valid {
				d.Memory = int8(value.Int64)
			}
		case deegoo.FieldLanguage:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field language", values[i])
			} else if value.Valid {
				d.Language = int8(value.Int64)
			}
		case deegoo.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				d.CreatedAt = value.Time
			}
		case deegoo.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_deegoo", values[i])
			} else if value.Valid {
				d.user_deegoo = new(string)
				*d.user_deegoo = value.String
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Deegoo.
// This includes values selected through modifiers, order, etc.
func (d *Deegoo) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Deegoo entity.
func (d *Deegoo) QueryUser() *UserQuery {
	return NewDeegooClient(d.config).QueryUser(d)
}

// Update returns a builder for updating this Deegoo.
// Note that you need to call Deegoo.Unwrap() before calling this method if this Deegoo
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Deegoo) Update() *DeegooUpdateOne {
	return NewDeegooClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Deegoo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Deegoo) Unwrap() *Deegoo {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Deegoo is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Deegoo) String() string {
	var builder strings.Builder
	builder.WriteString("Deegoo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("perception=")
	builder.WriteString(fmt.Sprintf("%v", d.Perception))
	builder.WriteString(", ")
	builder.WriteString("focus=")
	builder.WriteString(fmt.Sprintf("%v", d.Focus))
	builder.WriteString(", ")
	builder.WriteString("execution=")
	builder.WriteString(fmt.Sprintf("%v", d.Execution))
	builder.WriteString(", ")
	builder.WriteString("memory=")
	builder.WriteString(fmt.Sprintf("%v", d.Memory))
	builder.WriteString(", ")
	builder.WriteString("language=")
	builder.WriteString(fmt.Sprintf("%v", d.Language))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(d.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (d *Deegoo) MarshalJSON() ([]byte, error) {
	type Alias Deegoo
	return json.Marshal(&struct {
		*Alias
		DeegooEdges
	}{
		Alias:       (*Alias)(d),
		DeegooEdges: d.Edges,
	})
}

// Deegoos is a parsable slice of Deegoo.
type Deegoos []*Deegoo
