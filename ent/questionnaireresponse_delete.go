// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/eesoymilk/health-statistic-api/ent/predicate"
	"github.com/eesoymilk/health-statistic-api/ent/questionnaireresponse"
)

// QuestionnaireResponseDelete is the builder for deleting a QuestionnaireResponse entity.
type QuestionnaireResponseDelete struct {
	config
	hooks    []Hook
	mutation *QuestionnaireResponseMutation
}

// Where appends a list predicates to the QuestionnaireResponseDelete builder.
func (qrd *QuestionnaireResponseDelete) Where(ps ...predicate.QuestionnaireResponse) *QuestionnaireResponseDelete {
	qrd.mutation.Where(ps...)
	return qrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (qrd *QuestionnaireResponseDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, qrd.sqlExec, qrd.mutation, qrd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (qrd *QuestionnaireResponseDelete) ExecX(ctx context.Context) int {
	n, err := qrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (qrd *QuestionnaireResponseDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(questionnaireresponse.Table, sqlgraph.NewFieldSpec(questionnaireresponse.FieldID, field.TypeUUID))
	if ps := qrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, qrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	qrd.mutation.done = true
	return affected, err
}

// QuestionnaireResponseDeleteOne is the builder for deleting a single QuestionnaireResponse entity.
type QuestionnaireResponseDeleteOne struct {
	qrd *QuestionnaireResponseDelete
}

// Where appends a list predicates to the QuestionnaireResponseDelete builder.
func (qrdo *QuestionnaireResponseDeleteOne) Where(ps ...predicate.QuestionnaireResponse) *QuestionnaireResponseDeleteOne {
	qrdo.qrd.mutation.Where(ps...)
	return qrdo
}

// Exec executes the deletion query.
func (qrdo *QuestionnaireResponseDeleteOne) Exec(ctx context.Context) error {
	n, err := qrdo.qrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{questionnaireresponse.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (qrdo *QuestionnaireResponseDeleteOne) ExecX(ctx context.Context) {
	if err := qrdo.Exec(ctx); err != nil {
		panic(err)
	}
}
