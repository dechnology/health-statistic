package handlers

import (
	"net/url"

	"firebase.google.com/go/v4/messaging"
	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/google/uuid"
)

type Handler struct {
	DB                          *ent.Client
	FCM                         *messaging.Client
	RegistrationQuestionnaireId uuid.UUID
	Auth0Issuer                 url.URL
	Auth0Audience               string
}
