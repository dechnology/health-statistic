package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Questionnaire holds the schema definition for the Questionnaire entity.
type Questionnaire struct {
	ent.Schema
}

// Fields of the Questionnaire.
func (Questionnaire) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Questionnaire.
func (Questionnaire) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("questions", Question.Type),
		edge.To("responses", UserQuestionnaire.Type),
	}
}
