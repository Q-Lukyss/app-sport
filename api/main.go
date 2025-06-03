// @title API App-Sport
// @version 1.0
// @description API de suivi de poids et séances d'entraînement
// @termsOfService http://swagger.io/terms/

// @contact.name Quentin Lachery
// @contact.email quentin.lkss@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8020
// @BasePath /
package main

import (
	"app-sport/config"
	"app-sport/docs"
	"app-sport/models"
	"app-sport/routes"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	err := config.DB.AutoMigrate(&models.Poids{})
	if err != nil {
		log.Println("Une erreur de Migration s'est produite")
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Println("✅ Serveur démarré sur le port 8020")

	// Swagger setup
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Toutes tes routes API (comme /poids)
	routes.SetupRoutes(r)

	r.Run(":8020")
}
