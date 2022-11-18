package handlers

import (
	"tgbotapi/apiclients"
	"tgbotapi/tgbot"
)

type AppBotHandler struct {
	*AppContext
	HandlerFunc func(*AppContext, tgbot.HandlerContext) error
}

type AppContext struct {
	quotesApiClient *apiclients.QuotesAPIClient
}

func (ah *AppBotHandler) Handler(c tgbot.HandlerContext) error {
	return ah.HandlerFunc(ah.AppContext, c)
}

func SetupApp(bot *tgbot.Bot) {
	context := AppContext{
		quotesApiClient: apiclients.NewQuotesAPIClient(),
	}
	bot.OnText("/random_quote", (&AppBotHandler{&context, RandomQuote}).Handler)
	bot.OnText("/say_hello", SayHello)
}
