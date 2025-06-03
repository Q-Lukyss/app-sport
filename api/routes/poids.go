package routes

import (
	"app-sport/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterPoidsRoutes(r *gin.Engine) {
	r.GET("/poids", controllers.GetAllPoids)
	r.POST("/poids", controllers.CreatePoids)
	r.PUT("/poids/:id", controllers.UpdatePoids)
	r.DELETE("/poids/:id", controllers.DeletePoids)
}
