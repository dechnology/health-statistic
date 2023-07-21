package types

import (
	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/google/uuid"
)

//	@Description	The json body for creating a new question.
type BaseQuestion struct {
	// The question body
	Body string `json:"body"`
	// The question type, currently we accept string but in the future this
	// field will be enums.
	Type string `json:"type"`
}

//	@Description	The json body for creating a new answer in a new response.
type BaseAnswer struct {
	// The answer body.
	Body string `json:"body"`
	// The question this answer relates to, the question also needs to be in
	// the same questionnaire as the response.
	QuestionId uuid.UUID `json:"question_id"`
}

//	@Description	BaseQuestionnaire
type BaseQuestionnaire struct {
	// The name of the questionnaire
	Name string `json:"name"`
	// The initial questions in this questionnaire. This field may be empty
	// and you can add questions later using post request to
	// `quesionnaires/:id/new/question`.
	Questions []*BaseQuestion `json:"questions"`
}

type BaseResponse struct {
	// The user ID of the user who submit the response.
	UserId string `json:"user_id"`
	// The answers to all questions in a questionnaire.
	Answers []*BaseAnswer `json:"answers"`
}

type Response struct {
	ent.QuestionnaireResponse
	UserId  uuid.UUID     `json:"user_id"`
	Answers []*BaseAnswer `json:"answers"`
}

type QuestionWithQuestionnaire struct {
	ent.Question
	Questionnaire ent.Questionnaire
}

type ResponseWithQuestionnaire struct {
	Response
	Questionnaire ent.Questionnaire
}
