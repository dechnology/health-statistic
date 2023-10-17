package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// HealthKit holds the schema definition for the HealthKit entity.
type HealthKit struct {
	ent.Schema
}

// Fields of the HealthKit.
func (HealthKit) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("data", map[string]interface{}{}),
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
