package models

type Role struct {
	BaseModel
	Name string `gorm:"not null;unique"`
}
