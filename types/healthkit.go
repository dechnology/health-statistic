package types

import (
	"time"

	"github.com/google/uuid"
)

type BaseHealthKit struct {
	ID        uuid.UUID `json:"id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	StepCount float64   `json:"step_count"`
}
