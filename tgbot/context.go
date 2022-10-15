package tgbot

import "errors"

type Context interface {

}

type HandlerContext struct {
	b *Bot
	u Update
}

type HandlerFunc func(HandlerContext) error

func (c *HandlerContext) Send(what any) (*Message, error) {
	switch object := what.(type) {
	case string:
		return c.b.SendMessage(c.u.Message.Sender, object)
	default:
		return nil, errors.New("Unsupported what")
	}
}
