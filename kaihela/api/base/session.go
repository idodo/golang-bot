package base

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"github.com/gookit/event"
	log "github.com/sirupsen/logrus"
	event2 "golang-bot/kaihela/api/base/event"
	"io"
)

const EVENT_RECEIVE_FRAME = "EVENT-GLOBAL-RECEIVE_FRAME"
const EventDataFrameKey = "frame"
const EventDataSessionKey = "session"

type Session struct {
	Compressed          int
	ReceiveFrameHandler func(frame *event2.FrameMap) error
	ProcessDataHandler  func(data ast.Node) []byte
}

func (s *Session) On(message string, handler event.Listener) {
	event.On(message, handler)
}
func (s *Session) Trigger(eventName string, params event.M) {
	event.Trigger(eventName, params)
}
func (s *Session) ReceiveData(data []byte) error {
	if s.Compressed == 1 {
		b := bytes.NewReader(data)
		r, err := zlib.NewReader(b)
		if err != nil {
			return err
		}

		data, err = io.ReadAll(r)
		if err != nil {
			log.Error(err)
			return err
		}

	}
	dataNode, err := sonic.Get(data)
	if err != nil {
		log.Error("Json Unmarshal err.", err)
		return err
	}
	if s.ProcessDataHandler != nil {
		data = s.ProcessDataHandler(dataNode)
	}
	frame := event2.ParseFrameMapByData(data)
	log.WithField("frame", frame).Info("Receive frame from server")
	if frame != nil {
		if s.ReceiveFrameHandler != nil {
			return s.ReceiveFrameHandler(frame)
		} else {
			return s.ReceiveFrame(frame)
		}
	} else {
		log.Warnf("数据不是合法的frame", string(data))
	}
	return nil

}

func (s *Session) ReceiveFrame(frame *event2.FrameMap) error {
	event.Trigger(EVENT_RECEIVE_FRAME, map[string]interface{}{"frame": frame})
	if frame.SignalType == event2.SIG_EVENT {
		eventType := frame.Data["type"]
		channelType := frame.Data["channel_type"].(string)
		if eventType != "" {
			name := fmt.Sprintf("%s_%d", channelType, int64(eventType.(float64)))
			event.Trigger(name, map[string]interface{}{EventDataFrameKey: frame, EventDataSessionKey: s})
		}
	}
	return nil

}
