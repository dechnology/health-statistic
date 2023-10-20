package schema

import (
	"errors"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
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
		field.Int("birth_year").
			Positive().
			Validate(func(y int) error {
				if time.Now().Year()-y < 45 {
					return errors.New("the user must be 45 years old or above")
				}
				return nil
			}),
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
			"retired",
			"others",
		),
		field.Enum("marriage").Values(
			"single",
			"married",
			"divorced",
			"widowed",
		),
		field.Enum("medical_history").Values(
			"high_blood_pressure",
			"hyperlipidemia",
			"diabetes",
			"heart_disease",
			"stroke",
			"mental_illness",
			"dementia",
			"none_of_the_above",
		),
		field.Enum("medication_status").Values(
			"cardiovascular_drugs",
			"psychiatric_drugs",
			"other_drugs",
			"no_drugs",
		),
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
		edge.To("questionnaire_responses", QuestionnaireResponse.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("notifications", Notification.Type).
			Annotations(entsql.OnDelete(entsql.SetNull)),
		edge.To("prices", Price.Type),
		edge.To("mycards", MyCard.Type).Required(),
		edge.To("healthkit", HealthKit.Type),
		edge.To("deegoo", Deegoo.Type),
	}
}
