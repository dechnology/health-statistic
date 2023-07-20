// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/eesoymilk/health-statistic-api/ent/predicate"
	"github.com/eesoymilk/health-statistic-api/ent/questionnaireresponse"
	"github.com/eesoymilk/health-statistic-api/ent/user"
	"github.com/google/uuid"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetCreatedAt sets the "created_at" field.
func (uu *UserUpdate) SetCreatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreatedAt(*t)
	}
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetFirstName sets the "first_name" field.
func (uu *UserUpdate) SetFirstName(s string) *UserUpdate {
	uu.mutation.SetFirstName(s)
	return uu
}

// SetLastName sets the "last_name" field.
func (uu *UserUpdate) SetLastName(s string) *UserUpdate {
	uu.mutation.SetLastName(s)
	return uu
}

// SetBirthYear sets the "birth_year" field.
func (uu *UserUpdate) SetBirthYear(i int) *UserUpdate {
	uu.mutation.ResetBirthYear()
	uu.mutation.SetBirthYear(i)
	return uu
}

// AddBirthYear adds i to the "birth_year" field.
func (uu *UserUpdate) AddBirthYear(i int) *UserUpdate {
	uu.mutation.AddBirthYear(i)
	return uu
}

// SetHeight sets the "height" field.
func (uu *UserUpdate) SetHeight(f float64) *UserUpdate {
	uu.mutation.ResetHeight()
	uu.mutation.SetHeight(f)
	return uu
}

// AddHeight adds f to the "height" field.
func (uu *UserUpdate) AddHeight(f float64) *UserUpdate {
	uu.mutation.AddHeight(f)
	return uu
}

// SetWeight sets the "weight" field.
func (uu *UserUpdate) SetWeight(f float64) *UserUpdate {
	uu.mutation.ResetWeight()
	uu.mutation.SetWeight(f)
	return uu
}

// AddWeight adds f to the "weight" field.
func (uu *UserUpdate) AddWeight(f float64) *UserUpdate {
	uu.mutation.AddWeight(f)
	return uu
}

// SetGender sets the "gender" field.
func (uu *UserUpdate) SetGender(u user.Gender) *UserUpdate {
	uu.mutation.SetGender(u)
	return uu
}

// SetEducationLevel sets the "education_level" field.
func (uu *UserUpdate) SetEducationLevel(ul user.EducationLevel) *UserUpdate {
	uu.mutation.SetEducationLevel(ul)
	return uu
}

// SetOccupation sets the "occupation" field.
func (uu *UserUpdate) SetOccupation(u user.Occupation) *UserUpdate {
	uu.mutation.SetOccupation(u)
	return uu
}

// SetMarriage sets the "marriage" field.
func (uu *UserUpdate) SetMarriage(u user.Marriage) *UserUpdate {
	uu.mutation.SetMarriage(u)
	return uu
}

// SetMedicalHistory sets the "medical_history" field.
func (uu *UserUpdate) SetMedicalHistory(s string) *UserUpdate {
	uu.mutation.SetMedicalHistory(s)
	return uu
}

// SetNillableMedicalHistory sets the "medical_history" field if the given value is not nil.
func (uu *UserUpdate) SetNillableMedicalHistory(s *string) *UserUpdate {
	if s != nil {
		uu.SetMedicalHistory(*s)
	}
	return uu
}

// ClearMedicalHistory clears the value of the "medical_history" field.
func (uu *UserUpdate) ClearMedicalHistory() *UserUpdate {
	uu.mutation.ClearMedicalHistory()
	return uu
}

// SetMedicationStatus sets the "medication_status" field.
func (uu *UserUpdate) SetMedicationStatus(s string) *UserUpdate {
	uu.mutation.SetMedicationStatus(s)
	return uu
}

// SetNillableMedicationStatus sets the "medication_status" field if the given value is not nil.
func (uu *UserUpdate) SetNillableMedicationStatus(s *string) *UserUpdate {
	if s != nil {
		uu.SetMedicationStatus(*s)
	}
	return uu
}

// ClearMedicationStatus clears the value of the "medication_status" field.
func (uu *UserUpdate) ClearMedicationStatus() *UserUpdate {
	uu.mutation.ClearMedicationStatus()
	return uu
}

