package message

import (
	"github.com/timoffmax/vocabulary-bot/model/api/request"
	"github.com/timoffmax/vocabulary-bot/model/api/request/payload"
	"github.com/timoffmax/vocabulary-bot/model/api/response"
)

const RouteAnswerCallbackQuery = "answerCallbackQuery"

type AnswerCallbackQuery struct {
	Payload         payload.TgPlSendMessage
	ResponsePayload bool
	Request         request.AuthorizedRequest
}

func NewAnswerCallbackQuery(payload payload.TgPlAnswerCallbackQuery) *SendMessage {
	reqRespInfo := &SendMessage{
		Request: request.AuthorizedRequest{
			ApiMethod:          RouteAnswerCallbackQuery,
			RequestMethod:      request.RequestMethodPOST,
			RequestContentType: request.ContentTypeJson,
			Payload:            payload,
		},
		ResponsePayload: &response.TgMessage{},
	}

	reqRespInfo.send()

	return reqRespInfo
}

func (service *AnswerCallbackQuery) send() bool {
	service.Request.Send(service.ResponsePayload)

	return service.ResponsePayload
}
