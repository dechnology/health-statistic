// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/eesoymilk/health-statistic-api/ent/questionnaireresponse"
	"github.com/eesoymilk/health-statistic-api/ent/user"
	"github.com/google/uuid"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetBirthYear sets the "birth_year" field.
func (uc *UserCreate) SetBirthYear(i int) *UserCreate {
	uc.mutation.SetBirthYear(i)
	return uc
}

// SetHeight sets the "height" field.
func (uc *UserCreate) SetHeight(f float64) *UserCreate {
	uc.mutation.SetHeight(f)
	return uc
}

// SetWeight sets the "weight" field.
func (uc *UserCreate) SetWeight(f float64) *UserCreate {
	uc.mutation.SetWeight(f)
	return uc
}

// SetGender sets the "gender" field.
func (uc *UserCreate) SetGender(u user.Gender) *UserCreate {
	uc.mutation.SetGender(u)
	return uc
}

// SetEducationLevel sets the "education_level" field.
func (uc *UserCreate) SetEducationLevel(ul user.EducationLevel) *UserCreate {
	uc.mutation.SetEducationLevel(ul)
	return uc
}

// SetOccupation sets the "occupation" field.
func (uc *UserCreate) SetOccupation(u user.Occupation) *UserCreate {
	uc.mutation.SetOccupation(u)
	return uc
}

// SetMarriage sets the "marriage" field.
func (uc *UserCreate) SetMarriage(u user.Marriage) *UserCreate {
	uc.mutation.SetMarriage(u)
	return uc
}

// SetMedicalHistory sets the "medical_history" field.
func (uc *UserCreate) SetMedicalHistory(s string) *UserCreate {
	uc.mutation.SetMedicalHistory(s)
	return uc
}

// SetNillableMedicalHistory sets the "medical_history" field if the given value is not nil.
func (uc *UserCreate) SetNillableMedicalHistory(s *string) *UserCreate {
	if s != nil {
		uc.SetMedicalHistory(*s)
	}
	return uc
}

// SetMedicationStatus sets the "medication_status" field.
func (uc *UserCreate) SetMedicationStatus(s string) *UserCreate {
	uc.mutation.SetMedicationStatus(s)
	return uc
}

// SetNillableMedicationStatus sets the "medication_status" field if the given value is not nil.
func (uc *UserCreate) SetNillableMedicationStatus(s *string) *UserCreate {
	if s != nil {
		uc.SetMedicationStatus(*s)
	}
	return uc
}

// SetDementedAmongDirectRelatives sets the "demented_among_direct_relatives" field.
func (uc *UserCreate) SetDementedAmongDirectRelatives(b bool) *UserCreate {
	uc.mutation.SetDementedAmongDirectRelatives(b)
	return uc
}

// SetHeadInjuryExperience sets the "head_injury_experience" field.
func (uc *UserCreate) SetHeadInjuryExperience(b bool) *UserCreate {
	uc.mutation.SetHeadInjuryExperience(b)
	return uc
}

// SetEarCondition sets the "ear_condition" field.
func (uc *UserCreate) SetEarCondition(value user.EarCondition) *UserCreate {
	uc.mutation.SetEarCondition(value)
	return uc
}

// SetEyesightCondition sets the "eyesight_condition" field.
func (uc *UserCreate) SetEyesightCondition(value user.EyesightCondition) *UserCreate {
	uc.mutation.SetEyesightCondition(value)
	return uc
}

// SetSmokingHabit sets the "smoking_habit" field.
func (uc *UserCreate) SetSmokingHabit(uh user.SmokingHabit) *UserCreate {
	uc.mutation.SetSmokingHabit(uh)
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(s string) *UserCreate {
	uc.mutation.SetID(s)
	return uc
}

// AddQuestionnaireResponseIDs adds the "questionnaire_responses" edge to the QuestionnaireResponse entity by IDs.
func (uc *UserCreate) AddQuestionnaireResponseIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddQuestionnaireResponseIDs(ids...)
	return uc
}

