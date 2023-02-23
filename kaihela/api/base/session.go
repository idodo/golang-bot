package base

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gookit/event"
	log "github.com/sirupsen/logrus"
	event2 "golang-bot/kaihela/api/base/event"
	"io"
)

const EVENT_RECEIVE_FRAME = "EVENT-GLOBAL-RECEIVE_FRAME"

type Session struct {
	Compressed          int
	ReceiveFrameHandler func(frame *event2.FrameMap) error
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
	_, err := sonic.Get(data)
	if err != nil {
		log.Error("Json Unmarshal err.", err)
		return err
	}
	data2 := s.ProcessData(data)
	frame := event2.ParseFrameMapByData(data2)
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

func (s *Session) ProcessData(data []byte) []byte {
	return data
}

func (s *Session) ReceiveFrame(frame *event2.FrameMap) error {
	event.Trigger(EVENT_RECEIVE_FRAME, map[string]interface{}{"frame": frame})
	if frame.SignalType == event2.SIG_EVENT {
		eventType := frame.Data["type"]
		channelType := frame.Data["channel_type"]
		if eventType != "" {
			event.Trigger(fmt.Sprint("%s_%s", eventType, channelType), map[string]interface{}{"frame": frame})
		}
	}
	return nil

}
