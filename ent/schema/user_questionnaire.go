package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserQuestionnaire holds the schema definition for the UserQuestionnaire entity.
type UserQuestionnaire struct {
	ent.Schema
}

// Fields of the UserQuestionnaire.
func (UserQuestionnaire) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the UserQuestionnaire.
func (UserQuestionnaire) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("questionnaires").
			Unique(),
		edge.From("questionnaire", Questionnaire.Type).
			Ref("responses").
			Unique(),
		edge.To("answers", Answer.Type),
	}
}
