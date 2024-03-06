package database

import "github.com/timoffmax/vocabulary-bot/database/entity"

func init() {
	EstablishConnection()
}

/*
Applies DB schema changes
*/
func Migrate() {
	err := DbConnection.AutoMigrate(
		entity.AnusMessage{},
	)

	if nil != err {
		panic(err)
	}
}
