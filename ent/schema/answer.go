package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Answer holds the schema definition for the Answer entity.
type Answer struct {
	ent.Schema
}

// Fields of the Answer.
func (Answer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Unique().
			Default(uuid.New),
		field.Time("created_at").
			Default(time.Now),
		field.Text("body").
			Optional().
			Nillable(),
	}
}

// Edges of the Answer.
func (Answer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("chosen", Choice.Type).
			Annotations(entsql.OnDelete(entsql.SetNull)),
		edge.From("question", Question.Type).
			Ref("answers").
			Required().
			Unique(),
		edge.From("questionnaire_response", QuestionnaireResponse.Type).
			Ref("answers").
			Required().
			Unique(),
	}
}
