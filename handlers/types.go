package handlers

import (
	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/google/uuid"
)

//	@Description	The json body for creating a new question.
type QuestionBody struct {
	// The question body
	Body string `json:"body"`
	// The question type, currently we accept string but in the future this
	// field will be enums.
	Type string `json:"type"`
}

//	@Description	The json body for creating a new answer in a new response.
type AnswerBody struct {
	// The answer body.
	Body string `json:"body"`
	// The question this answer relates to, the question also needs to be in
	// the same questionnaire as the response.
	QuestionId uuid.UUID `json:"question_id"`
}

// ===
type BaseQuestion struct {
}

type BaseAnswer struct {
}

type BaseQuestionnaire struct {
	// The name of the questionnaire
	Name string `json:"name"`
	// The initial questions in this questionnaire. This field may be empty
	// and you can add questions later using post request to
	// `quesionnaires/:id/new/question`.
	Questions []*QuestionBody `json:"questions"`
}

type BaseResponse struct {
}

// ===

//	@Description	The json body for creating a new response.
type QuestionnaireResponseBody struct {
	// The user ID of the user who submit the response.
	UserId string `json:"user_id"`
	// The answers to all questions in a questionnaire.
	Answers []*AnswerBody `json:"answers"`
}

//	@Description	The json body for creating a new questionnaire.
type QuestionnaireBody struct {
	// The name of the questionnaire
	Name string `json:"name"`
	// The initial questions in this questionnaire. This field may be empty
	// and you can add questions later using post request to
	// `quesionnaires/:id/new/question`.
	Questions []*QuestionBody `json:"questions"`
}

type Question struct {
	ent.Question
	Questionnaires []*ent.Questionnaire `json:"quesionnaires"`
}

type SingleQuestionnaireResponse struct {
	ent.QuestionnaireResponse
	Questionnaire *Questionnaire `json:"questionnaire"`
	Answers       []*ent.Answer  `json:"answers"`
}

type QuestionnaireResponse struct {
	ent.QuestionnaireResponse
	Answers []*ent.Answer `json:"answers"`
}

type Questionnaire struct {
	ent.Questionnaire
	Responses *[]QuestionnaireResponse
	Questions *[]ent.Question `json:"questions"`
}

type UserHandler struct {
	DB *ent.Client
}

type QuestionnaireHandler struct {
	DB *ent.Client
}

type ResponseHandler struct {
	DB *ent.Client
}

type QuestionHandler struct {
	DB *ent.Client
}
