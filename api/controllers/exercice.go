package controllers

import (
	"app-sport/config"
	"app-sport/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateExercice godoc
// @Summary Créer un exercice
// @Description Ajoute un nouvel exercice dans la base
// @Tags Exercices
// @Accept json
// @Produce json
// @Param exercice body models.ExerciceInput true "Exercice à créer"
// @Success 201 {object} models.Exercice
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /exercices [post]
func CreateExercice(c *gin.Context) {
	var input models.ExerciceInput

	// Liaison du JSON reçu
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exercice := models.Exercice{
		Nom:              input.Nom,
		GroupeMusculaire: input.GroupeMusculaire,
		Type:             input.Type,
	}

	// Enregistrement en base
	if err := config.DB.Create(&exercice).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, exercice)
}

// GetAllExercices godoc
// @Summary Liste des exercices
// @Description Retourne tous les exercices disponibles
// @Tags Exercices
// @Produce json
// @Success 200 {array} models.Exercice
// @Failure 500 {object} map[string]string
// @Router /exercices [get]
func GetAllExercices(c *gin.Context) {
	var exercices []models.Exercice

	if err := config.DB.Order("created_at desc").Find(&exercices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exercices)
}

// GetExerciceById godoc
// @Summary Détail d’un exercice
// @Description Retourne les informations d’un exercice par son ID
// @Tags Exercices
// @Produce json
// @Param id path int true "ID de l’exercice"
// @Success 200 {object} models.Exercice
// @Failure 404 {object} map[string]string
// @Router /exercices/{id} [get]
func GetExerciceById(c *gin.Context) {
	var exercice models.Exercice

	if err := config.DB.First(&exercice, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exercice non trouvé"})
		return
	}

	c.JSON(http.StatusOK, exercice)
}

// UpdateExercice godoc
// @Summary Modifier un exercice
// @Description Met à jour un exercice existant
// @Tags Exercices
// @Accept json
// @Produce json
// @Param id path int true "ID de l’exercice"
// @Param exercice body models.ExerciceInput true "Champs à modifier"
// @Success 200 {object} models.Exercice
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /exercices/{id} [put]
func UpdateExercice(c *gin.Context) {
	var exercice models.Exercice

	if err := config.DB.First(&exercice, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exercice non trouvé"})
		return
	}

	var input models.ExerciceInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exercice.Nom = input.Nom
	exercice.GroupeMusculaire = input.GroupeMusculaire
	exercice.Type = input.Type

	config.DB.Save(&exercice)

	c.JSON(http.StatusOK, exercice)
}

// DeleteExercice godoc
// @Summary Supprimer un exercice
// @Description Supprime un exercice par son ID
// @Tags Exercices
// @Produce json
// @Param id path int true "ID de l’exercice"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /exercices/{id} [delete]
func DeleteExercice(c *gin.Context) {
	var exercice models.Exercice

	if err := config.DB.First(&exercice, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exercice non trouvé"})
		return
	}

	config.DB.Delete(&exercice)
	c.JSON(http.StatusOK, gin.H{"message": "Exercice supprimé"})
}
