package logtail

type ContentType string

const (
	MsgPack ContentType = "application/msgpack"
	Json    ContentType = "application/json; charset=UTF-8"
)
