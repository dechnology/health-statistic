// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package deegoo

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the deegoo type in the database.
	Label = "deegoo"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPerception holds the string denoting the perception field in the database.
	FieldPerception = "perception"
	// FieldFocus holds the string denoting the focus field in the database.
	FieldFocus = "focus"
	// FieldExecution holds the string denoting the execution field in the database.
	FieldExecution = "execution"
	// FieldMemory holds the string denoting the memory field in the database.
	FieldMemory = "memory"
	// FieldLanguage holds the string denoting the language field in the database.
	FieldLanguage = "language"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the deegoo in the database.
	Table = "deegoos"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "deegoos"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_deegoo"
)

// Columns holds all SQL columns for deegoo fields.
var Columns = []string{
	FieldID,
	FieldPerception,
	FieldFocus,
	FieldExecution,
	FieldMemory,
	FieldLanguage,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "deegoos"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_deegoo",
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
	// PerceptionValidator is a validator for the "perception" field. It is called by the builders before save.
	PerceptionValidator func(int8) error
	// FocusValidator is a validator for the "focus" field. It is called by the builders before save.
	FocusValidator func(int8) error
	// ExecutionValidator is a validator for the "execution" field. It is called by the builders before save.
	ExecutionValidator func(int8) error
	// MemoryValidator is a validator for the "memory" field. It is called by the builders before save.
	MemoryValidator func(int8) error
	// LanguageValidator is a validator for the "language" field. It is called by the builders before save.
	LanguageValidator func(int8) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Deegoo queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPerception orders the results by the perception field.
func ByPerception(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPerception, opts...).ToFunc()
}

// ByFocus orders the results by the focus field.
func ByFocus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFocus, opts...).ToFunc()
}

// ByExecution orders the results by the execution field.
func ByExecution(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExecution, opts...).ToFunc()
}

// ByMemory orders the results by the memory field.
func ByMemory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemory, opts...).ToFunc()
}

// ByLanguage orders the results by the language field.
func ByLanguage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLanguage, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
