package main

import (
	"os"

	tgbot "github.com/yoihito/tgbotapi/tgbot"
)

func main() {
	botSettings := tgbot.Settings{
		Token:  os.Getenv("TG_BOT_API_KEY"),
		ApiUrl: "https://api.telegram.org/bot",
		Poller: tgbot.NewPoller(),
	}
	bot, _ := tgbot.NewBot(botSettings)
	bot.OnText("/say_hello", func(c tgbot.HandlerContext) error {
		c.Send("Hello!")
		return nil
	})
	bot.Start()
}
