package schema

import "entgo.io/ent"

// Price holds the schema definition for the Price entity.
type Price struct {
	ent.Schema
}

// Fields of the Price.
func (Price) Fields() []ent.Field {
	return nil
}

// Edges of the Price.
func (Price) Edges() []ent.Edge {
	return nil
}
