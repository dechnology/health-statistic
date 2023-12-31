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
	"github.com/eesoymilk/health-statistic-api/ent/answer"
	"github.com/eesoymilk/health-statistic-api/ent/choice"
	"github.com/eesoymilk/health-statistic-api/ent/predicate"
	"github.com/eesoymilk/health-statistic-api/ent/question"
	"github.com/eesoymilk/health-statistic-api/ent/questionnaireresponse"
	"github.com/google/uuid"
)

// AnswerUpdate is the builder for updating Answer entities.
type AnswerUpdate struct {
	config
	hooks    []Hook
	mutation *AnswerMutation
}

// Where appends a list predicates to the AnswerUpdate builder.
func (au *AnswerUpdate) Where(ps ...predicate.Answer) *AnswerUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AnswerUpdate) SetCreatedAt(t time.Time) *AnswerUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AnswerUpdate) SetNillableCreatedAt(t *time.Time) *AnswerUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// SetBody sets the "body" field.
func (au *AnswerUpdate) SetBody(s string) *AnswerUpdate {
	au.mutation.SetBody(s)
	return au
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (au *AnswerUpdate) SetNillableBody(s *string) *AnswerUpdate {
	if s != nil {
		au.SetBody(*s)
	}
	return au
}

// ClearBody clears the value of the "body" field.
func (au *AnswerUpdate) ClearBody() *AnswerUpdate {
	au.mutation.ClearBody()
	return au
}

// AddChosenIDs adds the "chosen" edge to the Choice entity by IDs.
func (au *AnswerUpdate) AddChosenIDs(ids ...uuid.UUID) *AnswerUpdate {
	au.mutation.AddChosenIDs(ids...)
	return au
}

// AddChosen adds the "chosen" edges to the Choice entity.
func (au *AnswerUpdate) AddChosen(c ...*Choice) *AnswerUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return au.AddChosenIDs(ids...)
}

// SetQuestionID sets the "question" edge to the Question entity by ID.
func (au *AnswerUpdate) SetQuestionID(id uuid.UUID) *AnswerUpdate {
	au.mutation.SetQuestionID(id)
	return au
}

// SetQuestion sets the "question" edge to the Question entity.
func (au *AnswerUpdate) SetQuestion(q *Question) *AnswerUpdate {
	return au.SetQuestionID(q.ID)
}

// SetQuestionnaireResponseID sets the "questionnaire_response" edge to the QuestionnaireResponse entity by ID.
func (au *AnswerUpdate) SetQuestionnaireResponseID(id uuid.UUID) *AnswerUpdate {
	au.mutation.SetQuestionnaireResponseID(id)
	return au
}

// SetQuestionnaireResponse sets the "questionnaire_response" edge to the QuestionnaireResponse entity.
func (au *AnswerUpdate) SetQuestionnaireResponse(q *QuestionnaireResponse) *AnswerUpdate {
	return au.SetQuestionnaireResponseID(q.ID)
}

// Mutation returns the AnswerMutation object of the builder.
func (au *AnswerUpdate) Mutation() *AnswerMutation {
	return au.mutation
}

// ClearChosen clears all "chosen" edges to the Choice entity.
func (au *AnswerUpdate) ClearChosen() *AnswerUpdate {
	au.mutation.ClearChosen()
	return au
}

// RemoveChosenIDs removes the "chosen" edge to Choice entities by IDs.
func (au *AnswerUpdate) RemoveChosenIDs(ids ...uuid.UUID) *AnswerUpdate {
	au.mutation.RemoveChosenIDs(ids...)
	return au
}

// RemoveChosen removes "chosen" edges to Choice entities.
func (au *AnswerUpdate) RemoveChosen(c ...*Choice) *AnswerUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return au.RemoveChosenIDs(ids...)
}

// ClearQuestion clears the "question" edge to the Question entity.
func (au *AnswerUpdate) ClearQuestion() *AnswerUpdate {
	au.mutation.ClearQuestion()
	return au
}

