package models

import (
	"github.com/google/uuid"
)

type User struct {
	BaseModel
	Username string    `gorm:"not null;unique"`
	Email    string    `gorm:"not null;unique"`
	Password string    `gorm:"not null"`
	RoleID   uuid.UUID `gorm:"not null"`

	Role      *Role     `gorm:"foreignKey:RoleID"`
	Addresses []Address `gorm:"foreignKey:UserID"`
	Todos     []Todo    `gorm:"foreignKey:UserID"`
}
