// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFcmToken holds the string denoting the fcm_token field in the database.
	FieldFcmToken = "fcm_token"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldBirthYear holds the string denoting the birth_year field in the database.
	FieldBirthYear = "birth_year"
	// FieldHeight holds the string denoting the height field in the database.
	FieldHeight = "height"
	// FieldWeight holds the string denoting the weight field in the database.
	FieldWeight = "weight"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldEducationLevel holds the string denoting the education_level field in the database.
	FieldEducationLevel = "education_level"
	// FieldOccupation holds the string denoting the occupation field in the database.
	FieldOccupation = "occupation"
	// FieldMarriage holds the string denoting the marriage field in the database.
	FieldMarriage = "marriage"
	// FieldMedicalHistory holds the string denoting the medical_history field in the database.
	FieldMedicalHistory = "medical_history"
	// FieldMedicationStatus holds the string denoting the medication_status field in the database.
	FieldMedicationStatus = "medication_status"
	// FieldDementedAmongDirectRelatives holds the string denoting the demented_among_direct_relatives field in the database.
	FieldDementedAmongDirectRelatives = "demented_among_direct_relatives"
	// FieldHeadInjuryExperience holds the string denoting the head_injury_experience field in the database.
	FieldHeadInjuryExperience = "head_injury_experience"
	// FieldEarCondition holds the string denoting the ear_condition field in the database.
	FieldEarCondition = "ear_condition"
	// FieldEyesightCondition holds the string denoting the eyesight_condition field in the database.
	FieldEyesightCondition = "eyesight_condition"
	// FieldSmokingHabit holds the string denoting the smoking_habit field in the database.
	FieldSmokingHabit = "smoking_habit"
	// FieldDataConsent holds the string denoting the data_consent field in the database.
	FieldDataConsent = "data_consent"
	// EdgeQuestionnaireResponses holds the string denoting the questionnaire_responses edge name in mutations.
	EdgeQuestionnaireResponses = "questionnaire_responses"
	// EdgeNotifications holds the string denoting the notifications edge name in mutations.
	EdgeNotifications = "notifications"
	// EdgePrices holds the string denoting the prices edge name in mutations.
	EdgePrices = "prices"
	// EdgeMycards holds the string denoting the mycards edge name in mutations.
	EdgeMycards = "mycards"
	// EdgeHealthkit holds the string denoting the healthkit edge name in mutations.
	EdgeHealthkit = "healthkit"
	// EdgeDeegoo holds the string denoting the deegoo edge name in mutations.
	EdgeDeegoo = "deegoo"
	// MyCardFieldID holds the string denoting the ID field of the MyCard.
	MyCardFieldID = "card_number"
	// Table holds the table name of the user in the database.
	Table = "users"
	// QuestionnaireResponsesTable is the table that holds the questionnaire_responses relation/edge.
	QuestionnaireResponsesTable = "questionnaire_responses"
	// QuestionnaireResponsesInverseTable is the table name for the QuestionnaireResponse entity.
	// It exists in this package in order to avoid circular dependency with the "questionnaireresponse" package.
	QuestionnaireResponsesInverseTable = "questionnaire_responses"
	// QuestionnaireResponsesColumn is the table column denoting the questionnaire_responses relation/edge.
	QuestionnaireResponsesColumn = "user_questionnaire_responses"
	// NotificationsTable is the table that holds the notifications relation/edge.
	NotificationsTable = "notifications"
	// NotificationsInverseTable is the table name for the Notification entity.
	// It exists in this package in order to avoid circular dependency with the "notification" package.
	NotificationsInverseTable = "notifications"
	// NotificationsColumn is the table column denoting the notifications relation/edge.
	NotificationsColumn = "user_notifications"
	// PricesTable is the table that holds the prices relation/edge.
	PricesTable = "prices"
	// PricesInverseTable is the table name for the Price entity.
	// It exists in this package in order to avoid circular dependency with the "price" package.
	PricesInverseTable = "prices"
	// PricesColumn is the table column denoting the prices relation/edge.
	PricesColumn = "user_prices"
	// MycardsTable is the table that holds the mycards relation/edge.
	MycardsTable = "my_cards"
	// MycardsInverseTable is the table name for the MyCard entity.
	// It exists in this package in order to avoid circular dependency with the "mycard" package.
	MycardsInverseTable = "my_cards"
	// MycardsColumn is the table column denoting the mycards relation/edge.
	MycardsColumn = "user_mycards"
	// HealthkitTable is the table that holds the healthkit relation/edge.
	HealthkitTable = "health_kits"
	// HealthkitInverseTable is the table name for the HealthKit entity.
	// It exists in this package in order to avoid circular dependency with the "healthkit" package.
	HealthkitInverseTable = "health_kits"
	// HealthkitColumn is the table column denoting the healthkit relation/edge.
	HealthkitColumn = "user_healthkit"
	// DeegooTable is the table that holds the deegoo relation/edge.
	DeegooTable = "deegoos"
	// DeegooInverseTable is the table name for the Deegoo entity.
	// It exists in this package in order to avoid circular dependency with the "deegoo" package.
	DeegooInverseTable = "deegoos"
	// DeegooColumn is the table column denoting the deegoo relation/edge.
	DeegooColumn = "user_deegoo"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldFcmToken,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldBirthYear,
	FieldHeight,
	FieldWeight,
	FieldGender,
	FieldEducationLevel,
	FieldOccupation,
	FieldMarriage,
	FieldMedicalHistory,
	FieldMedicationStatus,
	FieldDementedAmongDirectRelatives,
	FieldHeadInjuryExperience,
	FieldEarCondition,
	FieldEyesightCondition,
	FieldSmokingHabit,
	FieldDataConsent,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// BirthYearValidator is a validator for the "birth_year" field. It is called by the builders before save.
	BirthYearValidator func(int) error
	// HeightValidator is a validator for the "height" field. It is called by the builders before save.
	HeightValidator func(float64) error
	// WeightValidator is a validator for the "weight" field. It is called by the builders before save.
	WeightValidator func(float64) error
	// DefaultDataConsent holds the default value on creation for the "data_consent" field.
	DefaultDataConsent bool
)

