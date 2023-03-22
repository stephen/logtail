package logtail

type Marshaler func(data any) ([]byte, error)
