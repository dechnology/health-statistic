package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Reward holds the schema definition for the Reward entity.
type Reward struct {
	ent.Schema
}

// Fields of the Reward.
func (Reward) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").Values("mycard", "line"),
		field.String("code"),
		field.Time("created_at").Default(time.Now),
		field.Time("taken_at").Optional(),
	}
}

// Edges of the Reward.
func (Reward) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("recipient", User.Type).
			Unique(),
		edge.To("notifications", Notification.Type),
	}
}
