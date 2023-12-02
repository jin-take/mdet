package models

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 論理削除を行うモデルはこちらを使用する
// Gormはdeleted_atがカラムにあると自動的に論理削除を行う
// https://gorm.io/docs/delete.html#Soft-Delete
type GormBase struct {
	ID        string         `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type User struct {
	Base
}
