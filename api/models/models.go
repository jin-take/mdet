package models

import (
	"time"
)

type Base struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

type User struct {
	Base
	UUID string `gorm:"not null;unique;default:uuid()"`
}
