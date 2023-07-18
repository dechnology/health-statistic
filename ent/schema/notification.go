package schema

import "entgo.io/ent"

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return nil
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return nil
}
