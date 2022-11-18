package handlers

import (
	"math/rand"
	"tgbotapi/tgbot"
)

func RandomQuote(a *AppContext, c tgbot.HandlerContext) error {
	quotes, err := a.quotesApiClient.GetQuotes()
	if err != nil {
		return err
	}

	n := rand.Int() % len(quotes)

	c.Send(quotes[n].Text)

	return nil
}
