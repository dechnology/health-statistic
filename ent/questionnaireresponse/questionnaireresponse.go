// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package questionnaireresponse

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the questionnaireresponse type in the database.
	Label = "questionnaire_response"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeQuestionnaire holds the string denoting the questionnaire edge name in mutations.
	EdgeQuestionnaire = "questionnaire"
	// EdgeAnswers holds the string denoting the answers edge name in mutations.
	EdgeAnswers = "answers"
	// Table holds the table name of the questionnaireresponse in the database.
	Table = "questionnaire_responses"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "questionnaire_responses"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_questionnaire_responses"
	// QuestionnaireTable is the table that holds the questionnaire relation/edge.
	QuestionnaireTable = "questionnaire_responses"
	// QuestionnaireInverseTable is the table name for the Questionnaire entity.
	// It exists in this package in order to avoid circular dependency with the "questionnaire" package.
	QuestionnaireInverseTable = "questionnaires"
	// QuestionnaireColumn is the table column denoting the questionnaire relation/edge.
	QuestionnaireColumn = "questionnaire_questionnaire_responses"
	// AnswersTable is the table that holds the answers relation/edge.
	AnswersTable = "answers"
	// AnswersInverseTable is the table name for the Answer entity.
	// It exists in this package in order to avoid circular dependency with the "answer" package.
	AnswersInverseTable = "answers"
	// AnswersColumn is the table column denoting the answers relation/edge.
	AnswersColumn = "questionnaire_response_answers"
)

// Columns holds all SQL columns for questionnaireresponse fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "questionnaire_responses"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"questionnaire_questionnaire_responses",
	"user_questionnaire_responses",
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the QuestionnaireResponse queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
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

// ByQuestionnaireField orders the results by questionnaire field.
func ByQuestionnaireField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newQuestionnaireStep(), sql.OrderByField(field, opts...))
	}
}

// ByAnswersCount orders the results by answers count.
func ByAnswersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAnswersStep(), opts...)
	}
}

// ByAnswers orders the results by answers terms.
func ByAnswers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAnswersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newQuestionnaireStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(QuestionnaireInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, QuestionnaireTable, QuestionnaireColumn),
	)
}
func newAnswersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AnswersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AnswersTable, AnswersColumn),
	)
}
