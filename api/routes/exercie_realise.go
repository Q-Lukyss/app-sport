package routes

import (
	"app-sport/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterExerciceRealiseRoutes(r *gin.Engine) {
	r.POST("/seances/:id/exercices", controllers.CreateExerciceRealise)
	r.GET("/seances/:id/exercices", controllers.GetExercicesBySeance)
	r.PUT("/exercice-realise/:id", controllers.UpdateExerciceRealise)
	r.DELETE("/exercice-realise/:id", controllers.DeleteExerciceRealise)
}
