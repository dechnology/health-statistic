package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			Immutable(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("first_name").MaxLen(30),
		field.String("last_name").MaxLen(30),
		field.Int("birth_year").
			Positive(),
		field.Float("height").
			Positive(),
		field.Float("weight").
			Positive(),
		field.Enum("gender").Values(
			"male",
			"female",
			"nonbinary",
		),
		field.Enum("education_level").Values(
			"elementry_school_or_below",
			"middle_school",
			"high_school",
			"bachelor",
			"master",
			"doctorate",
		),
		field.Enum("occupation").Values(
			"student",
			"government_employee",
			"service_industry",
			"industry_and_commerce",
			"freelancer",
			"domestic",
		),
		field.Enum("marriage").Values(
			"single",
			"married",
			"divorced",
			"widowed",
		),
		field.String("medical_history").
			MaxLen(100).
			Optional(),
		field.String("medication_status").
			MaxLen(100).
			Optional(),
		field.Bool("demented_among_direct_relatives"),
		field.Bool("head_injury_experience"),
		field.Enum("ear_condition").Values(
			"normal",
			"slightly_affecting_conversation",
			"need_hearing_aid",
		),
		field.Enum("eyesight_condition").Values(
			"normal",
			"slightly_affecting_reading",
			"need_glasses",
		),
		field.Enum("smoking_habit").Values(
			"none",
			"sometimes",
			"everyday",
		),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("questionnaire_responses", QuestionnaireResponse.Type),
		edge.To("notifications", Notification.Type),
	}
}
