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
	"github.com/eesoymilk/health-statistic-api/ent/healthkit"
	"github.com/eesoymilk/health-statistic-api/ent/user"
	"github.com/google/uuid"
)

// HealthKit is the model entity for the HealthKit schema.
type HealthKit struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// StartDate holds the value of the "start_date" field.
	StartDate time.Time `json:"start_date,omitempty"`
	// EndDate holds the value of the "end_date" field.
	EndDate time.Time `json:"end_date,omitempty"`
	// StepCount holds the value of the "step_count" field.
	StepCount float64 `json:"step_count,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HealthKitQuery when eager-loading is set.
	Edges          HealthKitEdges `json:"-"`
	user_healthkit *string
	selectValues   sql.SelectValues
}

// HealthKitEdges holds the relations/edges for other nodes in the graph.
type HealthKitEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HealthKitEdges) UserOrErr() (*User, error) {
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
func (*HealthKit) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case healthkit.FieldStepCount:
			values[i] = new(sql.NullFloat64)
		case healthkit.FieldStartDate, healthkit.FieldEndDate:
			values[i] = new(sql.NullTime)
		case healthkit.FieldID:
			values[i] = new(uuid.UUID)
		case healthkit.ForeignKeys[0]: // user_healthkit
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the HealthKit fields.
func (hk *HealthKit) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case healthkit.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				hk.ID = *value
			}
		case healthkit.FieldStartDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_date", values[i])
			} else if value.Valid {
				hk.StartDate = value.Time
			}
		case healthkit.FieldEndDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_date", values[i])
			} else if value.Valid {
				hk.EndDate = value.Time
			}
		case healthkit.FieldStepCount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field step_count", values[i])
			} else if value.Valid {
				hk.StepCount = value.Float64
			}
		case healthkit.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_healthkit", values[i])
			} else if value.Valid {
				hk.user_healthkit = new(string)
				*hk.user_healthkit = value.String
			}
		default:
			hk.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the HealthKit.
// This includes values selected through modifiers, order, etc.
func (hk *HealthKit) Value(name string) (ent.Value, error) {
	return hk.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the HealthKit entity.
func (hk *HealthKit) QueryUser() *UserQuery {
	return NewHealthKitClient(hk.config).QueryUser(hk)
}

// Update returns a builder for updating this HealthKit.
// Note that you need to call HealthKit.Unwrap() before calling this method if this HealthKit
// was returned from a transaction, and the transaction was committed or rolled back.
func (hk *HealthKit) Update() *HealthKitUpdateOne {
	return NewHealthKitClient(hk.config).UpdateOne(hk)
}

// Unwrap unwraps the HealthKit entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (hk *HealthKit) Unwrap() *HealthKit {
	_tx, ok := hk.config.driver.(*txDriver)
	if !ok {
		panic("ent: HealthKit is not a transactional entity")
	}
	hk.config.driver = _tx.drv
	return hk
}

// String implements the fmt.Stringer.
func (hk *HealthKit) String() string {
	var builder strings.Builder
	builder.WriteString("HealthKit(")
	builder.WriteString(fmt.Sprintf("id=%v, ", hk.ID))
	builder.WriteString("start_date=")
	builder.WriteString(hk.StartDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_date=")
	builder.WriteString(hk.EndDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("step_count=")
	builder.WriteString(fmt.Sprintf("%v", hk.StepCount))
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (hk *HealthKit) MarshalJSON() ([]byte, error) {
	type Alias HealthKit
	return json.Marshal(&struct {
		*Alias
		HealthKitEdges
	}{
		Alias:          (*Alias)(hk),
		HealthKitEdges: hk.Edges,
	})
}

// HealthKits is a parsable slice of HealthKit.
type HealthKits []*HealthKit