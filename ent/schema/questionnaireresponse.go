package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// QuestionnaireResponse holds the schema definition for the QuestionnaireResponse entity.
type QuestionnaireResponse struct {
	ent.Schema
}

// Fields of the QuestionnaireResponse.
func (QuestionnaireResponse) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the QuestionnaireResponse.
func (QuestionnaireResponse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("questionnaire_responses"),
		edge.From("questionnaire", Questionnaire.Type).
			Ref("questionnaire_responses"),
		edge.To("answers", Answer.Type),
	}
}
