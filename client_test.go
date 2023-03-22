package logtail

import (
	"os"
	"testing"
)

func TestClient_Send(t *testing.T) {
	type FakeLog struct {
		User string `json:"user"`
		Msg  string `json:"msg"`
	}
	authToken := os.Getenv("LOGTAIL_AUTH_TOKEN")
	client := NewClient(authToken)
	log := []FakeLog{
		{
			User: "gopher",
			Msg:  "Coding Happy 🐱‍🚀",
		},
		{
			User: "czyt",
			Msg:  "I ❤ golang",
		},
	}
	n, err := client.Send(log)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("send log:", n)
}
