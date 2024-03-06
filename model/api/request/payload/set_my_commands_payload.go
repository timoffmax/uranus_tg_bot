package payload

import "github.com/timoffmax/vocabulary-bot/model/api/response"

/*
*
@link https://core.telegram.org/bots/api#setmycommands
*/
type TgPlSetMyCommands struct {
	Commands []response.TgBotCommand `json:"commands"`
}
