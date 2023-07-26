package handlers

import (
	"github.com/eesoymilk/health-statistic-api/ent"
)

type Handler struct {
	DB *ent.Client
}
