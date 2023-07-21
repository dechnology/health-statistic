package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MyCard holds the schema definition for the MyCard entity.
type MyCard struct {
	ent.Schema
}

// Fields of the MyCard.
func (MyCard) Fields() []ent.Field {
	return []ent.Field{
		field.String("card_number").MinLen(12).MaxLen(12),
		field.String("card_password").MinLen(12).MaxLen(12),
		field.Time("created_at").Default(time.Now),
		field.Time("taken_at").Optional(),
	}
}

// Edges of the MyCard.
func (MyCard) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("recipient", User.Type).
			Unique(),
		edge.To("notifications", Notification.Type),
	}
}
