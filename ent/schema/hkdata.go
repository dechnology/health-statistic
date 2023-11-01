package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// HKData holds the schema definition for the HKData entity.
type HKData struct {
	ent.Schema
}

// Fields of the HKData.
func (HKData) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Unique().
			Default(uuid.New),
		// you need to manually set the data id value
		field.String("data_id"),
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
