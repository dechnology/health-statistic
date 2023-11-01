// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package hkdata

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the hkdata type in the database.
	Label = "hk_data"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDataID holds the string denoting the data_id field in the database.
	FieldDataID = "data_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldStartTimestamp holds the string denoting the start_timestamp field in the database.
	FieldStartTimestamp = "start_timestamp"
	// FieldEndTimestamp holds the string denoting the end_timestamp field in the database.
	FieldEndTimestamp = "end_timestamp"
	// FieldTimezoneID holds the string denoting the timezone_id field in the database.
	FieldTimezoneID = "timezone_id"
	// EdgeHealthkit holds the string denoting the healthkit edge name in mutations.
	EdgeHealthkit = "healthkit"
	// Table holds the table name of the hkdata in the database.
	Table = "hk_data"
	// HealthkitTable is the table that holds the healthkit relation/edge.
	HealthkitTable = "hk_data"
	// HealthkitInverseTable is the table name for the HealthKit entity.
	// It exists in this package in order to avoid circular dependency with the "healthkit" package.
	HealthkitInverseTable = "health_kits"
	// HealthkitColumn is the table column denoting the healthkit relation/edge.
	HealthkitColumn = "health_kit_data"
)

// Columns holds all SQL columns for hkdata fields.
var Columns = []string{
	FieldID,
	FieldDataID,
	FieldType,
	FieldValue,
	FieldStartTimestamp,
	FieldEndTimestamp,
	FieldTimezoneID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "hk_data"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"health_kit_data",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the HKData queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDataID orders the results by the data_id field.
func ByDataID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDataID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByValue orders the results by the value field.
func ByValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValue, opts...).ToFunc()
}

// ByStartTimestamp orders the results by the start_timestamp field.
func ByStartTimestamp(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartTimestamp, opts...).ToFunc()
}

// ByEndTimestamp orders the results by the end_timestamp field.
func ByEndTimestamp(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndTimestamp, opts...).ToFunc()
}

// ByTimezoneID orders the results by the timezone_id field.
func ByTimezoneID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimezoneID, opts...).ToFunc()
}

// ByHealthkitField orders the results by healthkit field.
func ByHealthkitField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHealthkitStep(), sql.OrderByField(field, opts...))
	}
}
func newHealthkitStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HealthkitInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, HealthkitTable, HealthkitColumn),
	)
}
