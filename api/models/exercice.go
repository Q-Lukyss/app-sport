package models

import (
	"gorm.io/gorm"
	"time"
)

type Exercice struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	Nom              string         `json:"nom" gorm:"uniqueIndex;not null"`
	GroupeMusculaire string         `json:"groupe_musculaire"`
	Type             string         `json:"type"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"-" gorm:"index"` // masqué du JSON Swagger
}

type ExerciceInput struct {
	Nom              string `json:"nom"`
	GroupeMusculaire string `json:"groupe_musculaire"`
	Type             string `json:"type"`
}