// SetDementedAmongDirectRelatives sets the "demented_among_direct_relatives" field.
func (uu *UserUpdate) SetDementedAmongDirectRelatives(b bool) *UserUpdate {
	uu.mutation.SetDementedAmongDirectRelatives(b)
	return uu
}

// SetHeadInjuryExperience sets the "head_injury_experience" field.
func (uu *UserUpdate) SetHeadInjuryExperience(b bool) *UserUpdate {
	uu.mutation.SetHeadInjuryExperience(b)
	return uu
}

// SetEarCondition sets the "ear_condition" field.
func (uu *UserUpdate) SetEarCondition(uc user.EarCondition) *UserUpdate {
	uu.mutation.SetEarCondition(uc)
	return uu
}

// SetEyesightCondition sets the "eyesight_condition" field.
func (uu *UserUpdate) SetEyesightCondition(uc user.EyesightCondition) *UserUpdate {
	uu.mutation.SetEyesightCondition(uc)
	return uu
}

// SetSmokingHabit sets the "smoking_habit" field.
func (uu *UserUpdate) SetSmokingHabit(uh user.SmokingHabit) *UserUpdate {
	uu.mutation.SetSmokingHabit(uh)
	return uu
}

// AddQuestionnaireResponseIDs adds the "questionnaire_responses" edge to the QuestionnaireResponse entity by IDs.
func (uu *UserUpdate) AddQuestionnaireResponseIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddQuestionnaireResponseIDs(ids...)
	return uu
}

// AddQuestionnaireResponses adds the "questionnaire_responses" edges to the QuestionnaireResponse entity.
func (uu *UserUpdate) AddQuestionnaireResponses(q ...*QuestionnaireResponse) *UserUpdate {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return uu.AddQuestionnaireResponseIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearQuestionnaireResponses clears all "questionnaire_responses" edges to the QuestionnaireResponse entity.
func (uu *UserUpdate) ClearQuestionnaireResponses() *UserUpdate {
	uu.mutation.ClearQuestionnaireResponses()
	return uu
}

// RemoveQuestionnaireResponseIDs removes the "questionnaire_responses" edge to QuestionnaireResponse entities by IDs.
func (uu *UserUpdate) RemoveQuestionnaireResponseIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveQuestionnaireResponseIDs(ids...)
	return uu
}

