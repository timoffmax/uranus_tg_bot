package service_api

import (
	"github.com/timoffmax/vocabulary-bot/model/api"
	"math"
)

type ButtonSettings struct {
	Label    string
	Callback string
}

func CreateSingleButton(label string, callback string) [][]api.TgInlineKeyboardButton {
	result := make([][]api.TgInlineKeyboardButton, 1)

	buttonsRow := make([]api.TgInlineKeyboardButton, 1)
	buttonsRow[0] = api.TgInlineKeyboardButton{
		Text:         label,
		CallbackData: callback,
	}
	result[0] = buttonsRow

	return result
}

func CreateTwoButtonsInline(l1 string, c1 string, l2 string, c2 string) [][]api.TgInlineKeyboardButton {
	result := make([][]api.TgInlineKeyboardButton, 1)

	buttonsRow := make([]api.TgInlineKeyboardButton, 2)
	buttonsRow[0] = api.TgInlineKeyboardButton{
		Text:         l1,
		CallbackData: c1,
	}
	buttonsRow[1] = api.TgInlineKeyboardButton{
		Text:         l2,
		CallbackData: c2,
	}

	result[0] = buttonsRow

	return result
}

func CreateButtons(buttons []ButtonSettings, qtyPerRow int) [][]api.TgInlineKeyboardButton {
	qtyButtons := len(buttons)
	qtyRows := int(math.Ceil(float64(qtyButtons) / float64(qtyPerRow)))
	currentIdx := 0
	maxIdx := qtyButtons - 1

	result := make([][]api.TgInlineKeyboardButton, qtyRows)

	for i := 0; i < qtyRows; i++ {
		buttonsRow := make([]api.TgInlineKeyboardButton, qtyPerRow)

		for j := 0; j < qtyPerRow; j++ {
			if currentIdx > maxIdx {
				break
			}

			buttonSettings := buttons[currentIdx]
			buttonsRow[j] = api.TgInlineKeyboardButton{
				Text:         buttonSettings.Label,
				CallbackData: buttonSettings.Callback,
			}

			currentIdx++
		}

		result[i] = buttonsRow
	}

	return result
}
