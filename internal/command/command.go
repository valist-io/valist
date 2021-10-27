package command

type contextKey string

const (
	ClientKey = contextKey("client")
	ConfigKey = contextKey("config")
)
