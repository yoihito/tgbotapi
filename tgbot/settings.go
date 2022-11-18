package tgbot

import "net/http"

type Settings struct {
	Token string
	ApiUrl string
	Poller Poller
	HttpClient *http.Client
}
