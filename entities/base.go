package entities

import (
	"time"
)

type Base struct {
	ID        UniqueEntityID `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
}
