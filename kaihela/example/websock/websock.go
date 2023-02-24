package main

import (
	log "github.com/sirupsen/logrus"
	"golang-bot/kaihela/api/base"
	"golang-bot/kaihela/example/conf"
	"golang-bot/kaihela/example/handler"
)

func main() {
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel)

	session := base.NewWebSocketSession(conf.Token, conf.BaseUrl, "./session.pid", "", 1)
	session.On(base.EventReceiveFrame, &handler.ReceiveFrameHandler{})
	session.On("GROUP*", &handler.GroupEventHandler{})
	session.On("GROUP_9", &handler.GroupTextEventHandler{Token: conf.Token, BaseUrl: conf.BaseUrl})
	session.Start()
}
