package logtail

type Log struct {
	DateTimeOffSet int64          `json:"dt"`
	Message        string         `json:"message"`
	Level          string         `json:"level"`
	Context        map[string]any `json:"context"`
}
