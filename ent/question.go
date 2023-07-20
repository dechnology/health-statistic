// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/eesoymilk/health-statistic-api/ent/question"
	"github.com/eesoymilk/health-statistic-api/ent/questionnaire"
	"github.com/google/uuid"
)

// Question is the model entity for the Question schema.
type Question struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Body holds the value of the "body" field.
	Body string `json:"body,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the QuestionQuery when eager-loading is set.
	Edges                   QuestionEdges `json:"-"`
	questionnaire_questions *uuid.UUID
	selectValues            sql.SelectValues
}

// QuestionEdges holds the relations/edges for other nodes in the graph.
type QuestionEdges struct {
	// Questionnaire holds the value of the questionnaire edge.
	Questionnaire *Questionnaire `json:"questionnaire,omitempty"`
	// Answers holds the value of the answers edge.
	Answers []*Answer `json:"answers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// QuestionnaireOrErr returns the Questionnaire value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e QuestionEdges) QuestionnaireOrErr() (*Questionnaire, error) {
	if e.loadedTypes[0] {
		if e.Questionnaire == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: questionnaire.Label}
		}
		return e.Questionnaire, nil
	}
	return nil, &NotLoadedError{edge: "questionnaire"}
}

// AnswersOrErr returns the Answers value or an error if the edge
// was not loaded in eager-loading.
func (e QuestionEdges) AnswersOrErr() ([]*Answer, error) {
	if e.loadedTypes[1] {
		return e.Answers, nil
	}
	return nil, &NotLoadedError{edge: "answers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Question) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case question.FieldBody, question.FieldType:
			values[i] = new(sql.NullString)
		case question.FieldID:
			values[i] = new(uuid.UUID)
		case question.ForeignKeys[0]: // questionnaire_questions
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Question fields.
func (q *Question) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case question.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				q.ID = *value
			}
		case question.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				q.Body = value.String
			}
		case question.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				q.Type = value.String
			}
		case question.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field questionnaire_questions", values[i])
			} else if value.Valid {
				q.questionnaire_questions = new(uuid.UUID)
				*q.questionnaire_questions = *value.S.(*uuid.UUID)
			}
		default:
			q.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Question.
// This includes values selected through modifiers, order, etc.
func (q *Question) Value(name string) (ent.Value, error) {
	return q.selectValues.Get(name)
}

// QueryQuestionnaire queries the "questionnaire" edge of the Question entity.
func (q *Question) QueryQuestionnaire() *QuestionnaireQuery {
	return NewQuestionClient(q.config).QueryQuestionnaire(q)
}

// QueryAnswers queries the "answers" edge of the Question entity.
func (q *Question) QueryAnswers() *AnswerQuery {
	return NewQuestionClient(q.config).QueryAnswers(q)
}

// Update returns a builder for updating this Question.
// Note that you need to call Question.Unwrap() before calling this method if this Question
// was returned from a transaction, and the transaction was committed or rolled back.
func (q *Question) Update() *QuestionUpdateOne {
	return NewQuestionClient(q.config).UpdateOne(q)
}

// Unwrap unwraps the Question entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (q *Question) Unwrap() *Question {
	_tx, ok := q.config.driver.(*txDriver)
	if !ok {
		panic("ent: Question is not a transactional entity")
	}
	q.config.driver = _tx.drv
	return q
}

// String implements the fmt.Stringer.
func (q *Question) String() string {
	var builder strings.Builder
	builder.WriteString("Question(")
	builder.WriteString(fmt.Sprintf("id=%v, ", q.ID))
	builder.WriteString("body=")
	builder.WriteString(q.Body)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(q.Type)
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (q *Question) MarshalJSON() ([]byte, error) {
	type Alias Question
	return json.Marshal(&struct {
		*Alias
		QuestionEdges
	}{
		Alias:         (*Alias)(q),
		QuestionEdges: q.Edges,
	})
}

// Questions is a parsable slice of Question.
type Questions []*Question