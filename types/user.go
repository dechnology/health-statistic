package types

import "github.com/eesoymilk/health-statistic-api/ent"

type BaseUser struct {
	ID                           string  `json:"id"`
	BirthYear                    int     `json:"birth_year" example:"2000"`
	Height                       float64 `json:"height" example:"185.2"`
	Weight                       float64 `json:"weight" example:"80.3"`
	Gender                       string  `json:"gender" example:"male"`
	EducationLevel               string  `json:"education_level" example:"doctorate" enums:"elementry_school_or_below,middle_school,high_school,bachelor,master,doctorate"`
	Occupation                   string  `json:"occupation" example:"student" enums:"student,government_employee,service_industry,industry_and_commerce,freelancer,domestic"`
	Marriage                     string  `json:"marriage" example:"single" enums:"single,married,divorced,widowed"`
	MedicalHistory               string  `json:"medical_history" example:"none"`
	MedicationStatus             string  `json:"medication_status" example:"none"`
	DementedAmongDirectRelatives bool    `json:"demented_among_direct_relatives" example:"false"`
	HeadInjuryExperience         bool    `json:"head_injury_experience" example:"false"`
	EarCondition                 string  `json:"ear_condition" example:"normal" enums:"normal,slightly_affecting_conversation,need_hearing_aid"`
	EyesightCondition            string  `json:"eyesight_condition" example:"normal" enums:"normal,slightly_affecting_reading,need_glasses"`
	SmokingHabit                 string  `json:"smoking_habit" example:"none" enums:"none,sometimes,everyday"`
	DataConsent                  bool    `json:"data_consent" example:"true"`
}

type RegisterData struct {
	User     BaseUser                    `json:"user"`
	Response ResponseWithQuestionnaireId `json:"response"`
}

type RegisterResponse struct {
	User     *ent.User                  `json:"user"`
	Response *ent.QuestionnaireResponse `json:"response"`
}

type FcmTokenRequest struct {
	FcmToken string `json:"fcm_token"`
}
