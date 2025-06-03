package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("⚠️ Aucun fichier .env trouvé (utilisation des variables système)")
	}
}
