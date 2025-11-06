package dto

import (
	"time"

	"github.com/google/uuid"
)

type TimeRequest struct {
	StartTime string `form:"startTime" validate:"required"`
	EndTime   string `form:"endTime" validate:"required"`
}

type TimeResponse struct {
	UUID      uuid.UUID
	StartTime string
	EndTime   string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
