package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// HealthKit holds the schema definition for the HealthKit entity.
type HealthKit struct {
	ent.Schema
}

// Fields of the HealthKit.
func (HealthKit) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Default(uuid.New).
			Unique(),
		field.Time("start_date"),
		field.Time("end_date"),
		field.Float("step_count").Min(0),
	}
}

// Edges of the HealthKit.
func (HealthKit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("healthkit").
			Unique(),
	}
}
