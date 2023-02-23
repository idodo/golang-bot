package main

import (
	"github.com/gookit/event"
	log "github.com/sirupsen/logrus"
	"golang-bot/kaihela/api/base"
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
	log.WithField("event", e).Info("收到频道内的文字消息.")
	return nil
}

func main() {
	token := "1/MTU1Njk=/b1r8dGxLxugSmcTop5jdjg=="
	session := base.NewWebSocketSession(token, "https://www.kookapp.cn/api", "./session.pid", "", 1)
	session.On(base.EVENT_RECEIVE_FRAME, &ReceiveFrameHandler{})
	session.On("GROUP*", &GroupEventHandler{})
	session.On("GROUP_1", &GroupTextEventHandler{})
	session.Start()
}