// AddQuestionnaireResponses adds the "questionnaire_responses" edges to the QuestionnaireResponse entity.
func (uc *UserCreate) AddQuestionnaireResponses(q ...*QuestionnaireResponse) *UserCreate {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return uc.AddQuestionnaireResponseIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		v := user.DefaultUpdatedAt()
		uc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "User.created_at"`)}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "User.updated_at"`)}
	}
	if _, ok := uc.mutation.BirthYear(); !ok {
		return &ValidationError{Name: "birth_year", err: errors.New(`ent: missing required field "User.birth_year"`)}
	}
	if v, ok := uc.mutation.BirthYear(); ok {
		if err := user.BirthYearValidator(v); err != nil {
			return &ValidationError{Name: "birth_year", err: fmt.Errorf(`ent: validator failed for field "User.birth_year": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Height(); !ok {
		return &ValidationError{Name: "height", err: errors.New(`ent: missing required field "User.height"`)}
	}
	if v, ok := uc.mutation.Height(); ok {
		if err := user.HeightValidator(v); err != nil {
			return &ValidationError{Name: "height", err: fmt.Errorf(`ent: validator failed for field "User.height": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Weight(); !ok {
		return &ValidationError{Name: "weight", err: errors.New(`ent: missing required field "User.weight"`)}
	}
	if v, ok := uc.mutation.Weight(); ok {
		if err := user.WeightValidator(v); err != nil {
			return &ValidationError{Name: "weight", err: fmt.Errorf(`ent: validator failed for field "User.weight": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Gender(); !ok {
		return &ValidationError{Name: "gender", err: errors.New(`ent: missing required field "User.gender"`)}
	}
	if v, ok := uc.mutation.Gender(); ok {
		if err := user.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "User.gender": %w`, err)}
		}
	}
	if _, ok := uc.mutation.EducationLevel(); !ok {
		return &ValidationError{Name: "education_level", err: errors.New(`ent: missing required field "User.education_level"`)}
	}
	if v, ok := uc.mutation.EducationLevel(); ok {
		if err := user.EducationLevelValidator(v); err != nil {
			return &ValidationError{Name: "education_level", err: fmt.Errorf(`ent: validator failed for field "User.education_level": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Occupation(); !ok {
		return &ValidationError{Name: "occupation", err: errors.New(`ent: missing required field "User.occupation"`)}
	}
	if v, ok := uc.mutation.Occupation(); ok {
		if err := user.OccupationValidator(v); err != nil {
			return &ValidationError{Name: "occupation", err: fmt.Errorf(`ent: validator failed for field "User.occupation": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Marriage(); !ok {
		return &ValidationError{Name: "marriage", err: errors.New(`ent: missing required field "User.marriage"`)}
	}
	if v, ok := uc.mutation.Marriage(); ok {
		if err := user.MarriageValidator(v); err != nil {
			return &ValidationError{Name: "marriage", err: fmt.Errorf(`ent: validator failed for field "User.marriage": %w`, err)}
		}
	}
	if v, ok := uc.mutation.MedicalHistory(); ok {
		if err := user.MedicalHistoryValidator(v); err != nil {
			return &ValidationError{Name: "medical_history", err: fmt.Errorf(`ent: validator failed for field "User.medical_history": %w`, err)}
		}
	}
	if v, ok := uc.mutation.MedicationStatus(); ok {
		if err := user.MedicationStatusValidator(v); err != nil {
			return &ValidationError{Name: "medication_status", err: fmt.Errorf(`ent: validator failed for field "User.medication_status": %w`, err)}
		}
	}
	if _, ok := uc.mutation.DementedAmongDirectRelatives(); !ok {
		return &ValidationError{Name: "demented_among_direct_relatives", err: errors.New(`ent: missing required field "User.demented_among_direct_relatives"`)}
	}
	if _, ok := uc.mutation.HeadInjuryExperience(); !ok {
		return &ValidationError{Name: "head_injury_experience", err: errors.New(`ent: missing required field "User.head_injury_experience"`)}
	}
	if _, ok := uc.mutation.EarCondition(); !ok {
		return &ValidationError{Name: "ear_condition", err: errors.New(`ent: missing required field "User.ear_condition"`)}
	}
	if v, ok := uc.mutation.EarCondition(); ok {
		if err := user.EarConditionValidator(v); err != nil {
			return &ValidationError{Name: "ear_condition", err: fmt.Errorf(`ent: validator failed for field "User.ear_condition": %w`, err)}
		}
	}
	if _, ok := uc.mutation.EyesightCondition(); !ok {
		return &ValidationError{Name: "eyesight_condition", err: errors.New(`ent: missing required field "User.eyesight_condition"`)}
	}
	if v, ok := uc.mutation.EyesightCondition(); ok {
		if err := user.EyesightConditionValidator(v); err != nil {
			return &ValidationError{Name: "eyesight_condition", err: fmt.Errorf(`ent: validator failed for field "User.eyesight_condition": %w`, err)}
		}
	}
	if _, ok := uc.mutation.SmokingHabit(); !ok {
		return &ValidationError{Name: "smoking_habit", err: errors.New(`ent: missing required field "User.smoking_habit"`)}
	}
	if v, ok := uc.mutation.SmokingHabit(); ok {
		if err := user.SmokingHabitValidator(v); err != nil {
			return &ValidationError{Name: "smoking_habit", err: fmt.Errorf(`ent: validator failed for field "User.smoking_habit": %w`, err)}
		}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected User.ID type: %T", _spec.ID.Value)
		}
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := uc.mutation.BirthYear(); ok {
		_spec.SetField(user.FieldBirthYear, field.TypeInt, value)
		_node.BirthYear = value
	}
	if value, ok := uc.mutation.Height(); ok {
		_spec.SetField(user.FieldHeight, field.TypeFloat64, value)
		_node.Height = value
	}
	if value, ok := uc.mutation.Weight(); ok {
		_spec.SetField(user.FieldWeight, field.TypeFloat64, value)
		_node.Weight = value
	}
	if value, ok := uc.mutation.Gender(); ok {
		_spec.SetField(user.FieldGender, field.TypeEnum, value)
		_node.Gender = value
	}
	if value, ok := uc.mutation.EducationLevel(); ok {
		_spec.SetField(user.FieldEducationLevel, field.TypeEnum, value)
		_node.EducationLevel = value
	}
	if value, ok := uc.mutation.Occupation(); ok {
		_spec.SetField(user.FieldOccupation, field.TypeEnum, value)
		_node.Occupation = value
	}
	if value, ok := uc.mutation.Marriage(); ok {
		_spec.SetField(user.FieldMarriage, field.TypeEnum, value)
		_node.Marriage = value
	}
	if value, ok := uc.mutation.MedicalHistory(); ok {
		_spec.SetField(user.FieldMedicalHistory, field.TypeString, value)
		_node.MedicalHistory = value
	}
	if value, ok := uc.mutation.MedicationStatus(); ok {
		_spec.SetField(user.FieldMedicationStatus, field.TypeString, value)
		_node.MedicationStatus = value
	}
	if value, ok := uc.mutation.DementedAmongDirectRelatives(); ok {
		_spec.SetField(user.FieldDementedAmongDirectRelatives, field.TypeBool, value)
		_node.DementedAmongDirectRelatives = value
	}
	if value, ok := uc.mutation.HeadInjuryExperience(); ok {
		_spec.SetField(user.FieldHeadInjuryExperience, field.TypeBool, value)
		_node.HeadInjuryExperience = value
	}
	if value, ok := uc.mutation.EarCondition(); ok {
		_spec.SetField(user.FieldEarCondition, field.TypeEnum, value)
		_node.EarCondition = value
	}
	if value, ok := uc.mutation.EyesightCondition(); ok {
		_spec.SetField(user.FieldEyesightCondition, field.TypeEnum, value)
		_node.EyesightCondition = value
	}
	if value, ok := uc.mutation.SmokingHabit(); ok {
		_spec.SetField(user.FieldSmokingHabit, field.TypeEnum, value)
		_node.SmokingHabit = value
	}
	if nodes := uc.mutation.QuestionnaireResponsesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
