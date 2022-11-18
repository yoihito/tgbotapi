package tgbot

import "errors"

type Context interface {

}

type HandlerContext struct {
	bot *Bot
	update Update
}

type HandlerFunc func(HandlerContext) error

func (c *HandlerContext) Send(what any) (*Message, error) {
	switch object := what.(type) {
	case string:
		return c.bot.SendMessage(c.update.Message.Sender, object)
	default:
		return nil, errors.New("Unsupported what")
	}
}