// Gender defines the type for the "gender" enum field.
type Gender string

// Gender values.
const (
	GenderMale      Gender = "male"
	GenderFemale    Gender = "female"
	GenderNonbinary Gender = "nonbinary"
)

func (ge Gender) String() string {
	return string(ge)
}

// GenderValidator is a validator for the "gender" field enum values. It is called by the builders before save.
func GenderValidator(ge Gender) error {
	switch ge {
	case GenderMale, GenderFemale, GenderNonbinary:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for gender field: %q", ge)
	}
}

// EducationLevel defines the type for the "education_level" enum field.
type EducationLevel string

// EducationLevel values.
const (
	EducationLevelElementrySchoolOrBelow EducationLevel = "elementry_school_or_below"
	EducationLevelMiddleSchool           EducationLevel = "middle_school"
	EducationLevelHighSchool             EducationLevel = "high_school"
	EducationLevelBachelor               EducationLevel = "bachelor"
	EducationLevelMaster                 EducationLevel = "master"
	EducationLevelDoctorate              EducationLevel = "doctorate"
)

func (el EducationLevel) String() string {
	return string(el)
}

// EducationLevelValidator is a validator for the "education_level" field enum values. It is called by the builders before save.
func EducationLevelValidator(el EducationLevel) error {
	switch el {
	case EducationLevelElementrySchoolOrBelow, EducationLevelMiddleSchool, EducationLevelHighSchool, EducationLevelBachelor, EducationLevelMaster, EducationLevelDoctorate:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for education_level field: %q", el)
	}
}

// Occupation defines the type for the "occupation" enum field.
type Occupation string

// Occupation values.
const (
	OccupationStudent             Occupation = "student"
	OccupationGovernmentEmployee  Occupation = "government_employee"
	OccupationServiceIndustry     Occupation = "service_industry"
	OccupationIndustryAndCommerce Occupation = "industry_and_commerce"
	OccupationFreelancer          Occupation = "freelancer"
	OccupationDomestic            Occupation = "domestic"
	OccupationRetired             Occupation = "retired"
	OccupationOthers              Occupation = "others"
)

func (o Occupation) String() string {
	return string(o)
}

// OccupationValidator is a validator for the "occupation" field enum values. It is called by the builders before save.
func OccupationValidator(o Occupation) error {
	switch o {
	case OccupationStudent, OccupationGovernmentEmployee, OccupationServiceIndustry, OccupationIndustryAndCommerce, OccupationFreelancer, OccupationDomestic, OccupationRetired, OccupationOthers:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for occupation field: %q", o)
	}
}

// Marriage defines the type for the "marriage" enum field.
type Marriage string

// Marriage values.
const (
	MarriageSingle   Marriage = "single"
	MarriageMarried  Marriage = "married"
	MarriageDivorced Marriage = "divorced"
	MarriageWidowed  Marriage = "widowed"
)

