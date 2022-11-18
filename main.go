package main

import (
	"os"
	"tgbotapi/handlers"
	"tgbotapi/tgbot"
)

func main() {
	botSettings := tgbot.Settings{
		Token:  os.Getenv("TG_BOT_API_KEY"),
		ApiUrl: "https://api.telegram.org/bot",
		Poller: tgbot.NewPoller(),
	}
	bot, _ := tgbot.NewBot(botSettings)
	handlers.SetupApp(bot)
	bot.Start()
}
