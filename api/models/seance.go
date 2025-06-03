package models

import (
	"gorm.io/gorm"
	"time"
)

type Seance struct {
	ID                uint              `json:"id" gorm:"primaryKey"`
	Date              time.Time         `json:"date"`
	Type              string            `json:"type"`
	Comment           string            `json:"comment,omitempty"`
	ExercicesRealises []ExerciceRealise `json:"exercices_realises" gorm:"foreignKey:SeanceID"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type SeanceInput struct {
	Date    string `json:"date"` // Format YYYY-MM-DD
	Type    string `json:"type"` // Muscu, natation, etc.
	Comment string `json:"comment,omitempty"`
}
