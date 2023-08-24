package application

import (
	"log"

	"github.com/team2/real_api/app/server"
	"github.com/team2/real_api/config"
	"github.com/joho/godotenv"
)

func Start() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := config.LoadConfig()
	server := server.NewServer(config)
	error := server.Start()
	if error != nil {
		log.Fatal("Error starting server: ", error)
	}
}
