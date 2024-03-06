package error

import "fmt"

type argumentError struct {
	CommandName    string
	Arguments      []string
	ErrorToDisplay string
}

func NewArgumentError(commandName string, arguments []string, errorText string) *argumentError {
	result := &argumentError{
		CommandName:    commandName,
		Arguments:      arguments,
		ErrorToDisplay: errorText,
	}

	return result
}

func (e *argumentError) Error() string {
	return e.ErrorToDisplay
}

func (e *argumentError) GetFriendlyError() string {
	return e.Error()
}

func (e *argumentError) GetLoggableError() string {
	result := fmt.Sprintf(
		"Error executing command %s with arguments %v. Details: %s",
		e.CommandName,
		e.Arguments,
		e.ErrorToDisplay,
	)

	return result
}
