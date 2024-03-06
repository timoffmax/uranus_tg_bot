package error

import "fmt"

type notFoundError struct {
	EntityName     string
	ErrorToDisplay string
}

func NewNotFoundError(entityName string, errorText string) *notFoundError {
	result := &notFoundError{
		EntityName:     entityName,
		ErrorToDisplay: errorText,
	}

	return result
}

func (e *notFoundError) Error() string {
	return e.ErrorToDisplay
}

func (e *notFoundError) GetFriendlyError() string {
	return e.Error()
}

func (e *notFoundError) GetLoggableError() string {
	result := fmt.Sprintf(
		"Entity %s not found. Details: %s",
		e.EntityName,
		e.ErrorToDisplay,
	)

	return result
}
