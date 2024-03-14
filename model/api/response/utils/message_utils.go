package utils

import (
	"github.com/timoffmax/vocabulary-bot/model/api/response"
	"regexp"
)

func IsAnusMessage(message *response.TgMessage) bool {
	regex := regexp.MustCompile(`(?i)(очко|([a@ая])([^a-zA-Z0-9А-Яа-я]+?)?([nнh])([^a-zA-Z0-9А-Яа-я]+?)?([uуy])([^a-zA-Z0-9А-Яа-я]+?)?([sсc$]))`)

	return regex.MatchString(message.Text)
}
