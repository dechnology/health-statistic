// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/eesoymilk/health-statistic-api/ent/answer"
	"github.com/eesoymilk/health-statistic-api/ent/choice"
	"github.com/eesoymilk/health-statistic-api/ent/predicate"
	"github.com/eesoymilk/health-statistic-api/ent/question"
	"github.com/eesoymilk/health-statistic-api/ent/questionnaire"
	"github.com/google/uuid"
)

// QuestionUpdate is the builder for updating Question entities.
type QuestionUpdate struct {
	config
	hooks    []Hook
	mutation *QuestionMutation
}

// Where appends a list predicates to the QuestionUpdate builder.
func (qu *QuestionUpdate) Where(ps ...predicate.Question) *QuestionUpdate {
	qu.mutation.Where(ps...)
	return qu
}

// SetType sets the "type" field.
func (qu *QuestionUpdate) SetType(q question.Type) *QuestionUpdate {
	qu.mutation.SetType(q)
	return qu
}

// SetBody sets the "body" field.
func (qu *QuestionUpdate) SetBody(s string) *QuestionUpdate {
	qu.mutation.SetBody(s)
	return qu
}

// SetOrder sets the "order" field.
func (qu *QuestionUpdate) SetOrder(i int) *QuestionUpdate {
	qu.mutation.ResetOrder()
	qu.mutation.SetOrder(i)
	return qu
}

// AddOrder adds i to the "order" field.
func (qu *QuestionUpdate) AddOrder(i int) *QuestionUpdate {
	qu.mutation.AddOrder(i)
	return qu
}

// SetQuestionnaireID sets the "questionnaire" edge to the Questionnaire entity by ID.
func (qu *QuestionUpdate) SetQuestionnaireID(id uuid.UUID) *QuestionUpdate {
	qu.mutation.SetQuestionnaireID(id)
	return qu
}

// SetQuestionnaire sets the "questionnaire" edge to the Questionnaire entity.
func (qu *QuestionUpdate) SetQuestionnaire(q *Questionnaire) *QuestionUpdate {
	return qu.SetQuestionnaireID(q.ID)
}

// AddChoiceIDs adds the "choices" edge to the Choice entity by IDs.
func (qu *QuestionUpdate) AddChoiceIDs(ids ...uuid.UUID) *QuestionUpdate {
	qu.mutation.AddChoiceIDs(ids...)
	return qu
}

// AddChoices adds the "choices" edges to the Choice entity.
func (qu *QuestionUpdate) AddChoices(c ...*Choice) *QuestionUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return qu.AddChoiceIDs(ids...)
}

// AddAnswerIDs adds the "answers" edge to the Answer entity by IDs.
func (qu *QuestionUpdate) AddAnswerIDs(ids ...uuid.UUID) *QuestionUpdate {
	qu.mutation.AddAnswerIDs(ids...)
	return qu
}

// AddAnswers adds the "answers" edges to the Answer entity.
func (qu *QuestionUpdate) AddAnswers(a ...*Answer) *QuestionUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return qu.AddAnswerIDs(ids...)
}

// Mutation returns the QuestionMutation object of the builder.
func (qu *QuestionUpdate) Mutation() *QuestionMutation {
	return qu.mutation
}

// ClearQuestionnaire clears the "questionnaire" edge to the Questionnaire entity.
func (qu *QuestionUpdate) ClearQuestionnaire() *QuestionUpdate {
	qu.mutation.ClearQuestionnaire()
	return qu
}

// ClearChoices clears all "choices" edges to the Choice entity.
func (qu *QuestionUpdate) ClearChoices() *QuestionUpdate {
	qu.mutation.ClearChoices()
	return qu
}

// RemoveChoiceIDs removes the "choices" edge to Choice entities by IDs.
func (qu *QuestionUpdate) RemoveChoiceIDs(ids ...uuid.UUID) *QuestionUpdate {
	qu.mutation.RemoveChoiceIDs(ids...)
	return qu
}

