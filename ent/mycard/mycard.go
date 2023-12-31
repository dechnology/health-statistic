// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package mycard

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the mycard type in the database.
	Label = "my_card"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "card_number"
	// FieldCardPassword holds the string denoting the card_password field in the database.
	FieldCardPassword = "card_password"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldTakenAt holds the string denoting the taken_at field in the database.
	FieldTakenAt = "taken_at"
	// EdgeRecipient holds the string denoting the recipient edge name in mutations.
	EdgeRecipient = "recipient"
	// EdgeNotifications holds the string denoting the notifications edge name in mutations.
	EdgeNotifications = "notifications"
	// UserFieldID holds the string denoting the ID field of the User.
	UserFieldID = "id"
	// NotificationFieldID holds the string denoting the ID field of the Notification.
	NotificationFieldID = "id"
	// Table holds the table name of the mycard in the database.
	Table = "my_cards"
	// RecipientTable is the table that holds the recipient relation/edge.
	RecipientTable = "my_cards"
	// RecipientInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	RecipientInverseTable = "users"
	// RecipientColumn is the table column denoting the recipient relation/edge.
	RecipientColumn = "user_mycards"
	// NotificationsTable is the table that holds the notifications relation/edge.
	NotificationsTable = "notifications"
	// NotificationsInverseTable is the table name for the Notification entity.
	// It exists in this package in order to avoid circular dependency with the "notification" package.
	NotificationsInverseTable = "notifications"
	// NotificationsColumn is the table column denoting the notifications relation/edge.
	NotificationsColumn = "my_card_notifications"
)

// Columns holds all SQL columns for mycard fields.
var Columns = []string{
	FieldID,
	FieldCardPassword,
	FieldCreatedAt,
	FieldTakenAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "my_cards"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_mycards",
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
	// CardPasswordValidator is a validator for the "card_password" field. It is called by the builders before save.
	CardPasswordValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the MyCard queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCardPassword orders the results by the card_password field.
func ByCardPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCardPassword, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByTakenAt orders the results by the taken_at field.
func ByTakenAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTakenAt, opts...).ToFunc()
}

// ByRecipientField orders the results by recipient field.
func ByRecipientField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRecipientStep(), sql.OrderByField(field, opts...))
	}
}

// ByNotificationsCount orders the results by notifications count.
func ByNotificationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNotificationsStep(), opts...)
	}
}

// ByNotifications orders the results by notifications terms.
func ByNotifications(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNotificationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newRecipientStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RecipientInverseTable, UserFieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, RecipientTable, RecipientColumn),
	)
}
func newNotificationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NotificationsInverseTable, NotificationFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, NotificationsTable, NotificationsColumn),
	)
}
