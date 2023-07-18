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
	"github.com/eesoymilk/health-statistic-api/ent/questionnaire"
	"github.com/eesoymilk/health-statistic-api/ent/questionnaireresponse"
	"github.com/eesoymilk/health-statistic-api/ent/user"
	"github.com/google/uuid"
)

// QuestionnaireResponse is the model entity for the QuestionnaireResponse schema.
type QuestionnaireResponse struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the QuestionnaireResponseQuery when eager-loading is set.
	Edges                                 QuestionnaireResponseEdges `json:"-"`
	questionnaire_questionnaire_responses *uuid.UUID
	user_questionnaire_responses          *string
	selectValues                          sql.SelectValues
}

// QuestionnaireResponseEdges holds the relations/edges for other nodes in the graph.
type QuestionnaireResponseEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Questionnaire holds the value of the questionnaire edge.
	Questionnaire *Questionnaire `json:"questionnaire,omitempty"`
	// Answers holds the value of the answers edge.
	Answers []*Answer `json:"answers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e QuestionnaireResponseEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// QuestionnaireOrErr returns the Questionnaire value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e QuestionnaireResponseEdges) QuestionnaireOrErr() (*Questionnaire, error) {
	if e.loadedTypes[1] {
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
func (e QuestionnaireResponseEdges) AnswersOrErr() ([]*Answer, error) {
	if e.loadedTypes[2] {
		return e.Answers, nil
	}
	return nil, &NotLoadedError{edge: "answers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*QuestionnaireResponse) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case questionnaireresponse.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case questionnaireresponse.FieldID:
			values[i] = new(uuid.UUID)
		case questionnaireresponse.ForeignKeys[0]: // questionnaire_questionnaire_responses
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case questionnaireresponse.ForeignKeys[1]: // user_questionnaire_responses
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the QuestionnaireResponse fields.
func (qr *QuestionnaireResponse) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case questionnaireresponse.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				qr.ID = *value
			}
		case questionnaireresponse.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				qr.CreatedAt = value.Time
			}
		case questionnaireresponse.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field questionnaire_questionnaire_responses", values[i])
			} else if value.Valid {
				qr.questionnaire_questionnaire_responses = new(uuid.UUID)
				*qr.questionnaire_questionnaire_responses = *value.S.(*uuid.UUID)
			}
		case questionnaireresponse.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_questionnaire_responses", values[i])
			} else if value.Valid {
				qr.user_questionnaire_responses = new(string)
				*qr.user_questionnaire_responses = value.String
			}
		default:
			qr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the QuestionnaireResponse.
// This includes values selected through modifiers, order, etc.
func (qr *QuestionnaireResponse) Value(name string) (ent.Value, error) {
	return qr.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the QuestionnaireResponse entity.
func (qr *QuestionnaireResponse) QueryUser() *UserQuery {
	return NewQuestionnaireResponseClient(qr.config).QueryUser(qr)
}

// QueryQuestionnaire queries the "questionnaire" edge of the QuestionnaireResponse entity.
func (qr *QuestionnaireResponse) QueryQuestionnaire() *QuestionnaireQuery {
	return NewQuestionnaireResponseClient(qr.config).QueryQuestionnaire(qr)
}

// QueryAnswers queries the "answers" edge of the QuestionnaireResponse entity.
func (qr *QuestionnaireResponse) QueryAnswers() *AnswerQuery {
	return NewQuestionnaireResponseClient(qr.config).QueryAnswers(qr)
}

// Update returns a builder for updating this QuestionnaireResponse.
// Note that you need to call QuestionnaireResponse.Unwrap() before calling this method if this QuestionnaireResponse
// was returned from a transaction, and the transaction was committed or rolled back.
func (qr *QuestionnaireResponse) Update() *QuestionnaireResponseUpdateOne {
	return NewQuestionnaireResponseClient(qr.config).UpdateOne(qr)
}

// Unwrap unwraps the QuestionnaireResponse entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (qr *QuestionnaireResponse) Unwrap() *QuestionnaireResponse {
	_tx, ok := qr.config.driver.(*txDriver)
	if !ok {
		panic("ent: QuestionnaireResponse is not a transactional entity")
	}
	qr.config.driver = _tx.drv
	return qr
}

// String implements the fmt.Stringer.
func (qr *QuestionnaireResponse) String() string {
	var builder strings.Builder
	builder.WriteString("QuestionnaireResponse(")
	builder.WriteString(fmt.Sprintf("id=%v, ", qr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(qr.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (qr *QuestionnaireResponse) MarshalJSON() ([]byte, error) {
	type Alias QuestionnaireResponse
	return json.Marshal(&struct {
		*Alias
		QuestionnaireResponseEdges
	}{
		Alias:                      (*Alias)(qr),
		QuestionnaireResponseEdges: qr.Edges,
	})
}

// QuestionnaireResponses is a parsable slice of QuestionnaireResponse.
type QuestionnaireResponses []*QuestionnaireResponse
