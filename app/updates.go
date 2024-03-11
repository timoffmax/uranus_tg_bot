package app

import (
	"fmt"
	"github.com/timoffmax/vocabulary-bot/database"
	"github.com/timoffmax/vocabulary-bot/database/entity"
	"github.com/timoffmax/vocabulary-bot/model/api/request/message"
	"github.com/timoffmax/vocabulary-bot/model/api/request/payload"
	"github.com/timoffmax/vocabulary-bot/model/api/response"
	"github.com/timoffmax/vocabulary-bot/model/api/response/utils"
	"github.com/timoffmax/vocabulary-bot/service/db"
	"time"
)

func listenForUpdates() {
	var offset int

	for {
		lastProcessedUpd := processRecentUpdates(offset)

		if nil != lastProcessedUpd {
			offset = lastProcessedUpd.UpdateId + 1
		}
	}
}

func processRecentUpdates(offset int) *response.TgUpdate {
	database.EstablishConnection()
	defer database.CloseConnection()

	var lastProcessedUpd *response.TgUpdate
	updates := getRecentUpdates(offset)

	for _, update := range updates {
		processUpdate(update)
		lastProcessedUpd = update
	}

	return lastProcessedUpd
}

func getRecentUpdates(offset int) []*response.TgUpdate {
	updates := message.NewGetUpdates(payload.TgPlGetUpdates{
		Offset:         offset,
		AllowedUpdates: []string{"message"},
	})

	respPayload := updates.ResponsePayload
	result := respPayload.Items

	return result
}

func processUpdate(upd *response.TgUpdate) {
	var origMessage response.TgMessage
	var origUser response.TgUser
	var origChat response.TgChat

	if 0 != upd.Message.MessageId {
		origMessage = upd.Message
		origUser = origMessage.From
		origChat = origMessage.Chat
	}

	if 0 == origUser.Id {
		return
	}

	if utils.IsAnusMessage(&origMessage) {
		var prevMessage = db.GetPreviousAnusMessage()

		if nil != prevMessage {
			var currentDate = time.Now().Unix()
			var prevDate = prevMessage.CreatedAt.Unix()
			var diffSecondsTotal = (currentDate - prevDate)
			var diffDays = diffSecondsTotal / 60 / 60 / 24
			var diffHours = (diffSecondsTotal / 60 / 60) % 24
			var diffMinutes = (diffSecondsTotal / 60) % 60
			var diffSeconds = diffSecondsTotal % 60

			requestPayload := payload.TgPlSendMessage{
				ChatId:    origChat.Id,
				ParseMode: payload.ParseModeHTML,
			}

			requestPayload.Text = fmt.Sprintf(
				"%d days %d hours %d minutes %d seconds without anus mentioning",
				uint64(diffDays),
				uint64(diffHours),
				uint64(diffMinutes),
				uint64(diffSeconds),
			)
			requestPayload.ReplyToMessageId = int(prevMessage.MessageId)
			message.NewSendMessage(requestPayload)
		}

		requestPayload := payload.TgPlSendMessage{
			ChatId:    origChat.Id,
			ParseMode: payload.ParseModeHTML,
		}
		requestPayload.Text = "I'm watching for ur anus"
		requestPayload.ReplyToMessageId = origMessage.MessageId
		message.NewSendMessage(requestPayload)

		anusMessage := entity.AnusMessage{
			MessageId: uint(origMessage.MessageId),
		}
		db.CreateAnusMessage(&anusMessage)
	}
}
