package controllers

import (
	"app-sport/config"
	"app-sport/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// CreatePoids godoc
// @Summary Créer une entrée de poids
// @Description Enregistre une nouvelle valeur de poids pour une date donnée
// @Tags Poids
// @Accept json
// @Produce json
// @Param poids body models.PoidsInput true "Poids à enregistrer"
// @Success 201 {object} models.Poids
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /poids [post]
func CreatePoids(c *gin.Context) {
	var input struct {
		Date   string  `json:"date"`   // "2025-06-03"
		Valeur float64 `json:"valeur"` // kg
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format de date invalide (YYYY-MM-DD)"})
		return
	}

	poids := models.Poids{Date: date, Valeur: input.Valeur}
	if err := config.DB.Create(&poids).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, poids)
}

// GetAllPoids godoc
// @Summary Liste des poids
// @Description Retourne toutes les entrées de poids enregistrées, triées par date décroissante
// @Tags Poids
// @Produce json
// @Success 200 {array} models.Poids
// @Failure 500 {object} map[string]string
// @Router /poids [get]
func GetAllPoids(c *gin.Context) {
	var poids []models.Poids
	if err := config.DB.Order("date desc").Find(&poids).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, poids)
}

// UpdatePoids godoc
// @Summary Modifier un poids existant
// @Description Met à jour la valeur du poids pour une entrée donnée
// @Tags Poids
// @Accept json
// @Produce json
// @Param id path int true "ID de l'entrée poids"
// @Param poids body models.PoidsUpdate true "Nouvelle valeur du poids"
// @Success 200 {object} models.Poids
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /poids/{id} [put]
func UpdatePoids(c *gin.Context) {
	var poids models.Poids
	if err := config.DB.First(&poids, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entrée non trouvée"})
		return
	}

	var input struct {
		Valeur float64 `json:"valeur"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	poids.Valeur = input.Valeur
	config.DB.Save(&poids)
	c.JSON(http.StatusOK, poids)
}

// DeletePoids godoc
// @Summary Supprimer une entrée poids
// @Description Supprime une entrée de poids par son ID
// @Tags Poids
// @Produce json
// @Param id path int true "ID de l'entrée poids"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /poids/{id} [delete]
func DeletePoids(c *gin.Context) {
	var poids models.Poids
	if err := config.DB.First(&poids, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entrée non trouvée"})
		return
	}
	config.DB.Delete(&poids)
	c.JSON(http.StatusOK, gin.H{"message": "Supprimé"})
}