// RemoveChoices removes "choices" edges to Choice entities.
func (qu *QuestionUpdate) RemoveChoices(c ...*Choice) *QuestionUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return qu.RemoveChoiceIDs(ids...)
}

// ClearAnswers clears all "answers" edges to the Answer entity.
func (qu *QuestionUpdate) ClearAnswers() *QuestionUpdate {
	qu.mutation.ClearAnswers()
	return qu
}

// RemoveAnswerIDs removes the "answers" edge to Answer entities by IDs.
func (qu *QuestionUpdate) RemoveAnswerIDs(ids ...uuid.UUID) *QuestionUpdate {
	qu.mutation.RemoveAnswerIDs(ids...)
	return qu
}

// RemoveAnswers removes "answers" edges to Answer entities.
func (qu *QuestionUpdate) RemoveAnswers(a ...*Answer) *QuestionUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return qu.RemoveAnswerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (qu *QuestionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, qu.sqlSave, qu.mutation, qu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (qu *QuestionUpdate) SaveX(ctx context.Context) int {
	affected, err := qu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (qu *QuestionUpdate) Exec(ctx context.Context) error {
	_, err := qu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qu *QuestionUpdate) ExecX(ctx context.Context) {
	if err := qu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (qu *QuestionUpdate) check() error {
	if v, ok := qu.mutation.GetType(); ok {
		if err := question.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Question.type": %w`, err)}
		}
	}
	if v, ok := qu.mutation.Body(); ok {
		if err := question.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf(`ent: validator failed for field "Question.body": %w`, err)}
		}
	}
	if v, ok := qu.mutation.Order(); ok {
		if err := question.OrderValidator(v); err != nil {
			return &ValidationError{Name: "order", err: fmt.Errorf(`ent: validator failed for field "Question.order": %w`, err)}
		}
	}
	if _, ok := qu.mutation.QuestionnaireID(); qu.mutation.QuestionnaireCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Question.questionnaire"`)
	}
	return nil
}

