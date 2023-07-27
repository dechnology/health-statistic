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
		field.String("id").
			MinLen(12).
			MaxLen(12).
			Unique().
			Immutable().
			StorageKey("card_number"),
		field.String("card_password").
			MinLen(12).
			MaxLen(12).
			Unique().
			Immutable(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("taken_at").
			Optional().
			Nillable(),
	}
}

// Edges of the MyCard.
func (MyCard) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("recipient", User.Type).
			Ref("mycards").
			Unique(),
		edge.To("notifications", Notification.Type),
	}
}
