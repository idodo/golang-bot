package base

import (
	"golang-bot/kaihela/api/base/event"
	"testing"
)

func TestParseFrame(t *testing.T) {
	f := event.ParseFrameMapByData([]byte(`{"d":"{}", "s":0, "sn":1}`))
	if f == nil {
		t.Error("f is nil")
	}
	if f.Data["type"] != "" {
		t.Error("type is not empty")
	}
}
