package controllers

import (
	"app-sport/config"
	"app-sport/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateExerciceRealise godoc
// @Summary Ajouter un exercice réalisé à une séance
// @Description Ajoute un exercice avec ses performances à une séance existante
// @Tags Exercices Réalisés
// @Accept json
// @Produce json
// @Param id path int true "ID de la séance"
// @Param input body models.ExerciceRealiseInput true "Détails de l'exercice"
// @Success 201 {object} models.ExerciceRealise
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /seances/{id}/exercices [post]
func CreateExerciceRealise(c *gin.Context) {
	seanceID, _ := strconv.Atoi(c.Param("id"))

	var input models.ExerciceRealiseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	er := models.ExerciceRealise{
		SeanceID:   uint(seanceID),
		ExerciceID: input.ExerciceID,
		Ordre:      input.Ordre,
		Series:     input.Series,
		Reps:       input.Reps,
		Poids:      input.Poids,
		RPE:        input.RPE,
		Comment:    input.Comment,
	}

	if err := config.DB.Create(&er).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, er)
}

// GetExercicesBySeance godoc
// @Summary Liste des exercices réalisés pour une séance
// @Description Retourne tous les exercices associés à une séance
// @Tags Exercices Réalisés
// @Produce json
// @Param id path int true "ID de la séance"
// @Success 200 {array} models.ExerciceRealise
// @Failure 500 {object} map[string]string
// @Router /seances/{id}/exercices [get]
func GetExercicesBySeance(c *gin.Context) {
	seanceID := c.Param("id")
	var exercices []models.ExerciceRealise

	if err := config.DB.Preload("Exercice").Where("seance_id = ?", seanceID).Order("ordre asc").Find(&exercices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exercices)
}

// UpdateExerciceRealise godoc
// @Summary Modifier un exercice réalisé
// @Description Met à jour les données d’un exercice réalisé
// @Tags Exercices Réalisés
// @Accept json
// @Produce json
// @Param id path int true "ID de l'exercice réalisé"
// @Param input body models.ExerciceRealiseInput true "Données à modifier"
// @Success 200 {object} models.ExerciceRealise
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /exercice-realise/{id} [put]
func UpdateExerciceRealise(c *gin.Context) {
	var er models.ExerciceRealise
	if err := config.DB.First(&er, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exercice non trouvé"})
		return
	}

	var input models.ExerciceRealiseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	er.ExerciceID = input.ExerciceID
	er.Ordre = input.Ordre
	er.Series = input.Series
	er.Reps = input.Reps
	er.Poids = input.Poids
	er.RPE = input.RPE
	er.Comment = input.Comment

	config.DB.Save(&er)
	c.JSON(http.StatusOK, er)
}

// DeleteExerciceRealise godoc
// @Summary Supprimer un exercice réalisé
// @Description Supprime une entrée exercice réalisé par son ID
// @Tags Exercices Réalisés
// @Produce json
// @Param id path int true "ID de l'exercice réalisé"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /exercice-realise/{id} [delete]
func DeleteExerciceRealise(c *gin.Context) {
	var er models.ExerciceRealise
	if err := config.DB.First(&er, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exercice non trouvé"})
		return
	}

	config.DB.Delete(&er)
	c.JSON(http.StatusOK, gin.H{"message": "Supprimé"})
}
