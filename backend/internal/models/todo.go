package models

import "github.com/google/uuid"

type Todo struct {
	BaseModel
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`

	UserID uuid.UUID
}
