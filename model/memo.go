package model

import (
	"time"

	"gorm.io/gorm"
)

type Memo struct {
	gorm.Model
	Title   string `json:"title" gorm:"not null"`
	Content string `json:"content"`
	User    User   `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId  uint   `json:"user_id" gorm:"not null"`
}

type MemoResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
