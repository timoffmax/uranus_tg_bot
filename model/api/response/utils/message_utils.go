package utils

import (
	"github.com/timoffmax/vocabulary-bot/model/api/response"
	"regexp"
)

func IsAnusMessage(message *response.TgMessage) bool {
	regex := regexp.MustCompile(`(?i)(anus|анус|очко)`)

	return regex.MatchString(message.Text)
}
