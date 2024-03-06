package app

import (
	"github.com/joho/godotenv"
	"github.com/timoffmax/vocabulary-bot/database"
	"log"
)

func Run() {
	initApp()
	listenForUpdates()
}

func initApp() {
	loadEnvFile()

	database.EstablishConnection()
	defer database.CloseConnection()
}

func initConfig() {
	loadEnvFile()
}

func initDB() {
	database.EstablishConnection()
	defer database.CloseConnection()
}

func loadEnvFile() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