// ClearQuestionnaireResponse clears the "questionnaire_response" edge to the QuestionnaireResponse entity.
func (au *AnswerUpdate) ClearQuestionnaireResponse() *AnswerUpdate {
	au.mutation.ClearQuestionnaireResponse()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AnswerUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AnswerUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AnswerUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AnswerUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AnswerUpdate) check() error {
	if _, ok := au.mutation.QuestionID(); au.mutation.QuestionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Answer.question"`)
	}
	if _, ok := au.mutation.QuestionnaireResponseID(); au.mutation.QuestionnaireResponseCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Answer.questionnaire_response"`)
	}
	return nil
}

func (au *AnswerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(answer.Table, answer.Columns, sqlgraph.NewFieldSpec(answer.FieldID, field.TypeUUID))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(answer.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.Body(); ok {
		_spec.SetField(answer.FieldBody, field.TypeString, value)
	}
	if au.mutation.BodyCleared() {
		_spec.ClearField(answer.FieldBody, field.TypeString)
	}
	if au.mutation.ChosenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   answer.ChosenTable,
			Columns: answer.ChosenPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(choice.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedChosenIDs(); len(nodes) > 0 && !au.mutation.ChosenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   answer.ChosenTable,
			Columns: answer.ChosenPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(choice.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.ChosenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   answer.ChosenTable,
			Columns: answer.ChosenPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(choice.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.QuestionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   answer.QuestionTable,
			Columns: []string{answer.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.QuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   answer.QuestionTable,
			Columns: []string{answer.QuestionColumn},
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
	if au.mutation.QuestionnaireResponseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   answer.QuestionnaireResponseTable,
			Columns: []string{answer.QuestionnaireResponseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionnaireresponse.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.QuestionnaireResponseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   answer.QuestionnaireResponseTable,
			Columns: []string{answer.QuestionnaireResponseColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{answer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AnswerUpdateOne is the builder for updating a single Answer entity.
type AnswerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AnswerMutation
}

// SetCreatedAt sets the "created_at" field.
func (auo *AnswerUpdateOne) SetCreatedAt(t time.Time) *AnswerUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AnswerUpdateOne) SetNillableCreatedAt(t *time.Time) *AnswerUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// SetBody sets the "body" field.
func (auo *AnswerUpdateOne) SetBody(s string) *AnswerUpdateOne {
	auo.mutation.SetBody(s)
	return auo
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (auo *AnswerUpdateOne) SetNillableBody(s *string) *AnswerUpdateOne {
	if s != nil {
		auo.SetBody(*s)
	}
	return auo
}

// ClearBody clears the value of the "body" field.
func (auo *AnswerUpdateOne) ClearBody() *AnswerUpdateOne {
	auo.mutation.ClearBody()
	return auo
}

// AddChosenIDs adds the "chosen" edge to the Choice entity by IDs.
func (auo *AnswerUpdateOne) AddChosenIDs(ids ...uuid.UUID) *AnswerUpdateOne {
	auo.mutation.AddChosenIDs(ids...)
	return auo
}

// AddChosen adds the "chosen" edges to the Choice entity.
func (auo *AnswerUpdateOne) AddChosen(c ...*Choice) *AnswerUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return auo.AddChosenIDs(ids...)
}

// SetQuestionID sets the "question" edge to the Question entity by ID.
func (auo *AnswerUpdateOne) SetQuestionID(id uuid.UUID) *AnswerUpdateOne {
	auo.mutation.SetQuestionID(id)
	return auo
}

// SetQuestion sets the "question" edge to the Question entity.
func (auo *AnswerUpdateOne) SetQuestion(q *Question) *AnswerUpdateOne {
	return auo.SetQuestionID(q.ID)
}

// SetQuestionnaireResponseID sets the "questionnaire_response" edge to the QuestionnaireResponse entity by ID.
func (auo *AnswerUpdateOne) SetQuestionnaireResponseID(id uuid.UUID) *AnswerUpdateOne {
	auo.mutation.SetQuestionnaireResponseID(id)
	return auo
}

// SetQuestionnaireResponse sets the "questionnaire_response" edge to the QuestionnaireResponse entity.
func (auo *AnswerUpdateOne) SetQuestionnaireResponse(q *QuestionnaireResponse) *AnswerUpdateOne {
	return auo.SetQuestionnaireResponseID(q.ID)
}

// Mutation returns the AnswerMutation object of the builder.
func (auo *AnswerUpdateOne) Mutation() *AnswerMutation {
	return auo.mutation
}

// ClearChosen clears all "chosen" edges to the Choice entity.
func (auo *AnswerUpdateOne) ClearChosen() *AnswerUpdateOne {
	auo.mutation.ClearChosen()
	return auo
}

// RemoveChosenIDs removes the "chosen" edge to Choice entities by IDs.
func (auo *AnswerUpdateOne) RemoveChosenIDs(ids ...uuid.UUID) *AnswerUpdateOne {
	auo.mutation.RemoveChosenIDs(ids...)
	return auo
}

// RemoveChosen removes "chosen" edges to Choice entities.
func (auo *AnswerUpdateOne) RemoveChosen(c ...*Choice) *AnswerUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return auo.RemoveChosenIDs(ids...)
}

// ClearQuestion clears the "question" edge to the Question entity.
func (auo *AnswerUpdateOne) ClearQuestion() *AnswerUpdateOne {
	auo.mutation.ClearQuestion()
	return auo
}

// ClearQuestionnaireResponse clears the "questionnaire_response" edge to the QuestionnaireResponse entity.
func (auo *AnswerUpdateOne) ClearQuestionnaireResponse() *AnswerUpdateOne {
	auo.mutation.ClearQuestionnaireResponse()
	return auo
}

// Where appends a list predicates to the AnswerUpdate builder.
func (auo *AnswerUpdateOne) Where(ps ...predicate.Answer) *AnswerUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AnswerUpdateOne) Select(field string, fields ...string) *AnswerUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Answer entity.
func (auo *AnswerUpdateOne) Save(ctx context.Context) (*Answer, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AnswerUpdateOne) SaveX(ctx context.Context) *Answer {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AnswerUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AnswerUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AnswerUpdateOne) check() error {
	if _, ok := auo.mutation.QuestionID(); auo.mutation.QuestionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Answer.question"`)
	}
	if _, ok := auo.mutation.QuestionnaireResponseID(); auo.mutation.QuestionnaireResponseCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Answer.questionnaire_response"`)
	}
	return nil
}

func (auo *AnswerUpdateOne) sqlSave(ctx context.Context) (_node *Answer, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(answer.Table, answer.Columns, sqlgraph.NewFieldSpec(answer.FieldID, field.TypeUUID))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Answer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, answer.FieldID)
		for _, f := range fields {
			if !answer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != answer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(answer.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.Body(); ok {
		_spec.SetField(answer.FieldBody, field.TypeString, value)
	}
	if auo.mutation.BodyCleared() {
		_spec.ClearField(answer.FieldBody, field.TypeString)
	}
	if auo.mutation.ChosenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   answer.ChosenTable,
			Columns: answer.ChosenPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(choice.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedChosenIDs(); len(nodes) > 0 && !auo.mutation.ChosenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   answer.ChosenTable,
			Columns: answer.ChosenPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(choice.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.ChosenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   answer.ChosenTable,
			Columns: answer.ChosenPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(choice.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.QuestionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   answer.QuestionTable,
			Columns: []string{answer.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.QuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   answer.QuestionTable,
			Columns: []string{answer.QuestionColumn},
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
	if auo.mutation.QuestionnaireResponseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   answer.QuestionnaireResponseTable,
			Columns: []string{answer.QuestionnaireResponseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionnaireresponse.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.QuestionnaireResponseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   answer.QuestionnaireResponseTable,
			Columns: []string{answer.QuestionnaireResponseColumn},
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
	_node = &Answer{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{answer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