func (qu *QuestionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := qu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(question.Table, question.Columns, sqlgraph.NewFieldSpec(question.FieldID, field.TypeUUID))
	if ps := qu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := qu.mutation.GetType(); ok {
		_spec.SetField(question.FieldType, field.TypeEnum, value)
	}
	if value, ok := qu.mutation.Body(); ok {
		_spec.SetField(question.FieldBody, field.TypeString, value)
	}
	if value, ok := qu.mutation.Order(); ok {
		_spec.SetField(question.FieldOrder, field.TypeInt, value)
	}
	if value, ok := qu.mutation.AddedOrder(); ok {
		_spec.AddField(question.FieldOrder, field.TypeInt, value)
	}
	if qu.mutation.QuestionnaireCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   question.QuestionnaireTable,
			Columns: []string{question.QuestionnaireColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionnaire.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.QuestionnaireIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   question.QuestionnaireTable,
			Columns: []string{question.QuestionnaireColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionnaire.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qu.mutation.ChoicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.ChoicesTable,
			Columns: []string{question.ChoicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(choice.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.RemovedChoicesIDs(); len(nodes) > 0 && !qu.mutation.ChoicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.ChoicesTable,
			Columns: []string{question.ChoicesColumn},
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
	if nodes := qu.mutation.ChoicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.ChoicesTable,
			Columns: []string{question.ChoicesColumn},
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
	if qu.mutation.AnswersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.AnswersTable,
			Columns: []string{question.AnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.RemovedAnswersIDs(); len(nodes) > 0 && !qu.mutation.AnswersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.AnswersTable,
			Columns: []string{question.AnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.AnswersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.AnswersTable,
			Columns: []string{question.AnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, qu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{question.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	qu.mutation.done = true
	return n, nil
}

// QuestionUpdateOne is the builder for updating a single Question entity.
type QuestionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *QuestionMutation
}

// SetType sets the "type" field.
func (quo *QuestionUpdateOne) SetType(q question.Type) *QuestionUpdateOne {
	quo.mutation.SetType(q)
	return quo
}

// SetBody sets the "body" field.
func (quo *QuestionUpdateOne) SetBody(s string) *QuestionUpdateOne {
	quo.mutation.SetBody(s)
	return quo
}

// SetOrder sets the "order" field.
func (quo *QuestionUpdateOne) SetOrder(i int) *QuestionUpdateOne {
	quo.mutation.ResetOrder()
	quo.mutation.SetOrder(i)
	return quo
}

// AddOrder adds i to the "order" field.
func (quo *QuestionUpdateOne) AddOrder(i int) *QuestionUpdateOne {
	quo.mutation.AddOrder(i)
	return quo
}

// SetQuestionnaireID sets the "questionnaire" edge to the Questionnaire entity by ID.
func (quo *QuestionUpdateOne) SetQuestionnaireID(id uuid.UUID) *QuestionUpdateOne {
	quo.mutation.SetQuestionnaireID(id)
	return quo
}

// SetQuestionnaire sets the "questionnaire" edge to the Questionnaire entity.
func (quo *QuestionUpdateOne) SetQuestionnaire(q *Questionnaire) *QuestionUpdateOne {
	return quo.SetQuestionnaireID(q.ID)
}

// AddChoiceIDs adds the "choices" edge to the Choice entity by IDs.
func (quo *QuestionUpdateOne) AddChoiceIDs(ids ...uuid.UUID) *QuestionUpdateOne {
	quo.mutation.AddChoiceIDs(ids...)
	return quo
}

// AddChoices adds the "choices" edges to the Choice entity.
func (quo *QuestionUpdateOne) AddChoices(c ...*Choice) *QuestionUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return quo.AddChoiceIDs(ids...)
}

// AddAnswerIDs adds the "answers" edge to the Answer entity by IDs.
func (quo *QuestionUpdateOne) AddAnswerIDs(ids ...uuid.UUID) *QuestionUpdateOne {
	quo.mutation.AddAnswerIDs(ids...)
	return quo
}

// AddAnswers adds the "answers" edges to the Answer entity.
func (quo *QuestionUpdateOne) AddAnswers(a ...*Answer) *QuestionUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return quo.AddAnswerIDs(ids...)
}

// Mutation returns the QuestionMutation object of the builder.
func (quo *QuestionUpdateOne) Mutation() *QuestionMutation {
	return quo.mutation
}

// ClearQuestionnaire clears the "questionnaire" edge to the Questionnaire entity.
func (quo *QuestionUpdateOne) ClearQuestionnaire() *QuestionUpdateOne {
	quo.mutation.ClearQuestionnaire()
	return quo
}

// ClearChoices clears all "choices" edges to the Choice entity.
func (quo *QuestionUpdateOne) ClearChoices() *QuestionUpdateOne {
	quo.mutation.ClearChoices()
	return quo
}

// RemoveChoiceIDs removes the "choices" edge to Choice entities by IDs.
func (quo *QuestionUpdateOne) RemoveChoiceIDs(ids ...uuid.UUID) *QuestionUpdateOne {
	quo.mutation.RemoveChoiceIDs(ids...)
	return quo
}

// RemoveChoices removes "choices" edges to Choice entities.
func (quo *QuestionUpdateOne) RemoveChoices(c ...*Choice) *QuestionUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return quo.RemoveChoiceIDs(ids...)
}

// ClearAnswers clears all "answers" edges to the Answer entity.
func (quo *QuestionUpdateOne) ClearAnswers() *QuestionUpdateOne {
	quo.mutation.ClearAnswers()
	return quo
}

// RemoveAnswerIDs removes the "answers" edge to Answer entities by IDs.
func (quo *QuestionUpdateOne) RemoveAnswerIDs(ids ...uuid.UUID) *QuestionUpdateOne {
	quo.mutation.RemoveAnswerIDs(ids...)
	return quo
}

// RemoveAnswers removes "answers" edges to Answer entities.
func (quo *QuestionUpdateOne) RemoveAnswers(a ...*Answer) *QuestionUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return quo.RemoveAnswerIDs(ids...)
}

// Where appends a list predicates to the QuestionUpdate builder.
func (quo *QuestionUpdateOne) Where(ps ...predicate.Question) *QuestionUpdateOne {
	quo.mutation.Where(ps...)
	return quo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (quo *QuestionUpdateOne) Select(field string, fields ...string) *QuestionUpdateOne {
	quo.fields = append([]string{field}, fields...)
	return quo
}

// Save executes the query and returns the updated Question entity.
func (quo *QuestionUpdateOne) Save(ctx context.Context) (*Question, error) {
	return withHooks(ctx, quo.sqlSave, quo.mutation, quo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (quo *QuestionUpdateOne) SaveX(ctx context.Context) *Question {
	node, err := quo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (quo *QuestionUpdateOne) Exec(ctx context.Context) error {
	_, err := quo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (quo *QuestionUpdateOne) ExecX(ctx context.Context) {
	if err := quo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (quo *QuestionUpdateOne) check() error {
	if v, ok := quo.mutation.GetType(); ok {
		if err := question.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Question.type": %w`, err)}
		}
	}
	if v, ok := quo.mutation.Body(); ok {
		if err := question.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf(`ent: validator failed for field "Question.body": %w`, err)}
		}
	}
	if v, ok := quo.mutation.Order(); ok {
		if err := question.OrderValidator(v); err != nil {
			return &ValidationError{Name: "order", err: fmt.Errorf(`ent: validator failed for field "Question.order": %w`, err)}
		}
	}
	if _, ok := quo.mutation.QuestionnaireID(); quo.mutation.QuestionnaireCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Question.questionnaire"`)
	}
	return nil
}

func (quo *QuestionUpdateOne) sqlSave(ctx context.Context) (_node *Question, err error) {
	if err := quo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(question.Table, question.Columns, sqlgraph.NewFieldSpec(question.FieldID, field.TypeUUID))
	id, ok := quo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Question.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := quo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, question.FieldID)
		for _, f := range fields {
			if !question.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != question.FieldID {
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
	if value, ok := quo.mutation.GetType(); ok {
		_spec.SetField(question.FieldType, field.TypeEnum, value)
	}
	if value, ok := quo.mutation.Body(); ok {
		_spec.SetField(question.FieldBody, field.TypeString, value)
	}
	if value, ok := quo.mutation.Order(); ok {
		_spec.SetField(question.FieldOrder, field.TypeInt, value)
	}
	if value, ok := quo.mutation.AddedOrder(); ok {
		_spec.AddField(question.FieldOrder, field.TypeInt, value)
	}
	if quo.mutation.QuestionnaireCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   question.QuestionnaireTable,
			Columns: []string{question.QuestionnaireColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionnaire.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.QuestionnaireIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   question.QuestionnaireTable,
			Columns: []string{question.QuestionnaireColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionnaire.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if quo.mutation.ChoicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.ChoicesTable,
			Columns: []string{question.ChoicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(choice.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.RemovedChoicesIDs(); len(nodes) > 0 && !quo.mutation.ChoicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.ChoicesTable,
			Columns: []string{question.ChoicesColumn},
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
	if nodes := quo.mutation.ChoicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.ChoicesTable,
			Columns: []string{question.ChoicesColumn},
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
	if quo.mutation.AnswersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.AnswersTable,
			Columns: []string{question.AnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.RemovedAnswersIDs(); len(nodes) > 0 && !quo.mutation.AnswersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.AnswersTable,
			Columns: []string{question.AnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.AnswersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.AnswersTable,
			Columns: []string{question.AnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Question{config: quo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, quo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{question.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	quo.mutation.done = true
	return _node, nil
}
