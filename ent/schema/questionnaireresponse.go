package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// QuestionnaireResponse holds the schema definition for the QuestionnaireResponse entity.
type QuestionnaireResponse struct {
	ent.Schema
}

// Fields of the QuestionnaireResponse.
func (QuestionnaireResponse) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the QuestionnaireResponse.
func (QuestionnaireResponse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Ref("questionnaire_responses").
			Unique(),
		edge.From("questionnaire", Questionnaire.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Ref("questionnaire_responses").
			Unique(),
		edge.To("answers", Answer.Type),
	}
}
