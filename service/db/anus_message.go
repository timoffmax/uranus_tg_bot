package db

import (
	"errors"
	"github.com/timoffmax/vocabulary-bot/database"
	"github.com/timoffmax/vocabulary-bot/database/entity"
	"gorm.io/gorm"
)

func GetPreviousAnusMessage(chatId string) *entity.AnusMessage {
	result := &entity.AnusMessage{ChatId: chatId}

	conn := database.DbConnection
	qr := conn.Last(&result)
	notFound := errors.Is(qr.Error, gorm.ErrRecordNotFound)

	if (1 != qr.RowsAffected) || (true == notFound) {
		return nil
	}

	return result
}

func CreateAnusMessage(anusMessage *entity.AnusMessage) error {
	var err error

	conn := database.DbConnection
	queryResult := conn.Create(anusMessage)

	if nil != queryResult.Error {
		err = queryResult.Error
	}

	return err
}
