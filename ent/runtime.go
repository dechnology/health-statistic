// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/eesoymilk/health-statistic-api/ent/answer"
	"github.com/eesoymilk/health-statistic-api/ent/questionnaire"
	"github.com/eesoymilk/health-statistic-api/ent/questionnaireresponse"
	"github.com/eesoymilk/health-statistic-api/ent/schema"
	"github.com/eesoymilk/health-statistic-api/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	answerFields := schema.Answer{}.Fields()
	_ = answerFields
	// answerDescCreatedAt is the schema descriptor for created_at field.
	answerDescCreatedAt := answerFields[0].Descriptor()
	// answer.DefaultCreatedAt holds the default value on creation for the created_at field.
	answer.DefaultCreatedAt = answerDescCreatedAt.Default.(func() time.Time)
	questionnaireFields := schema.Questionnaire{}.Fields()
	_ = questionnaireFields
	// questionnaireDescCreatedAt is the schema descriptor for created_at field.
	questionnaireDescCreatedAt := questionnaireFields[1].Descriptor()
	// questionnaire.DefaultCreatedAt holds the default value on creation for the created_at field.
	questionnaire.DefaultCreatedAt = questionnaireDescCreatedAt.Default.(func() time.Time)
	questionnaireresponseFields := schema.QuestionnaireResponse{}.Fields()
	_ = questionnaireresponseFields
	// questionnaireresponseDescCreatedAt is the schema descriptor for created_at field.
	questionnaireresponseDescCreatedAt := questionnaireresponseFields[0].Descriptor()
	// questionnaireresponse.DefaultCreatedAt holds the default value on creation for the created_at field.
	questionnaireresponse.DefaultCreatedAt = questionnaireresponseDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[1].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[2].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescBirthYear is the schema descriptor for birth_year field.
	userDescBirthYear := userFields[3].Descriptor()
	// user.BirthYearValidator is a validator for the "birth_year" field. It is called by the builders before save.
	user.BirthYearValidator = userDescBirthYear.Validators[0].(func(int) error)
	// userDescHeight is the schema descriptor for height field.
	userDescHeight := userFields[4].Descriptor()
	// user.HeightValidator is a validator for the "height" field. It is called by the builders before save.
	user.HeightValidator = userDescHeight.Validators[0].(func(float64) error)
	// userDescWeight is the schema descriptor for weight field.
	userDescWeight := userFields[5].Descriptor()
	// user.WeightValidator is a validator for the "weight" field. It is called by the builders before save.
	user.WeightValidator = userDescWeight.Validators[0].(func(float64) error)
	// userDescMedicalHistory is the schema descriptor for medical_history field.
	userDescMedicalHistory := userFields[10].Descriptor()
	// user.MedicalHistoryValidator is a validator for the "medical_history" field. It is called by the builders before save.
	user.MedicalHistoryValidator = userDescMedicalHistory.Validators[0].(func(string) error)
	// userDescMedicationStatus is the schema descriptor for medication_status field.
	userDescMedicationStatus := userFields[11].Descriptor()
	// user.MedicationStatusValidator is a validator for the "medication_status" field. It is called by the builders before save.
	user.MedicationStatusValidator = userDescMedicationStatus.Validators[0].(func(string) error)
}
