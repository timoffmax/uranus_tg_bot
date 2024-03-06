package error

type UserFriendlyError interface {
	GetFriendlyError() string
}

type LoggableError interface {
	GetLoggableError() string
}
