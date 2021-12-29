package command

import (
	"github.com/valist-io/valist/log"
)

type contextKey string

const (
	ClientKey = contextKey("client")
	ConfigKey = contextKey("config")
)

var logger = log.New()
