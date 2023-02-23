package main

import (
	"golang-bot/kaihela/api/base"
	"net/http"
)

func main() {

	session := &base.WebhookSession{}

	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":3333", nil)

}
