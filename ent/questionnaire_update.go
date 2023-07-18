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
	"github.com/eesoymilk/health-statistic-api/ent/question"
	"github.com/eesoymilk/health-statistic-api/ent/questionnaire"
	"github.com/eesoymilk/health-statistic-api/ent/questionnaireresponse"
	"github.com/google/uuid"
)

// QuestionnaireUpdate is the builder for updating Questionnaire entities.
type QuestionnaireUpdate struct {
	config
	hooks    []Hook
	mutation *QuestionnaireMutation
}

// Where appends a list predicates to the QuestionnaireUpdate builder.
func (qu *QuestionnaireUpdate) Where(ps ...predicate.Questionnaire) *QuestionnaireUpdate {
	qu.mutation.Where(ps...)
	return qu
}

// SetName sets the "name" field.
func (qu *QuestionnaireUpdate) SetName(s string) *QuestionnaireUpdate {
	qu.mutation.SetName(s)
	return qu
}

// SetCreatedAt sets the "created_at" field.
func (qu *QuestionnaireUpdate) SetCreatedAt(t time.Time) *QuestionnaireUpdate {
	qu.mutation.SetCreatedAt(t)
	return qu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (qu *QuestionnaireUpdate) SetNillableCreatedAt(t *time.Time) *QuestionnaireUpdate {
	if t != nil {
		qu.SetCreatedAt(*t)
	}
	return qu
}

// AddQuestionIDs adds the "questions" edge to the Question entity by IDs.
func (qu *QuestionnaireUpdate) AddQuestionIDs(ids ...uuid.UUID) *QuestionnaireUpdate {
	qu.mutation.AddQuestionIDs(ids...)
	return qu
}

// AddQuestions adds the "questions" edges to the Question entity.
func (qu *QuestionnaireUpdate) AddQuestions(q ...*Question) *QuestionnaireUpdate {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return qu.AddQuestionIDs(ids...)
}

// AddQuestionnaireResponseIDs adds the "questionnaire_responses" edge to the QuestionnaireResponse entity by IDs.
func (qu *QuestionnaireUpdate) AddQuestionnaireResponseIDs(ids ...uuid.UUID) *QuestionnaireUpdate {
	qu.mutation.AddQuestionnaireResponseIDs(ids...)
	return qu
}

// AddQuestionnaireResponses adds the "questionnaire_responses" edges to the QuestionnaireResponse entity.
func (qu *QuestionnaireUpdate) AddQuestionnaireResponses(q ...*QuestionnaireResponse) *QuestionnaireUpdate {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return qu.AddQuestionnaireResponseIDs(ids...)
}

// Mutation returns the QuestionnaireMutation object of the builder.
func (qu *QuestionnaireUpdate) Mutation() *QuestionnaireMutation {
	return qu.mutation
}

// ClearQuestions clears all "questions" edges to the Question entity.
func (qu *QuestionnaireUpdate) ClearQuestions() *QuestionnaireUpdate {
	qu.mutation.ClearQuestions()
	return qu
}

// RemoveQuestionIDs removes the "questions" edge to Question entities by IDs.
func (qu *QuestionnaireUpdate) RemoveQuestionIDs(ids ...uuid.UUID) *QuestionnaireUpdate {
	qu.mutation.RemoveQuestionIDs(ids...)
	return qu
}

// RemoveQuestions removes "questions" edges to Question entities.
func (qu *QuestionnaireUpdate) RemoveQuestions(q ...*Question) *QuestionnaireUpdate {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return qu.RemoveQuestionIDs(ids...)
}

// ClearQuestionnaireResponses clears all "questionnaire_responses" edges to the QuestionnaireResponse entity.
func (qu *QuestionnaireUpdate) ClearQuestionnaireResponses() *QuestionnaireUpdate {
	qu.mutation.ClearQuestionnaireResponses()
	return qu
}

// RemoveQuestionnaireResponseIDs removes the "questionnaire_responses" edge to QuestionnaireResponse entities by IDs.
func (qu *QuestionnaireUpdate) RemoveQuestionnaireResponseIDs(ids ...uuid.UUID) *QuestionnaireUpdate {
	qu.mutation.RemoveQuestionnaireResponseIDs(ids...)
	return qu
}

// RemoveQuestionnaireResponses removes "questionnaire_responses" edges to QuestionnaireResponse entities.
func (qu *QuestionnaireUpdate) RemoveQuestionnaireResponses(q ...*QuestionnaireResponse) *QuestionnaireUpdate {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return qu.RemoveQuestionnaireResponseIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (qu *QuestionnaireUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, qu.sqlSave, qu.mutation, qu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (qu *QuestionnaireUpdate) SaveX(ctx context.Context) int {
	affected, err := qu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (qu *QuestionnaireUpdate) Exec(ctx context.Context) error {
	_, err := qu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qu *QuestionnaireUpdate) ExecX(ctx context.Context) {
	if err := qu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (qu *QuestionnaireUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(questionnaire.Table, questionnaire.Columns, sqlgraph.NewFieldSpec(questionnaire.FieldID, field.TypeUUID))
	if ps := qu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := qu.mutation.Name(); ok {
		_spec.SetField(questionnaire.FieldName, field.TypeString, value)
	}
	if value, ok := qu.mutation.CreatedAt(); ok {
		_spec.SetField(questionnaire.FieldCreatedAt, field.TypeTime, value)
	}
	if qu.mutation.QuestionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   questionnaire.QuestionsTable,
			Columns: []string{questionnaire.QuestionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.RemovedQuestionsIDs(); len(nodes) > 0 && !qu.mutation.QuestionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   questionnaire.QuestionsTable,
			Columns: []string{questionnaire.QuestionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.QuestionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   questionnaire.QuestionsTable,
			Columns: []string{questionnaire.QuestionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qu.mutation.QuestionnaireResponsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   questionnaire.QuestionnaireResponsesTable,
			Columns: []string{questionnaire.QuestionnaireResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionnaireresponse.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.RemovedQuestionnaireResponsesIDs(); len(nodes) > 0 && !qu.mutation.QuestionnaireResponsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   questionnaire.QuestionnaireResponsesTable,
			Columns: []string{questionnaire.QuestionnaireResponsesColumn},
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
	if nodes := qu.mutation.QuestionnaireResponsesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   questionnaire.QuestionnaireResponsesTable,
			Columns: []string{questionnaire.QuestionnaireResponsesColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, qu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{questionnaire.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	qu.mutation.done = true
	return n, nil
}

// QuestionnaireUpdateOne is the builder for updating a single Questionnaire entity.
type QuestionnaireUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *QuestionnaireMutation
}

// SetName sets the "name" field.
func (quo *QuestionnaireUpdateOne) SetName(s string) *QuestionnaireUpdateOne {
	quo.mutation.SetName(s)
	return quo
}

// SetCreatedAt sets the "created_at" field.
func (quo *QuestionnaireUpdateOne) SetCreatedAt(t time.Time) *QuestionnaireUpdateOne {
	quo.mutation.SetCreatedAt(t)
	return quo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (quo *QuestionnaireUpdateOne) SetNillableCreatedAt(t *time.Time) *QuestionnaireUpdateOne {
	if t != nil {
		quo.SetCreatedAt(*t)
	}
	return quo
}

// AddQuestionIDs adds the "questions" edge to the Question entity by IDs.
func (quo *QuestionnaireUpdateOne) AddQuestionIDs(ids ...uuid.UUID) *QuestionnaireUpdateOne {
	quo.mutation.AddQuestionIDs(ids...)
	return quo
}

// AddQuestions adds the "questions" edges to the Question entity.
func (quo *QuestionnaireUpdateOne) AddQuestions(q ...*Question) *QuestionnaireUpdateOne {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return quo.AddQuestionIDs(ids...)
}

// AddQuestionnaireResponseIDs adds the "questionnaire_responses" edge to the QuestionnaireResponse entity by IDs.
func (quo *QuestionnaireUpdateOne) AddQuestionnaireResponseIDs(ids ...uuid.UUID) *QuestionnaireUpdateOne {
	quo.mutation.AddQuestionnaireResponseIDs(ids...)
	return quo
}

// AddQuestionnaireResponses adds the "questionnaire_responses" edges to the QuestionnaireResponse entity.
func (quo *QuestionnaireUpdateOne) AddQuestionnaireResponses(q ...*QuestionnaireResponse) *QuestionnaireUpdateOne {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return quo.AddQuestionnaireResponseIDs(ids...)
}

// Mutation returns the QuestionnaireMutation object of the builder.
func (quo *QuestionnaireUpdateOne) Mutation() *QuestionnaireMutation {
	return quo.mutation
}

// ClearQuestions clears all "questions" edges to the Question entity.
func (quo *QuestionnaireUpdateOne) ClearQuestions() *QuestionnaireUpdateOne {
	quo.mutation.ClearQuestions()
	return quo
}

// RemoveQuestionIDs removes the "questions" edge to Question entities by IDs.
func (quo *QuestionnaireUpdateOne) RemoveQuestionIDs(ids ...uuid.UUID) *QuestionnaireUpdateOne {
	quo.mutation.RemoveQuestionIDs(ids...)
	return quo
}

// RemoveQuestions removes "questions" edges to Question entities.
func (quo *QuestionnaireUpdateOne) RemoveQuestions(q ...*Question) *QuestionnaireUpdateOne {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return quo.RemoveQuestionIDs(ids...)
}

// ClearQuestionnaireResponses clears all "questionnaire_responses" edges to the QuestionnaireResponse entity.
func (quo *QuestionnaireUpdateOne) ClearQuestionnaireResponses() *QuestionnaireUpdateOne {
	quo.mutation.ClearQuestionnaireResponses()
	return quo
}

// RemoveQuestionnaireResponseIDs removes the "questionnaire_responses" edge to QuestionnaireResponse entities by IDs.
func (quo *QuestionnaireUpdateOne) RemoveQuestionnaireResponseIDs(ids ...uuid.UUID) *QuestionnaireUpdateOne {
	quo.mutation.RemoveQuestionnaireResponseIDs(ids...)
	return quo
}

// RemoveQuestionnaireResponses removes "questionnaire_responses" edges to QuestionnaireResponse entities.
func (quo *QuestionnaireUpdateOne) RemoveQuestionnaireResponses(q ...*QuestionnaireResponse) *QuestionnaireUpdateOne {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return quo.RemoveQuestionnaireResponseIDs(ids...)
}

// Where appends a list predicates to the QuestionnaireUpdate builder.
func (quo *QuestionnaireUpdateOne) Where(ps ...predicate.Questionnaire) *QuestionnaireUpdateOne {
	quo.mutation.Where(ps...)
	return quo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (quo *QuestionnaireUpdateOne) Select(field string, fields ...string) *QuestionnaireUpdateOne {
	quo.fields = append([]string{field}, fields...)
	return quo
}

// Save executes the query and returns the updated Questionnaire entity.
func (quo *QuestionnaireUpdateOne) Save(ctx context.Context) (*Questionnaire, error) {
	return withHooks(ctx, quo.sqlSave, quo.mutation, quo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (quo *QuestionnaireUpdateOne) SaveX(ctx context.Context) *Questionnaire {
	node, err := quo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (quo *QuestionnaireUpdateOne) Exec(ctx context.Context) error {
	_, err := quo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (quo *QuestionnaireUpdateOne) ExecX(ctx context.Context) {
	if err := quo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (quo *QuestionnaireUpdateOne) sqlSave(ctx context.Context) (_node *Questionnaire, err error) {
	_spec := sqlgraph.NewUpdateSpec(questionnaire.Table, questionnaire.Columns, sqlgraph.NewFieldSpec(questionnaire.FieldID, field.TypeUUID))
	id, ok := quo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Questionnaire.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := quo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, questionnaire.FieldID)
		for _, f := range fields {
			if !questionnaire.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != questionnaire.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := quo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := quo.mutation.Name(); ok {
		_spec.SetField(questionnaire.FieldName, field.TypeString, value)
	}
	if value, ok := quo.mutation.CreatedAt(); ok {
		_spec.SetField(questionnaire.FieldCreatedAt, field.TypeTime, value)
	}
	if quo.mutation.QuestionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   questionnaire.QuestionsTable,
			Columns: []string{questionnaire.QuestionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.RemovedQuestionsIDs(); len(nodes) > 0 && !quo.mutation.QuestionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   questionnaire.QuestionsTable,
			Columns: []string{questionnaire.QuestionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.QuestionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   questionnaire.QuestionsTable,
			Columns: []string{questionnaire.QuestionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if quo.mutation.QuestionnaireResponsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   questionnaire.QuestionnaireResponsesTable,
			Columns: []string{questionnaire.QuestionnaireResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionnaireresponse.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.RemovedQuestionnaireResponsesIDs(); len(nodes) > 0 && !quo.mutation.QuestionnaireResponsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   questionnaire.QuestionnaireResponsesTable,
			Columns: []string{questionnaire.QuestionnaireResponsesColumn},
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
	if nodes := quo.mutation.QuestionnaireResponsesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   questionnaire.QuestionnaireResponsesTable,
			Columns: []string{questionnaire.QuestionnaireResponsesColumn},
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
	_node = &Questionnaire{config: quo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, quo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{questionnaire.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	quo.mutation.done = true
	return _node, nil
}
