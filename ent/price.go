// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/eesoymilk/health-statistic-api/ent/price"
)

// Price is the model entity for the Price schema.
type Price struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Price) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case price.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Price fields.
func (pr *Price) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case price.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Price.
// This includes values selected through modifiers, order, etc.
func (pr *Price) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// Update returns a builder for updating this Price.
// Note that you need to call Price.Unwrap() before calling this method if this Price
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Price) Update() *PriceUpdateOne {
	return NewPriceClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Price entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Price) Unwrap() *Price {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Price is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Price) String() string {
	var builder strings.Builder
	builder.WriteString("Price(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Prices is a parsable slice of Price.
type Prices []*Price
