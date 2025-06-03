package models

import (
	"time"
)

type Poids struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Date      time.Time `gorm:"uniqueIndex" json:"date"`
	Valeur    float64   `json:"valeur"` // en kg
	CreatedAt time.Time
	UpdatedAt time.Time
}

// PoidsInput est utilisé pour la création et la modification
type PoidsInput struct {
	Date   string  `json:"date"`   // format YYYY-MM-DD
	Valeur float64 `json:"valeur"` // en kg
}

// PoidsUpdate est utilisé pour modifier uniquement la valeur
type PoidsUpdate struct {
	Valeur float64 `json:"valeur"`
}
