package interfaces

import tg_api "github.com/timoffmax/vocabulary-bot/model/api"

type ResponseMessageInterface interface {
	SetText(text string)
	SetButtons(buttons tg_api.TgReplyMarkupInterface)
	SetParseMode(mode string)
}
