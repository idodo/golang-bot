package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang-bot/kaihela/api/base"
	"golang-bot/kaihela/example/conf"
	"golang-bot/kaihela/example/handler"
	"io"
	"net/http"
)

func main() {

	session := base.NewWebhookSession(conf.EncryptKey, conf.VerifyToken, 1)
	session.On(base.EventReceiveFrame, &handler.ReceiveFrameHandler{})
	session.On("GROUP*", &handler.GroupEventHandler{})
	session.On("GROUP_9", &handler.GroupTextEventHandler{})
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-Type", "application/json")
		defer req.Body.Close()
		body, err := io.ReadAll(req.Body)
		if err != nil {
			log.WithError(err).Error("Read req body error")
			return
		}
		err, resData := session.ReceiveData(body)
		if err != nil {
			log.WithError(err).Error("handle req err")
		}
		resp.Write(resData)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", conf.HTTPServerPort), nil))

}
