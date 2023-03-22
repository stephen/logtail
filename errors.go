package logtail

import (
	"errors"
)

var (
	InvalidSourceToken = errors.New("you provided an invalid source token")
	InvalidBodyFormat  = errors.New("invalid body format:The body is not a valid JSON or MessagePack")
)
