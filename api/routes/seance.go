package routes

import (
	"app-sport/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterSeanceRoutes(r *gin.Engine) {

	r.GET("/seances", controllers.GetAllSeances)
	r.GET("/seances/today", controllers.GetSeancesToday)
	r.GET("/seances/by-date", controllers.GetSeancesByDate)
	r.GET("/seances/:id", controllers.GetSeanceById)
	r.POST("/seances", controllers.CreateSeance)
	r.PUT("/seances/:id", controllers.UpdateSeance)
	r.DELETE("/seances/:id", controllers.DeleteSeance)
}
