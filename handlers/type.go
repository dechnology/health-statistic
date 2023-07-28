package handlers

import (
	"net/url"

	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/google/uuid"
)

type Handler struct {
	DB                          *ent.Client
	RegistrationQuestionnaireId uuid.UUID
	Auth0Issuer                 url.URL
	Auth0Audience               string
}
