package main

import (
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gookit/event"
	log "github.com/sirupsen/logrus"
	"golang-bot/kaihela/api/base"
	event2 "golang-bot/kaihela/api/base/event"
	"golang-bot/kaihela/api/helper"
)

const (
	Token   = "1/MTU1Njk=/b1r8dGxLxugSmcTop5jdjg=="
	BaseUrl = "https://www.kookapp.cn/api"
)

type ReceiveFrameHandler struct {
}

func (rf *ReceiveFrameHandler) Handle(e event.Event) error {
	log.WithField("event", e).Info("ReceiveFrameHandler receive frame.")
	return nil
}

type GroupEventHandler struct {
}

func (ge *GroupEventHandler) Handle(e event.Event) error {
	log.WithField("event", e).Info("GroupEventHandler receive event.")
	return nil
}

type GroupTextEventHandler struct {
}

func (gteh *GroupTextEventHandler) Handle(e event.Event) error {
	log.WithField("event", fmt.Sprintf("%+v", e.Data())).Info("收到频道内的文字消息.")
	err := func() error {
		if _, ok := e.Data()["frame"]; !ok {
			return errors.New("data has no frame field")
		}
		frame := e.Data()[base.EventDataFrameKey].(*event2.FrameMap)
		data, err := sonic.Marshal(frame.Data)
		if err != nil {
			return err
		}
		msgEvent := &event2.MessageKMarkdownEvent{}
		err = sonic.Unmarshal(data, msgEvent)
		if err != nil {
			return err
		}
		client := helper.NewApiHelper("/v3/message/create", Token, BaseUrl, "", "")
		if msgEvent.Author.Bot {
			log.Info("bot message")
			return nil
		}
		echoData := map[string]string{
			"channel_id": msgEvent.TargetId,
			"content":    "echo:" + msgEvent.KMarkdown.RawContent,
		}
		echoDataByte, err := sonic.Marshal(echoData)
		if err != nil {
			return err
		}
		resp, err := client.SetBody(echoDataByte).Post()
		if err != nil {
			return err
		}
		log.Infof("resp:%+v", resp)
		return nil
	}()
	if err != nil {
		log.WithError(err).Error("GroupTextEventHandler err")
	}

	return nil
}

func main() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel)

	session := base.NewWebSocketSession(Token, BaseUrl, "./session.pid", "", 1)
	session.On(base.EVENT_RECEIVE_FRAME, &ReceiveFrameHandler{})
	session.On("GROUP*", &GroupEventHandler{})
	session.On("GROUP_9", &GroupTextEventHandler{})
	session.Start()
}
