package handlers

import (
	"github.com/eesoymilk/health-statistic-api/ent"
)

type RegisterHandler struct {
	DB *ent.Client
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

type NotificationHandler struct {
	DB *ent.Client
}

type MyCardHandler struct {
	DB *ent.Client
}

type PriceHandler struct {
	DB *ent.Client
}
