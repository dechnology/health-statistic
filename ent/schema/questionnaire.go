package schema

import "entgo.io/ent"

// Questionnaire holds the schema definition for the Questionnaire entity.
type Questionnaire struct {
	ent.Schema
}

// Fields of the Questionnaire.
func (Questionnaire) Fields() []ent.Field {
	return nil
}

// Edges of the Questionnaire.
func (Questionnaire) Edges() []ent.Edge {
	return nil
}
