package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserContext struct {
	ID       uuid.UUID   `json:"id"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Role     RoleContext `json:"role"`
	jwt.RegisteredClaims
}

type RoleContext struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
