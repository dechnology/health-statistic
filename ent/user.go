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
	"github.com/eesoymilk/health-statistic-api/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// BirthYear holds the value of the "birth_year" field.
	BirthYear int `json:"birth_year,omitempty"`
	// Height holds the value of the "height" field.
	Height float64 `json:"height,omitempty"`
	// Weight holds the value of the "weight" field.
	Weight float64 `json:"weight,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender user.Gender `json:"gender,omitempty"`
	// EducationLevel holds the value of the "education_level" field.
	EducationLevel user.EducationLevel `json:"education_level,omitempty"`
	// Occupation holds the value of the "occupation" field.
	Occupation user.Occupation `json:"occupation,omitempty"`
	// Marriage holds the value of the "marriage" field.
	Marriage user.Marriage `json:"marriage,omitempty"`
	// MedicalHistory holds the value of the "medical_history" field.
	MedicalHistory string `json:"medical_history,omitempty"`
	// MedicationStatus holds the value of the "medication_status" field.
	MedicationStatus string `json:"medication_status,omitempty"`
	// DementedAmongDirectRelatives holds the value of the "demented_among_direct_relatives" field.
	DementedAmongDirectRelatives bool `json:"demented_among_direct_relatives,omitempty"`
	// HeadInjuryExperience holds the value of the "head_injury_experience" field.
	HeadInjuryExperience bool `json:"head_injury_experience,omitempty"`
	// EarCondition holds the value of the "ear_condition" field.
	EarCondition user.EarCondition `json:"ear_condition,omitempty"`
	// EyesightCondition holds the value of the "eyesight_condition" field.
	EyesightCondition user.EyesightCondition `json:"eyesight_condition,omitempty"`
	// SmokingHabit holds the value of the "smoking_habit" field.
	SmokingHabit user.SmokingHabit `json:"smoking_habit,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"-"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// QuestionnaireResponses holds the value of the questionnaire_responses edge.
	QuestionnaireResponses []*QuestionnaireResponse `json:"questionnaire_responses,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// QuestionnaireResponsesOrErr returns the QuestionnaireResponses value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) QuestionnaireResponsesOrErr() ([]*QuestionnaireResponse, error) {
	if e.loadedTypes[0] {
		return e.QuestionnaireResponses, nil
	}
	return nil, &NotLoadedError{edge: "questionnaire_responses"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldDementedAmongDirectRelatives, user.FieldHeadInjuryExperience:
			values[i] = new(sql.NullBool)
		case user.FieldHeight, user.FieldWeight:
			values[i] = new(sql.NullFloat64)
		case user.FieldBirthYear:
			values[i] = new(sql.NullInt64)
		case user.FieldID, user.FieldFirstName, user.FieldLastName, user.FieldGender, user.FieldEducationLevel, user.FieldOccupation, user.FieldMarriage, user.FieldMedicalHistory, user.FieldMedicationStatus, user.FieldEarCondition, user.FieldEyesightCondition, user.FieldSmokingHabit:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				u.ID = value.String
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				u.FirstName = value.String
			}
		case user.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				u.LastName = value.String
			}
		case user.FieldBirthYear:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field birth_year", values[i])
			} else if value.Valid {
				u.BirthYear = int(value.Int64)
			}
		case user.FieldHeight:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field height", values[i])
			} else if value.Valid {
				u.Height = value.Float64
			}
		case user.FieldWeight:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field weight", values[i])
			} else if value.Valid {
				u.Weight = value.Float64
			}
		case user.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				u.Gender = user.Gender(value.String)
			}
		case user.FieldEducationLevel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field education_level", values[i])
			} else if value.Valid {
				u.EducationLevel = user.EducationLevel(value.String)
			}
		case user.FieldOccupation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field occupation", values[i])
			} else if value.Valid {
				u.Occupation = user.Occupation(value.String)
			}
		case user.FieldMarriage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field marriage", values[i])
			} else if value.Valid {
				u.Marriage = user.Marriage(value.String)
			}
		case user.FieldMedicalHistory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field medical_history", values[i])
			} else if value.Valid {
				u.MedicalHistory = value.String
			}
		case user.FieldMedicationStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field medication_status", values[i])
			} else if value.Valid {
				u.MedicationStatus = value.String
			}
		case user.FieldDementedAmongDirectRelatives:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field demented_among_direct_relatives", values[i])
			} else if value.Valid {
				u.DementedAmongDirectRelatives = value.Bool
			}
		case user.FieldHeadInjuryExperience:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field head_injury_experience", values[i])
			} else if value.Valid {
				u.HeadInjuryExperience = value.Bool
			}
		case user.FieldEarCondition:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ear_condition", values[i])
			} else if value.Valid {
				u.EarCondition = user.EarCondition(value.String)
			}
		case user.FieldEyesightCondition:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field eyesight_condition", values[i])
			} else if value.Valid {
				u.EyesightCondition = user.EyesightCondition(value.String)
			}
		case user.FieldSmokingHabit:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field smoking_habit", values[i])
			} else if value.Valid {
				u.SmokingHabit = user.SmokingHabit(value.String)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryQuestionnaireResponses queries the "questionnaire_responses" edge of the User entity.
func (u *User) QueryQuestionnaireResponses() *QuestionnaireResponseQuery {
	return NewUserClient(u.config).QueryQuestionnaireResponses(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(u.FirstName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(u.LastName)
	builder.WriteString(", ")
	builder.WriteString("birth_year=")
	builder.WriteString(fmt.Sprintf("%v", u.BirthYear))
	builder.WriteString(", ")
	builder.WriteString("height=")
	builder.WriteString(fmt.Sprintf("%v", u.Height))
	builder.WriteString(", ")
	builder.WriteString("weight=")
	builder.WriteString(fmt.Sprintf("%v", u.Weight))
	builder.WriteString(", ")
	builder.WriteString("gender=")
	builder.WriteString(fmt.Sprintf("%v", u.Gender))
	builder.WriteString(", ")
	builder.WriteString("education_level=")
	builder.WriteString(fmt.Sprintf("%v", u.EducationLevel))
	builder.WriteString(", ")
	builder.WriteString("occupation=")
	builder.WriteString(fmt.Sprintf("%v", u.Occupation))
	builder.WriteString(", ")
	builder.WriteString("marriage=")
	builder.WriteString(fmt.Sprintf("%v", u.Marriage))
	builder.WriteString(", ")
	builder.WriteString("medical_history=")
	builder.WriteString(u.MedicalHistory)
	builder.WriteString(", ")
	builder.WriteString("medication_status=")
	builder.WriteString(u.MedicationStatus)
	builder.WriteString(", ")
	builder.WriteString("demented_among_direct_relatives=")
	builder.WriteString(fmt.Sprintf("%v", u.DementedAmongDirectRelatives))
	builder.WriteString(", ")
	builder.WriteString("head_injury_experience=")
	builder.WriteString(fmt.Sprintf("%v", u.HeadInjuryExperience))
	builder.WriteString(", ")
	builder.WriteString("ear_condition=")
	builder.WriteString(fmt.Sprintf("%v", u.EarCondition))
	builder.WriteString(", ")
	builder.WriteString("eyesight_condition=")
	builder.WriteString(fmt.Sprintf("%v", u.EyesightCondition))
	builder.WriteString(", ")
	builder.WriteString("smoking_habit=")
	builder.WriteString(fmt.Sprintf("%v", u.SmokingHabit))
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (u *User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		*Alias
		UserEdges
	}{
		Alias:     (*Alias)(u),
		UserEdges: u.Edges,
	})
}

// Users is a parsable slice of User.
type Users []*User
