package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Default(uuid.New).
			Unique(),
		field.Enum("type").
			Values("normal", "mycard", "price"),
		field.Time("sent_at").
			Optional(),
		field.Time("read_at").
			Optional(),
		field.Text("message"),
	}
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("recipient", User.Type).
			Ref("notifications").
			Unique(),
		edge.From("mycard", MyCard.Type).
			Ref("notifications").
			Unique(),
		edge.From("price", Price.Type).
			Ref("notifications").
			Unique(),
	}
}
