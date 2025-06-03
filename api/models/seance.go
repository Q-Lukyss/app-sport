package models

import (
	"gorm.io/gorm"
	"time"
)

type Seance struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Date      time.Time      `json:"date"`              // jour de la séance
	Type      string         `json:"type"`              // muscu, natation, cardio...
	Comment   string         `json:"comment,omitempty"` // facultatif
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type SeanceInput struct {
	Date    string `json:"date"` // "YYYY-MM-DD"
	Type    string `json:"type"` // muscu, piscine, etc.
	Comment string `json:"comment,omitempty"`
}
