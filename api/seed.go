package main

import (
	"app-sport/config"
	"app-sport/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func Init() {
	config.LoadEnv()
	config.ConnectDB()

	seedPoids()
	seedExercices()
	seedSeances()
}

func seedPoids() {
	poidsData := []models.Poids{
		{Date: parseDate("2025-05-05"), Valeur: 103.0},
		{Date: parseDate("2025-05-06"), Valeur: 102.9},
		{Date: parseDate("2025-05-07"), Valeur: 102.8},
		{Date: parseDate("2025-05-08"), Valeur: 102.7},
		{Date: parseDate("2025-05-09"), Valeur: 102.8},
	}
	for _, entry := range poidsData {
		exist := models.Poids{}
		if err := config.DB.Where("date = ?", entry.Date).First(&exist).Error; err == nil {
			// déjà présent
			continue
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			// autre erreur (ex: DB down)
			panic(err)
		}

		config.DB.Create(&entry)
		fmt.Println("Poids ajouté:", entry.Date, entry.Valeur)
	}
}

func seedExercices() {
	exercices := []models.Exercice{
		{Nom: "Tractions", GroupeMusculaire: "Dos", Type: "Musculation"},
		{Nom: "Dips", GroupeMusculaire: "Triceps", Type: "Musculation"},
		{Nom: "Pompes", GroupeMusculaire: "Pectoraux", Type: "Musculation"},
		{Nom: "Wrist Curl", GroupeMusculaire: "Avant-bras", Type: "Musculation"},
		{Nom: "Reverse Curl", GroupeMusculaire: "Avant-bras", Type: "Musculation"},
		{Nom: "Farmer Walk", GroupeMusculaire: "Grip", Type: "Musculation"},
	}
	for _, ex := range exercices {
		exist := models.Exercice{}
		if err := config.DB.Where("nom = ?", ex.Nom).First(&exist).Error; err == nil {
			continue
		}
		config.DB.Create(&ex)
		fmt.Println("Exercice ajouté:", ex.Nom)
	}
}

func seedSeances() {
	seance := models.Seance{
		Date:    parseDate("2025-06-01"),
		Type:    "Musculation",
		Comment: "Séance type tractions/dips/pompes",
	}
	config.DB.FirstOrCreate(&seance, models.Seance{Date: seance.Date, Type: seance.Type})

	exos := []struct {
		nom    string
		poids  float64
		reps   int
		series int
		rpe    int
	}{
		{"Tractions", 0, 3, 3, 7},
		{"Dips", 0, 8, 3, 8},
		{"Pompes", 0, 15, 3, 6},
	}
	for i, e := range exos {
		ex := models.Exercice{}
		if err := config.DB.Where("nom = ?", e.nom).First(&ex).Error; err != nil {
			fmt.Println("Exercice introuvable:", e.nom)
			continue
		}
		real := models.ExerciceRealise{
			SeanceID:   seance.ID,
			ExerciceID: ex.ID,
			Ordre:      i + 1,
			Series:     e.series,
			Reps:       e.reps,
			Poids:      e.poids,
			RPE:        float64(e.rpe),
			Comment:    "",
		}
		if err := config.DB.Create(&real).Error; err != nil {
			fmt.Println("❌ Erreur insertion exercice:", e.nom, err)
		} else {
			fmt.Println("✅ Exercice ajouté:", e.nom)
		}
	}

}

func parseDate(dateStr string) time.Time {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		panic("Date invalide: " + dateStr)
	}
	return t
}
