package utils

import (
	"github.com/timoffmax/vocabulary-bot/model/api/response"
	"regexp"
)

func IsAnusMessage(message *response.TgMessage) bool {
	regex := regexp.MustCompile(`(?i)(очко|([a@ая])([^a-zA-Z0-9]+?)?([nнh])([^a-zA-Z0-9]+?)?([uуy])([^a-zA-Z0-9]+?)?([sсc$]))`)

	return regex.MatchString(message.Text)
}