func (m Marriage) String() string {
	return string(m)
}

// MarriageValidator is a validator for the "marriage" field enum values. It is called by the builders before save.
func MarriageValidator(m Marriage) error {
	switch m {
	case MarriageSingle, MarriageMarried, MarriageDivorced, MarriageWidowed:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for marriage field: %q", m)
	}
}

// MedicalHistory defines the type for the "medical_history" enum field.
type MedicalHistory string

// MedicalHistory values.
const (
	MedicalHistoryHighBloodPressure MedicalHistory = "high_blood_pressure"
	MedicalHistoryHyperlipidemia    MedicalHistory = "hyperlipidemia"
	MedicalHistoryDiabetes          MedicalHistory = "diabetes"
	MedicalHistoryHeartDisease      MedicalHistory = "heart_disease"
	MedicalHistoryStroke            MedicalHistory = "stroke"
	MedicalHistoryMentalIllness     MedicalHistory = "mental_illness"
	MedicalHistoryDementia          MedicalHistory = "dementia"
	MedicalHistoryNoneOfTheAbove    MedicalHistory = "none_of_the_above"
)

func (mh MedicalHistory) String() string {
	return string(mh)
}

// MedicalHistoryValidator is a validator for the "medical_history" field enum values. It is called by the builders before save.
func MedicalHistoryValidator(mh MedicalHistory) error {
	switch mh {
	case MedicalHistoryHighBloodPressure, MedicalHistoryHyperlipidemia, MedicalHistoryDiabetes, MedicalHistoryHeartDisease, MedicalHistoryStroke, MedicalHistoryMentalIllness, MedicalHistoryDementia, MedicalHistoryNoneOfTheAbove:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for medical_history field: %q", mh)
	}
}

// MedicationStatus defines the type for the "medication_status" enum field.
type MedicationStatus string

// MedicationStatus values.
const (
	MedicationStatusCardiovascularDrugs MedicationStatus = "cardiovascular_drugs"
	MedicationStatusPsychiatricDrugs    MedicationStatus = "psychiatric_drugs"
	MedicationStatusOtherDrugs          MedicationStatus = "other_drugs"
	MedicationStatusNoDrugs             MedicationStatus = "no_drugs"
)

func (ms MedicationStatus) String() string {
	return string(ms)
}

// MedicationStatusValidator is a validator for the "medication_status" field enum values. It is called by the builders before save.
func MedicationStatusValidator(ms MedicationStatus) error {
	switch ms {
	case MedicationStatusCardiovascularDrugs, MedicationStatusPsychiatricDrugs, MedicationStatusOtherDrugs, MedicationStatusNoDrugs:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for medication_status field: %q", ms)
	}
}

// EarCondition defines the type for the "ear_condition" enum field.
type EarCondition string

// EarCondition values.
const (
	EarConditionNormal                        EarCondition = "normal"
	EarConditionSlightlyAffectingConversation EarCondition = "slightly_affecting_conversation"
	EarConditionNeedHearingAid                EarCondition = "need_hearing_aid"
)

func (ec EarCondition) String() string {
	return string(ec)
}

// EarConditionValidator is a validator for the "ear_condition" field enum values. It is called by the builders before save.
func EarConditionValidator(ec EarCondition) error {
	switch ec {
	case EarConditionNormal, EarConditionSlightlyAffectingConversation, EarConditionNeedHearingAid:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for ear_condition field: %q", ec)
	}
}

// EyesightCondition defines the type for the "eyesight_condition" enum field.
type EyesightCondition string

// EyesightCondition values.
const (
	EyesightConditionNormal                   EyesightCondition = "normal"
	EyesightConditionSlightlyAffectingReading EyesightCondition = "slightly_affecting_reading"
	EyesightConditionNeedGlasses              EyesightCondition = "need_glasses"
)

func (ec EyesightCondition) String() string {
	return string(ec)
}

// EyesightConditionValidator is a validator for the "eyesight_condition" field enum values. It is called by the builders before save.
func EyesightConditionValidator(ec EyesightCondition) error {
	switch ec {
	case EyesightConditionNormal, EyesightConditionSlightlyAffectingReading, EyesightConditionNeedGlasses:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for eyesight_condition field: %q", ec)
	}
}

// SmokingHabit defines the type for the "smoking_habit" enum field.
type SmokingHabit string

