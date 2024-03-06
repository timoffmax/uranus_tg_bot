package other

import (
	"github.com/timoffmax/vocabulary-bot/model/api/request"
	"github.com/timoffmax/vocabulary-bot/model/api/response"
)

const RouteGetMe = "getMe"

type GetMe struct {
	Payload         interface{}
	ResponsePayload *response.TgUser
	Request         request.AuthorizedRequest
}

func NewGetMe() *GetMe {
	result := &GetMe{
		Request: request.AuthorizedRequest{
			ApiMethod:     RouteGetMe,
			RequestMethod: request.RequestMethodGET,
			Payload:       nil,
		},
		ResponsePayload: &response.TgUser{},
	}

	return result
}

func (service *GetMe) Send() *response.TgUser {
	service.Request.Send(service.ResponsePayload)

	return service.ResponsePayload
}
