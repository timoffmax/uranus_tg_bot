package database

import (
	"fmt"
	"github.com/timoffmax/vocabulary-bot/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbConnection *gorm.DB

func init() {
	EstablishConnection()
	Migrate()
}

/*
*
Opens the DB connection
*/
func EstablishConnection() {
	var err error

	if nil == DbConnection {
		dsn := getDSN()

		DbConnection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			panic("Can't connect to the database")
		}
	}
}

func CloseConnection() {
	sqlDB, err := DbConnection.DB()

	if nil != err {
		sqlDB.Close()
	}
}

/*
*
Prepares DSN string
*/
func getDSN() string {
	host := config.GetDbHost()
	port := config.GetDbPort()
	dbName := config.GetDbName()
	dbUser := config.GetDbUser()
	dbPasswd := config.GetDbPassword()
	SSLMode := "disable"
	timezone := "UTC"

	result := fmt.Sprintf(
		"host=%s user=%s password='%s' dbname=%s port=%s sslmode=%s TimeZone=%s",
		host,
		dbUser,
		dbPasswd,
		dbName,
		port,
		SSLMode,
		timezone,
	)

	return result
}
