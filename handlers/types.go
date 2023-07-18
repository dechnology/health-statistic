package handlers

import (
	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/google/uuid"
)

type QuestionBody struct {
	Body string `json:"body"`
	Type string `json:"type"`
}

type AnswerBody struct {
	Body       string    `json:"body"`
	QuestionId uuid.UUID `json:"question_id"`
}

type QuestionnaireResponseBody struct {
	UserId  string        `json:"user_id"`
	Answers []*AnswerBody `json:"answers"`
}

type QuestionnaireBody struct {
	Name      string          `json:"name"`
	Questions []*QuestionBody `json:"questions"`
}

type Question struct {
	ent.Question
	Questionnaires []*ent.Questionnaire `json:"quesionnaires"`
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
