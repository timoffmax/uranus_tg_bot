package utils

import "net/http"

type SimpleResponse struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

var successStatuses = []int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}

func NewSimpleResponse(statusCode int, message string) SimpleResponse {
	success := inArray(statusCode, successStatuses)

	result := SimpleResponse{
		Success: success,
		Status:  statusCode,
		Message: message,
	}

	return result
}

func inArray(needle int, haystack []int) bool {
	result := false

	for _, item := range haystack {
		if needle == item {
			result = true
			break
		}
	}

	return result
}
