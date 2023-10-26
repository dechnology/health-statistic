package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// HKData holds the schema definition for the HKData entity.
type HKData struct {
	ent.Schema
}

// Fields of the HKData.
func (HKData) Fields() []ent.Field {
	return []ent.Field{
		// you need to manually set the id value
		field.String("id").
			Immutable().
			Unique(),
		field.String("type"),
		field.String("value"),
		field.String("start_timestamp"),
		field.String("end_timestamp"),
		field.String("timezone_id"),
	}
}

// Edges of the HKData.
func (HKData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("healthkit", HealthKit.Type).
			Ref("data").
			Unique(),
	}
}
