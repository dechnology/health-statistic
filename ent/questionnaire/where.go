// Code generated by ent, DO NOT EDIT.

package questionnaire

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/eesoymilk/health-statistic-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldEQ(FieldName, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldEQ(FieldCreatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldContainsFold(FieldName, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Questionnaire {
	return predicate.Questionnaire(sql.FieldLTE(FieldCreatedAt, v))
}

// HasQuestions applies the HasEdge predicate on the "questions" edge.
func HasQuestions() predicate.Questionnaire {
	return predicate.Questionnaire(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, QuestionsTable, QuestionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasQuestionsWith applies the HasEdge predicate on the "questions" edge with a given conditions (other predicates).
func HasQuestionsWith(preds ...predicate.Question) predicate.Questionnaire {
	return predicate.Questionnaire(func(s *sql.Selector) {
		step := newQuestionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasQuestionnaireResponses applies the HasEdge predicate on the "questionnaire_responses" edge.
func HasQuestionnaireResponses() predicate.Questionnaire {
	return predicate.Questionnaire(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, QuestionnaireResponsesTable, QuestionnaireResponsesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasQuestionnaireResponsesWith applies the HasEdge predicate on the "questionnaire_responses" edge with a given conditions (other predicates).
func HasQuestionnaireResponsesWith(preds ...predicate.QuestionnaireResponse) predicate.Questionnaire {
	return predicate.Questionnaire(func(s *sql.Selector) {
		step := newQuestionnaireResponsesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Questionnaire) predicate.Questionnaire {
	return predicate.Questionnaire(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Questionnaire) predicate.Questionnaire {
	return predicate.Questionnaire(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Questionnaire) predicate.Questionnaire {
	return predicate.Questionnaire(func(s *sql.Selector) {
		p(s.Not())
	})
}