// SmokingHabit values.
const (
	SmokingHabitNone      SmokingHabit = "none"
	SmokingHabitSometimes SmokingHabit = "sometimes"
	SmokingHabitEveryday  SmokingHabit = "everyday"
)

func (sh SmokingHabit) String() string {
	return string(sh)
}

// SmokingHabitValidator is a validator for the "smoking_habit" field enum values. It is called by the builders before save.
func SmokingHabitValidator(sh SmokingHabit) error {
	switch sh {
	case SmokingHabitNone, SmokingHabitSometimes, SmokingHabitEveryday:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for smoking_habit field: %q", sh)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFcmToken orders the results by the fcm_token field.
func ByFcmToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFcmToken, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByBirthYear orders the results by the birth_year field.
func ByBirthYear(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBirthYear, opts...).ToFunc()
}

// ByHeight orders the results by the height field.
func ByHeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHeight, opts...).ToFunc()
}

// ByWeight orders the results by the weight field.
func ByWeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWeight, opts...).ToFunc()
}

// ByGender orders the results by the gender field.
func ByGender(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGender, opts...).ToFunc()
}

// ByEducationLevel orders the results by the education_level field.
func ByEducationLevel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEducationLevel, opts...).ToFunc()
}

// ByOccupation orders the results by the occupation field.
func ByOccupation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOccupation, opts...).ToFunc()
}

// ByMarriage orders the results by the marriage field.
func ByMarriage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMarriage, opts...).ToFunc()
}

// ByMedicalHistory orders the results by the medical_history field.
func ByMedicalHistory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMedicalHistory, opts...).ToFunc()
}

// ByMedicationStatus orders the results by the medication_status field.
func ByMedicationStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMedicationStatus, opts...).ToFunc()
}

// ByDementedAmongDirectRelatives orders the results by the demented_among_direct_relatives field.
func ByDementedAmongDirectRelatives(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDementedAmongDirectRelatives, opts...).ToFunc()
}

// ByHeadInjuryExperience orders the results by the head_injury_experience field.
func ByHeadInjuryExperience(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHeadInjuryExperience, opts...).ToFunc()
}

// ByEarCondition orders the results by the ear_condition field.
func ByEarCondition(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEarCondition, opts...).ToFunc()
}

// ByEyesightCondition orders the results by the eyesight_condition field.
func ByEyesightCondition(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEyesightCondition, opts...).ToFunc()
}

// BySmokingHabit orders the results by the smoking_habit field.
func BySmokingHabit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSmokingHabit, opts...).ToFunc()
}

// ByDataConsent orders the results by the data_consent field.
func ByDataConsent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDataConsent, opts...).ToFunc()
}

// ByQuestionnaireResponsesCount orders the results by questionnaire_responses count.
func ByQuestionnaireResponsesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newQuestionnaireResponsesStep(), opts...)
	}
}

// ByQuestionnaireResponses orders the results by questionnaire_responses terms.
func ByQuestionnaireResponses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newQuestionnaireResponsesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByNotificationsCount orders the results by notifications count.
func ByNotificationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNotificationsStep(), opts...)
	}
}

// ByNotifications orders the results by notifications terms.
func ByNotifications(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNotificationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPricesCount orders the results by prices count.
func ByPricesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPricesStep(), opts...)
	}
}

// ByPrices orders the results by prices terms.
func ByPrices(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPricesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMycardsCount orders the results by mycards count.
func ByMycardsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMycardsStep(), opts...)
	}
}

// ByMycards orders the results by mycards terms.
func ByMycards(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMycardsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByHealthkitCount orders the results by healthkit count.
func ByHealthkitCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newHealthkitStep(), opts...)
	}
}

// ByHealthkit orders the results by healthkit terms.
func ByHealthkit(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHealthkitStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDeegooCount orders the results by deegoo count.
func ByDeegooCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDeegooStep(), opts...)
	}
}

// ByDeegoo orders the results by deegoo terms.
func ByDeegoo(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDeegooStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newQuestionnaireResponsesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(QuestionnaireResponsesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, QuestionnaireResponsesTable, QuestionnaireResponsesColumn),
	)
}
func newNotificationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NotificationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, NotificationsTable, NotificationsColumn),
	)
}
func newPricesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PricesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PricesTable, PricesColumn),
	)
}
func newMycardsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MycardsInverseTable, MyCardFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MycardsTable, MycardsColumn),
	)
}
func newHealthkitStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HealthkitInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, HealthkitTable, HealthkitColumn),
	)
}
func newDeegooStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DeegooInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DeegooTable, DeegooColumn),
	)
}
