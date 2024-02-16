package models

import (
	"time"

	"github.com/google/uuid"
)

func (uuidBase *UuidBase) BeforeCreate() (err error) {
	uuidBase.UUID = uuid.New().String()
	return
}

type UuidBase struct {
	UUID      string    `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

type Base struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

// User has many posts
// User has one user_details
type User struct {
	Base
	Posts      []Post
	UserDetail UserDetail
}

type UserDetail struct {
	UuidBase
	UserID       uint   `gorm:"not null;index"`
	UserName     string `gorm:"not null"`
	Email        string `gorm:"not null;unique"`
	ProfileImage string
}

// Post belongs to users
type Post struct {
	UuidBase
	UserID uint   `gorm:"not null;index"`
	User   User   `gorm:"foreignKey:UserID;references:ID"`
	Active bool   `gorm:"not null;default:true"`
	Body   string `gorm:"not null"`
}
