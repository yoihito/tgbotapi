package handlers

import "tgbotapi/tgbot"

func SayHello(c tgbot.HandlerContext) error {
	c.Send("Hello!")
	return nil
}
