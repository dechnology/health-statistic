package handlers

import (
	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/google/uuid"
)

type Handler struct {
	DB                          *ent.Client
	RegistrationQuestionnaireId uuid.UUID
}
