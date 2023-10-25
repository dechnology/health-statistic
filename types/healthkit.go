package types

import (
	"time"

	"github.com/google/uuid"
)

type BaseHKData struct {
	ID             uuid.UUID `json:"id"`
	Type           string    `json:"type"`
	Value          string    `json:"value"`
	StartTimestamp string    `json:"start_timestamp"`
	EndTimestamp   string    `json:"end_timestamp"`
	TimezoneID     string    `json:"timezone_id"`
}

type BaseHealthKit struct {
	StartTime time.Time    `json:"start_time"`
	EndTime   time.Time    `json:"end_time"`
	Data      []BaseHKData `json:"data"`
}
