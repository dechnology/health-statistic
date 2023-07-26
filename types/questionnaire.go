package types

import (
	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/google/uuid"
)

type BaseQuestion struct {
	// The question body
	Body    string   `json:"body" example:"你這週的心情如何？"`
	Type    string   `json:"type" enums:"short_answer,single_choice,multiple_choice" example:"single_choice"`
	Choices []string `json:"choices" example:"[是,否]"`
}

type BaseAnswer struct {
	// The question this answer relates to, the question also needs to be in
	// the same questionnaire as the response.
	QuestionId uuid.UUID `json:"question_id"`
	// The answer body.
	Body    *string       `json:"body" example:"我這週心情還不錯！"`
	Choices *[]ent.Choice `json:"choices"`
}

//	@Description	BaseQuestionnaire
type BaseQuestionnaire struct {
	// The name of the questionnaire
	Name string `json:"name" example:"問卷標題"`
	// The initial questions in this questionnaire. This field may be empty
	// and you can add questions later using post request to
	// `quesionnaires/:id/new/question`.
	Questions []*BaseQuestion `json:"questions"`
}

type BaseResponse struct {
	// The answers to all questions in a questionnaire.
	Answers []*BaseAnswer `json:"answers"`
}

type ResponseWithUserId struct {
	// The user ID of the user who submit the response.
	UserId string `json:"user_id"`
	BaseResponse
}

type ResponseWithQuestionnaireId struct {
	QuestionnaireId string `json:"quesionnaire_id"`
	BaseResponse
}

type Response struct {
	ent.QuestionnaireResponse
	ResponseWithUserId
}

type ResponseWithQuestionnaire struct {
	Response
	Questionnaire ent.Questionnaire
}

type QuestionWithQuestionnaire struct {
	ent.Question
	Questionnaire ent.Questionnaire
}

type QuestionnaireWithId struct {
	BaseQuestionnaire
	ID string `json:"id"`
}

type QuestionnaireWithQuestions struct {
	ent.Questionnaire
	Questions []ent.Question `json:"questions"`
}

type QuestionnaireDetails struct {
	ent.Questionnaire
	Questions []ent.Question `json:"questions"`
	Responses []Response     `json:"responses"`
}
