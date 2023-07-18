package schema

import "entgo.io/ent"

// Reward holds the schema definition for the Reward entity.
type Reward struct {
	ent.Schema
}

// Fields of the Reward.
func (Reward) Fields() []ent.Field {
	return nil
}

// Edges of the Reward.
func (Reward) Edges() []ent.Edge {
	return nil
}