// RemoveQuestionnaireResponses removes "questionnaire_responses" edges to QuestionnaireResponse entities.
func (uu *UserUpdate) RemoveQuestionnaireResponses(q ...*QuestionnaireResponse) *UserUpdate {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return uu.RemoveQuestionnaireResponseIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.FirstName(); ok {
		if err := user.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "User.first_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.LastName(); ok {
		if err := user.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "User.last_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.BirthYear(); ok {
		if err := user.BirthYearValidator(v); err != nil {
			return &ValidationError{Name: "birth_year", err: fmt.Errorf(`ent: validator failed for field "User.birth_year": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Height(); ok {
		if err := user.HeightValidator(v); err != nil {
			return &ValidationError{Name: "height", err: fmt.Errorf(`ent: validator failed for field "User.height": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Weight(); ok {
		if err := user.WeightValidator(v); err != nil {
			return &ValidationError{Name: "weight", err: fmt.Errorf(`ent: validator failed for field "User.weight": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Gender(); ok {
		if err := user.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "User.gender": %w`, err)}
		}
	}
	if v, ok := uu.mutation.EducationLevel(); ok {
		if err := user.EducationLevelValidator(v); err != nil {
			return &ValidationError{Name: "education_level", err: fmt.Errorf(`ent: validator failed for field "User.education_level": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Occupation(); ok {
		if err := user.OccupationValidator(v); err != nil {
			return &ValidationError{Name: "occupation", err: fmt.Errorf(`ent: validator failed for field "User.occupation": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Marriage(); ok {
		if err := user.MarriageValidator(v); err != nil {
			return &ValidationError{Name: "marriage", err: fmt.Errorf(`ent: validator failed for field "User.marriage": %w`, err)}
		}
	}
	if v, ok := uu.mutation.MedicalHistory(); ok {
		if err := user.MedicalHistoryValidator(v); err != nil {
			return &ValidationError{Name: "medical_history", err: fmt.Errorf(`ent: validator failed for field "User.medical_history": %w`, err)}
		}
	}
	if v, ok := uu.mutation.MedicationStatus(); ok {
		if err := user.MedicationStatusValidator(v); err != nil {
			return &ValidationError{Name: "medication_status", err: fmt.Errorf(`ent: validator failed for field "User.medication_status": %w`, err)}
		}
	}
	if v, ok := uu.mutation.EarCondition(); ok {
		if err := user.EarConditionValidator(v); err != nil {
			return &ValidationError{Name: "ear_condition", err: fmt.Errorf(`ent: validator failed for field "User.ear_condition": %w`, err)}
		}
	}
	if v, ok := uu.mutation.EyesightCondition(); ok {
		if err := user.EyesightConditionValidator(v); err != nil {
			return &ValidationError{Name: "eyesight_condition", err: fmt.Errorf(`ent: validator failed for field "User.eyesight_condition": %w`, err)}
		}
	}
	if v, ok := uu.mutation.SmokingHabit(); ok {
		if err := user.SmokingHabitValidator(v); err != nil {
			return &ValidationError{Name: "smoking_habit", err: fmt.Errorf(`ent: validator failed for field "User.smoking_habit": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if value, ok := uu.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if value, ok := uu.mutation.BirthYear(); ok {
		_spec.SetField(user.FieldBirthYear, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedBirthYear(); ok {
		_spec.AddField(user.FieldBirthYear, field.TypeInt, value)
	}
	if value, ok := uu.mutation.Height(); ok {
		_spec.SetField(user.FieldHeight, field.TypeFloat64, value)
	}
	if value, ok := uu.mutation.AddedHeight(); ok {
		_spec.AddField(user.FieldHeight, field.TypeFloat64, value)
	}
	if value, ok := uu.mutation.Weight(); ok {
		_spec.SetField(user.FieldWeight, field.TypeFloat64, value)
	}
	if value, ok := uu.mutation.AddedWeight(); ok {
		_spec.AddField(user.FieldWeight, field.TypeFloat64, value)
	}
	if value, ok := uu.mutation.Gender(); ok {
		_spec.SetField(user.FieldGender, field.TypeEnum, value)
	}
	if value, ok := uu.mutation.EducationLevel(); ok {
		_spec.SetField(user.FieldEducationLevel, field.TypeEnum, value)
	}
	if value, ok := uu.mutation.Occupation(); ok {
		_spec.SetField(user.FieldOccupation, field.TypeEnum, value)
	}
	if value, ok := uu.mutation.Marriage(); ok {
		_spec.SetField(user.FieldMarriage, field.TypeEnum, value)
	}
	if value, ok := uu.mutation.MedicalHistory(); ok {
		_spec.SetField(user.FieldMedicalHistory, field.TypeString, value)
	}
	if uu.mutation.MedicalHistoryCleared() {
		_spec.ClearField(user.FieldMedicalHistory, field.TypeString)
	}
	if value, ok := uu.mutation.MedicationStatus(); ok {
		_spec.SetField(user.FieldMedicationStatus, field.TypeString, value)
	}
	if uu.mutation.MedicationStatusCleared() {
		_spec.ClearField(user.FieldMedicationStatus, field.TypeString)
	}
	if value, ok := uu.mutation.DementedAmongDirectRelatives(); ok {
		_spec.SetField(user.FieldDementedAmongDirectRelatives, field.TypeBool, value)
	}
	if value, ok := uu.mutation.HeadInjuryExperience(); ok {
		_spec.SetField(user.FieldHeadInjuryExperience, field.TypeBool, value)
	}
	if value, ok := uu.mutation.EarCondition(); ok {
		_spec.SetField(user.FieldEarCondition, field.TypeEnum, value)
	}
	if value, ok := uu.mutation.EyesightCondition(); ok {
		_spec.SetField(user.FieldEyesightCondition, field.TypeEnum, value)
	}
	if value, ok := uu.mutation.SmokingHabit(); ok {
		_spec.SetField(user.FieldSmokingHabit, field.TypeEnum, value)
	}
	if uu.mutation.QuestionnaireResponsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.QuestionnaireResponsesTable,
			Columns: []string{user.QuestionnaireResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionnaireresponse.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedQuestionnaireResponsesIDs(); len(nodes) > 0 && !uu.mutation.QuestionnaireResponsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.QuestionnaireResponsesTable,
			Columns: []string{user.QuestionnaireResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionnaireresponse.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.QuestionnaireResponsesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.QuestionnaireResponsesTable,
			Columns: []string{user.QuestionnaireResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionnaireresponse.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetCreatedAt sets the "created_at" field.
func (uuo *UserUpdateOne) SetCreatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreatedAt(*t)
	}
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetFirstName sets the "first_name" field.
func (uuo *UserUpdateOne) SetFirstName(s string) *UserUpdateOne {
	uuo.mutation.SetFirstName(s)
	return uuo
}

// SetLastName sets the "last_name" field.
func (uuo *UserUpdateOne) SetLastName(s string) *UserUpdateOne {
	uuo.mutation.SetLastName(s)
	return uuo
}

// SetBirthYear sets the "birth_year" field.
func (uuo *UserUpdateOne) SetBirthYear(i int) *UserUpdateOne {
	uuo.mutation.ResetBirthYear()
	uuo.mutation.SetBirthYear(i)
	return uuo
}

// AddBirthYear adds i to the "birth_year" field.
func (uuo *UserUpdateOne) AddBirthYear(i int) *UserUpdateOne {
	uuo.mutation.AddBirthYear(i)
	return uuo
}

// SetHeight sets the "height" field.
func (uuo *UserUpdateOne) SetHeight(f float64) *UserUpdateOne {
	uuo.mutation.ResetHeight()
	uuo.mutation.SetHeight(f)
	return uuo
}

// AddHeight adds f to the "height" field.
func (uuo *UserUpdateOne) AddHeight(f float64) *UserUpdateOne {
	uuo.mutation.AddHeight(f)
	return uuo
}

// SetWeight sets the "weight" field.
func (uuo *UserUpdateOne) SetWeight(f float64) *UserUpdateOne {
	uuo.mutation.ResetWeight()
	uuo.mutation.SetWeight(f)
	return uuo
}

// AddWeight adds f to the "weight" field.
func (uuo *UserUpdateOne) AddWeight(f float64) *UserUpdateOne {
	uuo.mutation.AddWeight(f)
	return uuo
}

// SetGender sets the "gender" field.
func (uuo *UserUpdateOne) SetGender(u user.Gender) *UserUpdateOne {
	uuo.mutation.SetGender(u)
	return uuo
}

// SetEducationLevel sets the "education_level" field.
func (uuo *UserUpdateOne) SetEducationLevel(ul user.EducationLevel) *UserUpdateOne {
	uuo.mutation.SetEducationLevel(ul)
	return uuo
}

// SetOccupation sets the "occupation" field.
func (uuo *UserUpdateOne) SetOccupation(u user.Occupation) *UserUpdateOne {
	uuo.mutation.SetOccupation(u)
	return uuo
}

// SetMarriage sets the "marriage" field.
func (uuo *UserUpdateOne) SetMarriage(u user.Marriage) *UserUpdateOne {
	uuo.mutation.SetMarriage(u)
	return uuo
}

// SetMedicalHistory sets the "medical_history" field.
func (uuo *UserUpdateOne) SetMedicalHistory(s string) *UserUpdateOne {
	uuo.mutation.SetMedicalHistory(s)
	return uuo
}

// SetNillableMedicalHistory sets the "medical_history" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableMedicalHistory(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetMedicalHistory(*s)
	}
	return uuo
}

// ClearMedicalHistory clears the value of the "medical_history" field.
func (uuo *UserUpdateOne) ClearMedicalHistory() *UserUpdateOne {
	uuo.mutation.ClearMedicalHistory()
	return uuo
}

// SetMedicationStatus sets the "medication_status" field.
func (uuo *UserUpdateOne) SetMedicationStatus(s string) *UserUpdateOne {
	uuo.mutation.SetMedicationStatus(s)
	return uuo
}

// SetNillableMedicationStatus sets the "medication_status" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableMedicationStatus(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetMedicationStatus(*s)
	}
	return uuo
}

// ClearMedicationStatus clears the value of the "medication_status" field.
func (uuo *UserUpdateOne) ClearMedicationStatus() *UserUpdateOne {
	uuo.mutation.ClearMedicationStatus()
	return uuo
}

// SetDementedAmongDirectRelatives sets the "demented_among_direct_relatives" field.
func (uuo *UserUpdateOne) SetDementedAmongDirectRelatives(b bool) *UserUpdateOne {
	uuo.mutation.SetDementedAmongDirectRelatives(b)
	return uuo
}

// SetHeadInjuryExperience sets the "head_injury_experience" field.
func (uuo *UserUpdateOne) SetHeadInjuryExperience(b bool) *UserUpdateOne {
	uuo.mutation.SetHeadInjuryExperience(b)
	return uuo
}

// SetEarCondition sets the "ear_condition" field.
func (uuo *UserUpdateOne) SetEarCondition(uc user.EarCondition) *UserUpdateOne {
	uuo.mutation.SetEarCondition(uc)
	return uuo
}

// SetEyesightCondition sets the "eyesight_condition" field.
func (uuo *UserUpdateOne) SetEyesightCondition(uc user.EyesightCondition) *UserUpdateOne {
	uuo.mutation.SetEyesightCondition(uc)
	return uuo
}

// SetSmokingHabit sets the "smoking_habit" field.
func (uuo *UserUpdateOne) SetSmokingHabit(uh user.SmokingHabit) *UserUpdateOne {
	uuo.mutation.SetSmokingHabit(uh)
	return uuo
}

// AddQuestionnaireResponseIDs adds the "questionnaire_responses" edge to the QuestionnaireResponse entity by IDs.
func (uuo *UserUpdateOne) AddQuestionnaireResponseIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddQuestionnaireResponseIDs(ids...)
	return uuo
}

// AddQuestionnaireResponses adds the "questionnaire_responses" edges to the QuestionnaireResponse entity.
func (uuo *UserUpdateOne) AddQuestionnaireResponses(q ...*QuestionnaireResponse) *UserUpdateOne {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return uuo.AddQuestionnaireResponseIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearQuestionnaireResponses clears all "questionnaire_responses" edges to the QuestionnaireResponse entity.
func (uuo *UserUpdateOne) ClearQuestionnaireResponses() *UserUpdateOne {
	uuo.mutation.ClearQuestionnaireResponses()
	return uuo
}

// RemoveQuestionnaireResponseIDs removes the "questionnaire_responses" edge to QuestionnaireResponse entities by IDs.
func (uuo *UserUpdateOne) RemoveQuestionnaireResponseIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveQuestionnaireResponseIDs(ids...)
	return uuo
}

// RemoveQuestionnaireResponses removes "questionnaire_responses" edges to QuestionnaireResponse entities.
func (uuo *UserUpdateOne) RemoveQuestionnaireResponses(q ...*QuestionnaireResponse) *UserUpdateOne {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return uuo.RemoveQuestionnaireResponseIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.FirstName(); ok {
		if err := user.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "User.first_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.LastName(); ok {
		if err := user.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "User.last_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.BirthYear(); ok {
		if err := user.BirthYearValidator(v); err != nil {
			return &ValidationError{Name: "birth_year", err: fmt.Errorf(`ent: validator failed for field "User.birth_year": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Height(); ok {
		if err := user.HeightValidator(v); err != nil {
			return &ValidationError{Name: "height", err: fmt.Errorf(`ent: validator failed for field "User.height": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Weight(); ok {
		if err := user.WeightValidator(v); err != nil {
			return &ValidationError{Name: "weight", err: fmt.Errorf(`ent: validator failed for field "User.weight": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Gender(); ok {
		if err := user.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "User.gender": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.EducationLevel(); ok {
		if err := user.EducationLevelValidator(v); err != nil {
			return &ValidationError{Name: "education_level", err: fmt.Errorf(`ent: validator failed for field "User.education_level": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Occupation(); ok {
		if err := user.OccupationValidator(v); err != nil {
			return &ValidationError{Name: "occupation", err: fmt.Errorf(`ent: validator failed for field "User.occupation": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Marriage(); ok {
		if err := user.MarriageValidator(v); err != nil {
			return &ValidationError{Name: "marriage", err: fmt.Errorf(`ent: validator failed for field "User.marriage": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.MedicalHistory(); ok {
		if err := user.MedicalHistoryValidator(v); err != nil {
			return &ValidationError{Name: "medical_history", err: fmt.Errorf(`ent: validator failed for field "User.medical_history": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.MedicationStatus(); ok {
		if err := user.MedicationStatusValidator(v); err != nil {
			return &ValidationError{Name: "medication_status", err: fmt.Errorf(`ent: validator failed for field "User.medication_status": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.EarCondition(); ok {
		if err := user.EarConditionValidator(v); err != nil {
			return &ValidationError{Name: "ear_condition", err: fmt.Errorf(`ent: validator failed for field "User.ear_condition": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.EyesightCondition(); ok {
		if err := user.EyesightConditionValidator(v); err != nil {
			return &ValidationError{Name: "eyesight_condition", err: fmt.Errorf(`ent: validator failed for field "User.eyesight_condition": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.SmokingHabit(); ok {
		if err := user.SmokingHabitValidator(v); err != nil {
			return &ValidationError{Name: "smoking_habit", err: fmt.Errorf(`ent: validator failed for field "User.smoking_habit": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.BirthYear(); ok {
		_spec.SetField(user.FieldBirthYear, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedBirthYear(); ok {
		_spec.AddField(user.FieldBirthYear, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.Height(); ok {
		_spec.SetField(user.FieldHeight, field.TypeFloat64, value)
	}
	if value, ok := uuo.mutation.AddedHeight(); ok {
		_spec.AddField(user.FieldHeight, field.TypeFloat64, value)
	}
	if value, ok := uuo.mutation.Weight(); ok {
		_spec.SetField(user.FieldWeight, field.TypeFloat64, value)
	}
	if value, ok := uuo.mutation.AddedWeight(); ok {
		_spec.AddField(user.FieldWeight, field.TypeFloat64, value)
	}
	if value, ok := uuo.mutation.Gender(); ok {
		_spec.SetField(user.FieldGender, field.TypeEnum, value)
	}
	if value, ok := uuo.mutation.EducationLevel(); ok {
		_spec.SetField(user.FieldEducationLevel, field.TypeEnum, value)
	}
	if value, ok := uuo.mutation.Occupation(); ok {
		_spec.SetField(user.FieldOccupation, field.TypeEnum, value)
	}
	if value, ok := uuo.mutation.Marriage(); ok {
		_spec.SetField(user.FieldMarriage, field.TypeEnum, value)
	}
	if value, ok := uuo.mutation.MedicalHistory(); ok {
		_spec.SetField(user.FieldMedicalHistory, field.TypeString, value)
	}
	if uuo.mutation.MedicalHistoryCleared() {
		_spec.ClearField(user.FieldMedicalHistory, field.TypeString)
	}
	if value, ok := uuo.mutation.MedicationStatus(); ok {
		_spec.SetField(user.FieldMedicationStatus, field.TypeString, value)
	}
	if uuo.mutation.MedicationStatusCleared() {
		_spec.ClearField(user.FieldMedicationStatus, field.TypeString)
	}
	if value, ok := uuo.mutation.DementedAmongDirectRelatives(); ok {
		_spec.SetField(user.FieldDementedAmongDirectRelatives, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.HeadInjuryExperience(); ok {
		_spec.SetField(user.FieldHeadInjuryExperience, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.EarCondition(); ok {
		_spec.SetField(user.FieldEarCondition, field.TypeEnum, value)
	}
	if value, ok := uuo.mutation.EyesightCondition(); ok {
		_spec.SetField(user.FieldEyesightCondition, field.TypeEnum, value)
	}
	if value, ok := uuo.mutation.SmokingHabit(); ok {
		_spec.SetField(user.FieldSmokingHabit, field.TypeEnum, value)
	}
	if uuo.mutation.QuestionnaireResponsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.QuestionnaireResponsesTable,
			Columns: []string{user.QuestionnaireResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionnaireresponse.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedQuestionnaireResponsesIDs(); len(nodes) > 0 && !uuo.mutation.QuestionnaireResponsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.QuestionnaireResponsesTable,
			Columns: []string{user.QuestionnaireResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionnaireresponse.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.QuestionnaireResponsesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.QuestionnaireResponsesTable,
			Columns: []string{user.QuestionnaireResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionnaireresponse.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}