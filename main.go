package main

import (
	"net/http"
	"os"
	"tgbotapi/handlers"
	"tgbotapi/tgbot"
)

func main() {
	httpClient := &http.Client{}
	botSettings := tgbot.Settings{
		Token:  os.Getenv("TG_BOT_API_KEY"),
		ApiUrl: "https://api.telegram.org/bot",
		Poller: tgbot.NewPoller(),
		HttpClient: httpClient,
	}
	bot, _ := tgbot.NewBot(botSettings)
	handlers.SetupApp(bot, httpClient)
	bot.Start()
}
