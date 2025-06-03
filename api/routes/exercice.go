package routes

import (
	"app-sport/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterExerciceRoutes(r *gin.Engine) {
	r.GET("/exercices", controllers.GetAllExercices)
	r.GET("/exercices/:id", controllers.GetExerciceById)
	r.POST("/exercices", controllers.CreateExercice)
	r.PUT("/exercices/:id", controllers.UpdateExercice)
	r.DELETE("/exercices/:id", controllers.DeleteExercice)
}
