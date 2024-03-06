package command

import (
	"github.com/timoffmax/vocabulary-bot/model/api/request"
	"github.com/timoffmax/vocabulary-bot/model/api/request/payload"
	"github.com/timoffmax/vocabulary-bot/model/api/response"
)

const RouteSetMyCommands = "setMyCommands"

type SetMyCommands struct {
	Payload         payload.TgPlSetMyCommands
	ResponsePayload *response.AbstractBoolResponse
	Request         request.AuthorizedRequest
}

func NewSetMyCommands(payload payload.TgPlSetMyCommands) *SetMyCommands {
	reqRespInfo := &SetMyCommands{
		Request: request.AuthorizedRequest{
			ApiMethod:          RouteSetMyCommands,
			RequestMethod:      request.RequestMethodPOST,
			RequestContentType: request.ContentTypeJson,
			Payload:            payload,
		},
		ResponsePayload: &response.AbstractBoolResponse{},
	}

	reqRespInfo.send()

	return reqRespInfo
}

func (service *SetMyCommands) send() *response.AbstractBoolResponse {
	service.Request.Send(service.ResponsePayload)

	return service.ResponsePayload
}
