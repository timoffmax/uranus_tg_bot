package command

import (
	"github.com/timoffmax/vocabulary-bot/model/api/request"
	"github.com/timoffmax/vocabulary-bot/model/api/request/payload"
	"github.com/timoffmax/vocabulary-bot/model/api/response"
)

const RouteGetMyCommands = "getMyCommands"

type GetMyCommands struct {
	Payload         payload.TgPlSetMyCommands
	ResponsePayload *response.TgBotCommands
	Request         request.AuthorizedRequest
}

func NewGetMyCommands() *GetMyCommands {
	reqRespInfo := &GetMyCommands{
		Request: request.AuthorizedRequest{
			ApiMethod:     RouteGetMyCommands,
			RequestMethod: request.RequestMethodGET,
			Payload:       nil,
		},
		ResponsePayload: &response.TgBotCommands{},
	}

	reqRespInfo.send()

	return reqRespInfo
}

func (service *GetMyCommands) send() *response.TgBotCommands {
	service.Request.Send(service.ResponsePayload)

	return service.ResponsePayload
}
