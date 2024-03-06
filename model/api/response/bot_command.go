package response

import "encoding/json"

/*
*
An array of structs of type @see TgBotCommand
*/
type TgBotCommands struct {
	Items []*TgBotCommand
}

/*
*
@link https://core.telegram.org/bots/api#botcommand
*/
type TgBotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

func (r *TgBotCommands) UnmarshalJSON(body []byte) error {
	var items []*TgBotCommand

	if err := json.Unmarshal(body, &items); err != nil {
		return err
	}

	r.Items = items

	return nil
}
