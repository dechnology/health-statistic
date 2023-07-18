package schema

import (
	"time"

	"entgo.io/ent"
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
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Time("created_at").Default(time.Now),
		field.Text("body"),
	}
}

// Edges of the Answer.
func (Answer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("question", Question.Type).
			Ref("answers").
			Unique(),
		edge.From("questionnaire_response", QuestionnaireResponse.Type).
			Ref("answers").
			Unique(),
	}
}
