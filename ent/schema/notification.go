package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.Time("sent_at").Optional(),
		field.Time("read_at").Optional(),
		field.Text("message"),
	}
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("recipient", User.Type).Ref("notifications"),
		edge.From("reward", Reward.Type).
			Ref("notifications"),
		edge.From("price", Price.Type).
			Ref("notifications"),
	}
}
