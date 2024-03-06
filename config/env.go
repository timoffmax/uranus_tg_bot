package config

import (
	"github.com/joho/godotenv"
	"os"
)

const tgBotToken = "TG_BOT_TOKEN"

const dbHost = "DB_HOST"
const dbPort = "DB_PORT"
const dbUser = "DB_USER"
const dbPassword = "DB_PASSWORD"
const dbName = "DB_NAME"

func init() {
	loadConfig()
}

func loadConfig() {
	godotenv.Load()
}

func GetTgBotToken() string {
	return os.Getenv(tgBotToken)
}

func GetDbHost() string {
	return os.Getenv(dbHost)
}

func GetDbPort() string {
	return os.Getenv(dbPort)
}

func GetDbUser() string {
	return os.Getenv(dbUser)
}

func GetDbPassword() string {
	return os.Getenv(dbPassword)
}

func GetDbName() string {
	return os.Getenv(dbName)
}
