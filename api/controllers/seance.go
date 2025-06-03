package controllers

import (
	"app-sport/config"
	"app-sport/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// CreateSeance godoc
// @Summary Créer une séance
// @Description Ajoute une nouvelle séance
// @Tags Séances
// @Accept json
// @Produce json
// @Param seance body models.SeanceInput true "Séance à enregistrer"
// @Success 201 {object} models.Seance
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /seances [post]
func CreateSeance(c *gin.Context) {
	var input models.SeanceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	date, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format de date invalide (YYYY-MM-DD)"})
		return
	}

	seance := models.Seance{
		Date:    date,
		Type:    input.Type,
		Comment: input.Comment,
	}

	if err := config.DB.Create(&seance).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, seance)
}

// GetAllSeances godoc
// @Summary Liste des séances
// @Description Retourne toutes les séances enregistrées
// @Tags Séances
// @Produce json
// @Success 200 {array} models.Seance
// @Failure 500 {object} map[string]string
// @Router /seances [get]
func GetAllSeances(c *gin.Context) {
	var seances []models.Seance
	if err := config.DB.Order("date desc").Find(&seances).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, seances)
}

// GetSeanceById godoc
// @Summary Détail d'une séance
// @Description Retourne une séance par son ID
// @Tags Séances
// @Produce json
// @Param id path int true "ID de la séance"
// @Success 200 {object} models.Seance
// @Failure 404 {object} map[string]string
// @Router /seances/{id} [get]
func GetSeanceById(c *gin.Context) {
	var seance models.Seance
	if err := config.DB.First(&seance, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Séance non trouvée"})
		return
	}
	c.JSON(http.StatusOK, seance)
}

// GetSeancesToday godoc
// @Summary Séances du jour
// @Description Récupère toutes les séances prévues pour la date du jour
// @Tags Séances
// @Produce json
// @Success 200 {array} models.Seance
// @Failure 500 {object} map[string]string
// @Router /seances/today [get]
func GetSeancesToday(c *gin.Context) {
	today := time.Now().Format("2006-01-02")

	var seances []models.Seance
	if err := config.DB.Where("date = ?", today).Preload("ExercicesRealises").Preload("ExercicesRealises.Exercice").Find(&seances).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, seances)
}

// GetSeancesByDate godoc
// @Summary Séances par date
// @Description Récupère toutes les séances pour une date spécifique
// @Tags Séances
// @Produce json
// @Param date query string true "Date au format YYYY-MM-DD"
// @Success 200 {array} models.Seance
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /seances/by-date [get]
func GetSeancesByDate(c *gin.Context) {
	dateStr := c.Query("date")
	if dateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Paramètre 'date' requis (YYYY-MM-DD)"})
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format de date invalide. Utilisez YYYY-MM-DD"})
		return
	}

	var seances []models.Seance
	if err := config.DB.Where("date = ?", date).Preload("ExercicesRealises").Preload("ExercicesRealises.Exercice").Find(&seances).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, seances)
}

// UpdateSeance godoc
// @Summary Modifier une séance
// @Description Met à jour une séance existante
// @Tags Séances
// @Accept json
// @Produce json
// @Param id path int true "ID de la séance"
// @Param seance body models.SeanceInput true "Champs à modifier"
// @Success 200 {object} models.Seance
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /seances/{id} [put]
func UpdateSeance(c *gin.Context) {
	var seance models.Seance
	if err := config.DB.First(&seance, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Séance non trouvée"})
		return
	}
	var input models.SeanceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	date, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format de date invalide (YYYY-MM-DD)"})
		return
	}
	seance.Date = date
	seance.Type = input.Type
	seance.Comment = input.Comment

	config.DB.Save(&seance)
	c.JSON(http.StatusOK, seance)
}

// DeleteSeance godoc
// @Summary Supprimer une séance
// @Description Supprime une séance par ID
// @Tags Séances
// @Produce json
// @Param id path int true "ID de la séance"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /seances/{id} [delete]
func DeleteSeance(c *gin.Context) {
	var seance models.Seance
	if err := config.DB.First(&seance, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Séance non trouvée"})
		return
	}
	config.DB.Delete(&seance)
	c.JSON(http.StatusOK, gin.H{"message": "Séance supprimée"})
}
