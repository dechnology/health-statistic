package handlers

import (
	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/google/uuid"
)

type QuestionnaireHandler struct {
	DB *ent.Client
}

type ResponseHandler struct {
	DB *ent.Client
}

type QuestionHandler struct {
	DB *ent.Client
}

type Question struct {
	Body string `json:"body"`
	Type string `json:"type"`
}

type Answer struct {
	QuestionId uuid.UUID `json:"question_id"`
	Body       string    `json:"body"`
}

type QuestionnaireBody struct {
	Name      string     `json:"name"`
	Questions []Question `json:"questions"`
}

type ResponseBody struct {
	UserId  string   `json:"user_id"`
	Answers []Answer `json:"answers"`
}
