package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	RegisterPoidsRoutes(r) // délégué à routes/poids.go
	RegisterExerciceRoutes(r)
}
