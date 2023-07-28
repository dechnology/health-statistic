package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Question holds the schema definition for the Question entity.
type Question struct {
	ent.Schema
}

// Fields of the Question.
func (Question) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Default(uuid.New).
			Unique(),
		field.Enum("type").
			Values(
				"short_answer",
				"single_choice",
				"multiple_choice",
			),
		field.Text("body").
			NotEmpty(),
		field.Int("order").
			NonNegative(),
	}
}

// Edges of the Question.
func (Question) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("questionnaire", Questionnaire.Type).
			Ref("questions").
			Required().
			Unique(),
		edge.To("choices", Choice.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("answers", Answer.Type),
	}
}
