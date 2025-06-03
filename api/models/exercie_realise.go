package models

type ExerciceRealise struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	SeanceID   uint    `json:"seance_id"`
	ExerciceID uint    `json:"exercice_id"`
	Ordre      int     `json:"ordre"`
	Series     int     `json:"series"`
	Reps       int     `json:"reps"`
	Poids      float64 `json:"poids"`
	RPE        float64 `json:"rpe"`
	Comment    string  `json:"comment,omitempty"`

	Exercice Exercice `json:"exercice" gorm:"foreignKey:ExerciceID"`
}

type ExerciceRealiseInput struct {
	ExerciceID uint    `json:"exercice_id"`
	Ordre      int     `json:"ordre"`
	Series     int     `json:"series"`
	Reps       int     `json:"reps"`
	Poids      float64 `json:"poids"`
	RPE        float64 `json:"rpe"`
	Comment    string  `json:"comment,omitempty"`
}
